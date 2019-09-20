#!/usr/bin/python3
"""Convert KapeFiles to a Velociraptor artifact.

Kape is a popular target file collector used to triage a live
system. It operates by scanning a group of yaml files called Targets
which describe simple rules about how to find these files. Although
Kape is not open source, the rules driving its actions are stored in
the KapeFiles repository on github.

This script converts the Target yaml files within the KapeFiles
repository to a data driven Velociraptor artifact.

Simply point this script at the root a directory containing ".tkape"
files and we will generate the artifact to stdout.

"""

import argparse
import io
import sys
import csv
import re
import os
import yaml

BLACKLISTED = ["!ALL.tkape"]


class KapeContext:
    groups = {}
    rows = [["Id", "Name", "Category", "Glob", "Comment"]]
    kape_files = []
    kape_data = {}

def read_targets(ctx, project_path):
    for root, dirs, files in os.walk(
            project_path + "/Targets", topdown=False):

        for name in files:
            if not name.endswith(".tkape") or name in BLACKLISTED:
                continue

            ctx.kape_files.append(name)

            try:
                full_path = os.path.join(root, name)
                ctx.kape_data[name] = yaml.safe_load(open(full_path).read())
            except Exception as e:
                print ("Unable to parse %s: %s" % (full_path, e))

            ctx.groups[name] = set()

    for name, data in ctx.kape_data.items():
        for target in data["Targets"]:
            glob = target.get("Path", "")

            if target.get("Recursive"):
                glob += "/**"

            mask = target.get("FileMask")
            if mask:
                glob += "/" + mask

            # Expand the targets in the glob
            if ".tkape" in glob:
                continue

            row_id = len(ctx.rows)
            ctx.groups[name].add(row_id)

            ctx.rows.append([
                row_id,
                target["Name"],
                target.get("Category", ""),
                strip_drive(glob),
                target.get("Comment", "")])

    for i in range(3):
        for name, data in ctx.kape_data.items():
            for target in data["Targets"]:
                glob = target.get("Path", "")
                if ".tkape" in glob:
                    deps = ctx.groups.get(glob)
                    if deps is None:
                        sys.stderr.write("Unable to process dependency %s (%s)\n" %
                                         (name, glob))
                        #import pdb; pdb.set_trace()
                        continue

                    for dependency in deps:
                        ctx.groups[name].add(dependency)

def sanitize(name):
    name = name.replace(".tkape", "")
    name = re.sub("[^a-zA-Z0-9]", "_", name)
    return name


def strip_drive(name):
    return re.sub("^[a-zA-Z]:\\\\", "", name)

def get_csv(rows):
    out = io.StringIO()
    writer = csv.writer(out)
    for row in rows:
        writer.writerow(row)

    return out.getvalue()

def get_yara(ctx):
    result = []
    for i, row in enumerate(ctx.rows[1:]):
        glob = row[3]   # Glob
        glob = re.sub("[/\\\\]+", "\\/", glob)
        glob = glob.replace("**", ".*")
        glob = re.sub("([^\\.])\\*", "\\1[^\\/]*", glob)

        result.append("rule T%s {strings: $=/%s/ nocase condition: any of them}" % (
            i, glob))
    return "\n".join(result)

def format(ctx):
    template = """name: Windows.Kape.Targets
description: |

    Kape is a popular bulk collector tool for triaging a system
    quickly. While KAPE itself is not an opensource tool, the logic it
    uses to decide which files to collect is encoded in YAML files
    hosted on the KapeFiles project
    (https://github.com/EricZimmerman/KapeFiles) and released under an
    MIT license.

    This artifact is automatically generated from these YAML files,
    contributed and maintained by the community. This artifact only
    encapsulates the KAPE "Targets" - basically a bunch of glob
    expressions used for collecting files on the endpoint. We do not
    do any post processing these files - we just collect them.

    We recommend that timeouts and upload limits be used
    conservatively with this artifact because we can upload really
    vast quantities of data very quickly.

reference:
  - https://www.kroll.com/en/insights/publications/cyber/kroll-artifact-parser-extractor-kape
  - https://github.com/EricZimmerman/KapeFiles

parameters:
%(parameters)s
  - name: KapeRules
    type: hidden
    description: A CSV file controlling the different Kape Target Rules
    default: |
%(csv)s
  - name: KapeTargets
    type: hidden
    description: Each parameter above represents a group of rules to be triggered. This table specifies which rule IDs will be included when the parameter is checked.
    default: |
%(rules)s
  - name: Device
    default: C:\\
  - name: Accessor
    default: auto

sources:
  - queries:
      # Select all the rules Ids to be included depending on the group selection.
      - |
        LET targets <= SELECT * FROM parse_csv(filename=KapeTargets, accessor="data")
        WHERE get(item=scope(), member=Group)

      # Filter only the rules in the rule table that have an Id we want.
      - |
        LET specs <= SELECT Id, Device + Glob AS Glob
        FROM parse_csv(filename=KapeRules, accessor="data")
        WHERE Id in array(array=targets.RuleIds)

      # Join all the rules into a single Glob plugin. This ensure we
      # only make one pass over the filesystem.
      - |
        LET result = SELECT timestamp(epoch=now()) AS CopiedOnTimestamp,
               FullPath AS SourceFile, Size,
               timestamp(epoch=Ctime.Sec) AS CreatedOnUtc,
               timestamp(epoch=Mtime.Sec) AS ModifiedOnUtc,
               timestamp(epoch=Atime.Sec) AS LastAccessedOnUtc,
               upload(file=FullPath, accessor=Accessor) AS Upload
        FROM glob(globs=specs.Glob, accessor=Accessor)
        WHERE NOT IsDir

      - |
        SELECT CopiedOnTimestamp, SourceFile, Upload.Path AS DestinationFile,
               Size AS FileSize, Upload.sha256 AS SourceFileSha256,
               CreatedOnUtc, ModifiedOnUtc, LastAccessedOnUtc
        FROM result

"""
    parameters_str = ""
    rules = [["Group", "RuleIds"]]

    for k, v in sorted(ctx.groups.items()):
        parameters_str += "  - name: %s\n    description: \"%s (by %s): %s\"\n    type: bool\n" % (
            sanitize(k),
            ctx.kape_data[k].get("Description"),
            ctx.kape_data[k].get("Author"),
            ", ".join([ctx.rows[x][1] for x in v]))

        ids = ['%s' % x for x in v]
        if len(ids) > 0:
            rules.append([sanitize(k), sorted(v)])

    print (template % dict(
        parameters=parameters_str,
        rules="\n".join(["      " + x for x in get_csv(rules).splitlines()]),
        csv="\n".join(["      " + x for x in get_csv(ctx.rows).splitlines()]),
        yara="\n".join(["      " + x for x in get_yara(ctx).splitlines()]),
    ))


if __name__ == "__main__":
    argument_parser = argparse.ArgumentParser()
    argument_parser.add_argument("kape_file_path", help="Path to the KapeFiles project")

    args = argument_parser.parse_args()

    ctx = KapeContext()
    read_targets(ctx, args.kape_file_path)
    format(ctx)
