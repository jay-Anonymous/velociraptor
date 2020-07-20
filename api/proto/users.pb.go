// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	proto1 "www.velocidex.com/golang/velociraptor/acls/proto"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GUISettings_UIMode int32

const (
	GUISettings_BASIC    GUISettings_UIMode = 0
	GUISettings_ADVANCED GUISettings_UIMode = 1
	GUISettings_DEBUG    GUISettings_UIMode = 2
)

var GUISettings_UIMode_name = map[int32]string{
	0: "BASIC",
	1: "ADVANCED",
	2: "DEBUG",
}

var GUISettings_UIMode_value = map[string]int32{
	"BASIC":    0,
	"ADVANCED": 1,
	"DEBUG":    2,
}

func (x GUISettings_UIMode) String() string {
	return proto.EnumName(GUISettings_UIMode_name, int32(x))
}

func (GUISettings_UIMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1, 0}
}

type ApiGrrUser_UserType int32

const (
	ApiGrrUser_USER_TYPE_NONE     ApiGrrUser_UserType = 0
	ApiGrrUser_USER_TYPE_STANDARD ApiGrrUser_UserType = 1
	ApiGrrUser_USER_TYPE_ADMIN    ApiGrrUser_UserType = 2
)

var ApiGrrUser_UserType_name = map[int32]string{
	0: "USER_TYPE_NONE",
	1: "USER_TYPE_STANDARD",
	2: "USER_TYPE_ADMIN",
}

var ApiGrrUser_UserType_value = map[string]int32{
	"USER_TYPE_NONE":     0,
	"USER_TYPE_STANDARD": 1,
	"USER_TYPE_ADMIN":    2,
}

func (x ApiGrrUser_UserType) String() string {
	return proto.EnumName(ApiGrrUser_UserType_name, int32(x))
}

func (ApiGrrUser_UserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{4, 0}
}

type UserNotification_Type int32

const (
	UserNotification_TYPE_UNSET                                  UserNotification_Type = 0
	UserNotification_TYPE_CLIENT_INTERROGATED                    UserNotification_Type = 1
	UserNotification_TYPE_CLIENT_APPROVAL_REQUESTED              UserNotification_Type = 2
	UserNotification_TYPE_HUNT_APPROVAL_REQUESTED                UserNotification_Type = 3
	UserNotification_TYPE_CRON_JOB_APPROVAL_REQUESTED            UserNotification_Type = 4
	UserNotification_TYPE_CLIENT_APPROVAL_GRANTED                UserNotification_Type = 5
	UserNotification_TYPE_HUNT_APPROVAL_GRANTED                  UserNotification_Type = 6
	UserNotification_TYPE_CRON_JOB_APPROVAL_GRANTED              UserNotification_Type = 7
	UserNotification_TYPE_VFS_FILE_COLLECTED                     UserNotification_Type = 8
	UserNotification_TYPE_VFS_FILE_COLLECTION_FAILED             UserNotification_Type = 9
	UserNotification_TYPE_HUNT_STOPPED                           UserNotification_Type = 10
	UserNotification_TYPE_FILE_ARCHIVE_GENERATED                 UserNotification_Type = 11
	UserNotification_TYPE_FILE_ARCHIVE_GENERATION_FAILED         UserNotification_Type = 12
	UserNotification_TYPE_FLOW_RUN_COMPLETED                     UserNotification_Type = 13
	UserNotification_TYPE_FLOW_RUN_FAILED                        UserNotification_Type = 14
	UserNotification_TYPE_VFS_LIST_DIRECTORY_COMPLETED           UserNotification_Type = 15
	UserNotification_TYPE_VFS_RECURSIVE_LIST_DIRECTORY_COMPLETED UserNotification_Type = 16
)

var UserNotification_Type_name = map[int32]string{
	0:  "TYPE_UNSET",
	1:  "TYPE_CLIENT_INTERROGATED",
	2:  "TYPE_CLIENT_APPROVAL_REQUESTED",
	3:  "TYPE_HUNT_APPROVAL_REQUESTED",
	4:  "TYPE_CRON_JOB_APPROVAL_REQUESTED",
	5:  "TYPE_CLIENT_APPROVAL_GRANTED",
	6:  "TYPE_HUNT_APPROVAL_GRANTED",
	7:  "TYPE_CRON_JOB_APPROVAL_GRANTED",
	8:  "TYPE_VFS_FILE_COLLECTED",
	9:  "TYPE_VFS_FILE_COLLECTION_FAILED",
	10: "TYPE_HUNT_STOPPED",
	11: "TYPE_FILE_ARCHIVE_GENERATED",
	12: "TYPE_FILE_ARCHIVE_GENERATION_FAILED",
	13: "TYPE_FLOW_RUN_COMPLETED",
	14: "TYPE_FLOW_RUN_FAILED",
	15: "TYPE_VFS_LIST_DIRECTORY_COMPLETED",
	16: "TYPE_VFS_RECURSIVE_LIST_DIRECTORY_COMPLETED",
}

var UserNotification_Type_value = map[string]int32{
	"TYPE_UNSET":                                  0,
	"TYPE_CLIENT_INTERROGATED":                    1,
	"TYPE_CLIENT_APPROVAL_REQUESTED":              2,
	"TYPE_HUNT_APPROVAL_REQUESTED":                3,
	"TYPE_CRON_JOB_APPROVAL_REQUESTED":            4,
	"TYPE_CLIENT_APPROVAL_GRANTED":                5,
	"TYPE_HUNT_APPROVAL_GRANTED":                  6,
	"TYPE_CRON_JOB_APPROVAL_GRANTED":              7,
	"TYPE_VFS_FILE_COLLECTED":                     8,
	"TYPE_VFS_FILE_COLLECTION_FAILED":             9,
	"TYPE_HUNT_STOPPED":                           10,
	"TYPE_FILE_ARCHIVE_GENERATED":                 11,
	"TYPE_FILE_ARCHIVE_GENERATION_FAILED":         12,
	"TYPE_FLOW_RUN_COMPLETED":                     13,
	"TYPE_FLOW_RUN_FAILED":                        14,
	"TYPE_VFS_LIST_DIRECTORY_COMPLETED":           15,
	"TYPE_VFS_RECURSIVE_LIST_DIRECTORY_COMPLETED": 16,
}

func (x UserNotification_Type) String() string {
	return proto.EnumName(UserNotification_Type_name, int32(x))
}

func (UserNotification_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{6, 0}
}

type UserNotification_State int32

const (
	UserNotification_STATE_UNSET       UserNotification_State = 0
	UserNotification_STATE_PENDING     UserNotification_State = 1
	UserNotification_STATE_NOT_PENDING UserNotification_State = 2
)

var UserNotification_State_name = map[int32]string{
	0: "STATE_UNSET",
	1: "STATE_PENDING",
	2: "STATE_NOT_PENDING",
}

var UserNotification_State_value = map[string]int32{
	"STATE_UNSET":       0,
	"STATE_PENDING":     1,
	"STATE_NOT_PENDING": 2,
}

func (x UserNotification_State) String() string {
	return proto.EnumName(UserNotification_State_name, int32(x))
}

func (UserNotification_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{6, 1}
}

type VelociraptorUser struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PasswordHash         []byte               `protobuf:"bytes,2,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	PasswordSalt         []byte               `protobuf:"bytes,3,opt,name=password_salt,json=passwordSalt,proto3" json:"password_salt,omitempty"`
	Email                string               `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Picture              string               `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty"`
	VerifiedEmail        bool                 `protobuf:"varint,6,opt,name=verified_email,json=verifiedEmail,proto3" json:"verified_email,omitempty"`
	ReadOnly             bool                 `protobuf:"varint,7,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	Locked               bool                 `protobuf:"varint,8,opt,name=locked,proto3" json:"locked,omitempty"`
	Permissions          *proto1.ApiClientACL `protobuf:"bytes,9,opt,name=Permissions,proto3" json:"Permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VelociraptorUser) Reset()         { *m = VelociraptorUser{} }
func (m *VelociraptorUser) String() string { return proto.CompactTextString(m) }
func (*VelociraptorUser) ProtoMessage()    {}
func (*VelociraptorUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *VelociraptorUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VelociraptorUser.Unmarshal(m, b)
}
func (m *VelociraptorUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VelociraptorUser.Marshal(b, m, deterministic)
}
func (m *VelociraptorUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VelociraptorUser.Merge(m, src)
}
func (m *VelociraptorUser) XXX_Size() int {
	return xxx_messageInfo_VelociraptorUser.Size(m)
}
func (m *VelociraptorUser) XXX_DiscardUnknown() {
	xxx_messageInfo_VelociraptorUser.DiscardUnknown(m)
}

var xxx_messageInfo_VelociraptorUser proto.InternalMessageInfo

func (m *VelociraptorUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VelociraptorUser) GetPasswordHash() []byte {
	if m != nil {
		return m.PasswordHash
	}
	return nil
}

func (m *VelociraptorUser) GetPasswordSalt() []byte {
	if m != nil {
		return m.PasswordSalt
	}
	return nil
}

func (m *VelociraptorUser) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *VelociraptorUser) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *VelociraptorUser) GetVerifiedEmail() bool {
	if m != nil {
		return m.VerifiedEmail
	}
	return false
}

func (m *VelociraptorUser) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *VelociraptorUser) GetLocked() bool {
	if m != nil {
		return m.Locked
	}
	return false
}

func (m *VelociraptorUser) GetPermissions() *proto1.ApiClientACL {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// Next field: 4
type GUISettings struct {
	Mode                 GUISettings_UIMode `protobuf:"varint,1,opt,name=mode,proto3,enum=proto.GUISettings_UIMode" json:"mode,omitempty"`
	CanaryMode           bool               `protobuf:"varint,3,opt,name=canary_mode,json=canaryMode,proto3" json:"canary_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GUISettings) Reset()         { *m = GUISettings{} }
func (m *GUISettings) String() string { return proto.CompactTextString(m) }
func (*GUISettings) ProtoMessage()    {}
func (*GUISettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *GUISettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GUISettings.Unmarshal(m, b)
}
func (m *GUISettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GUISettings.Marshal(b, m, deterministic)
}
func (m *GUISettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GUISettings.Merge(m, src)
}
func (m *GUISettings) XXX_Size() int {
	return xxx_messageInfo_GUISettings.Size(m)
}
func (m *GUISettings) XXX_DiscardUnknown() {
	xxx_messageInfo_GUISettings.DiscardUnknown(m)
}

var xxx_messageInfo_GUISettings proto.InternalMessageInfo

func (m *GUISettings) GetMode() GUISettings_UIMode {
	if m != nil {
		return m.Mode
	}
	return GUISettings_BASIC
}

func (m *GUISettings) GetCanaryMode() bool {
	if m != nil {
		return m.CanaryMode
	}
	return false
}

type UILink struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UILink) Reset()         { *m = UILink{} }
func (m *UILink) String() string { return proto.CompactTextString(m) }
func (*UILink) ProtoMessage()    {}
func (*UILink) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{2}
}

func (m *UILink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UILink.Unmarshal(m, b)
}
func (m *UILink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UILink.Marshal(b, m, deterministic)
}
func (m *UILink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UILink.Merge(m, src)
}
func (m *UILink) XXX_Size() int {
	return xxx_messageInfo_UILink.Size(m)
}
func (m *UILink) XXX_DiscardUnknown() {
	xxx_messageInfo_UILink.DiscardUnknown(m)
}

var xxx_messageInfo_UILink proto.InternalMessageInfo

func (m *UILink) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *UILink) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// These traits are used by the AdminUI Angular app to disable certain
// UI elements based on the user's permission set.
type ApiGrrUserInterfaceTraits struct {
	Permissions     *proto1.ApiClientACL `protobuf:"bytes,9,opt,name=Permissions,proto3" json:"Permissions,omitempty"`
	AuthUsingGoogle bool                 `protobuf:"varint,19,opt,name=auth_using_google,json=authUsingGoogle,proto3" json:"auth_using_google,omitempty"`
	Picture         string               `protobuf:"bytes,20,opt,name=picture,proto3" json:"picture,omitempty"`
	Links           []*UILink            `protobuf:"bytes,21,rep,name=links,proto3" json:"links,omitempty"`
	// An opaque setting object stored by the GUI.
	UiSettings           string   `protobuf:"bytes,1,opt,name=ui_settings,json=uiSettings,proto3" json:"ui_settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiGrrUserInterfaceTraits) Reset()         { *m = ApiGrrUserInterfaceTraits{} }
func (m *ApiGrrUserInterfaceTraits) String() string { return proto.CompactTextString(m) }
func (*ApiGrrUserInterfaceTraits) ProtoMessage()    {}
func (*ApiGrrUserInterfaceTraits) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{3}
}

func (m *ApiGrrUserInterfaceTraits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGrrUserInterfaceTraits.Unmarshal(m, b)
}
func (m *ApiGrrUserInterfaceTraits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGrrUserInterfaceTraits.Marshal(b, m, deterministic)
}
func (m *ApiGrrUserInterfaceTraits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGrrUserInterfaceTraits.Merge(m, src)
}
func (m *ApiGrrUserInterfaceTraits) XXX_Size() int {
	return xxx_messageInfo_ApiGrrUserInterfaceTraits.Size(m)
}
func (m *ApiGrrUserInterfaceTraits) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGrrUserInterfaceTraits.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGrrUserInterfaceTraits proto.InternalMessageInfo

func (m *ApiGrrUserInterfaceTraits) GetPermissions() *proto1.ApiClientACL {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *ApiGrrUserInterfaceTraits) GetAuthUsingGoogle() bool {
	if m != nil {
		return m.AuthUsingGoogle
	}
	return false
}

func (m *ApiGrrUserInterfaceTraits) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *ApiGrrUserInterfaceTraits) GetLinks() []*UILink {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *ApiGrrUserInterfaceTraits) GetUiSettings() string {
	if m != nil {
		return m.UiSettings
	}
	return ""
}

type ApiGrrUser struct {
	Username             string                     `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Settings             *GUISettings               `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
	InterfaceTraits      *ApiGrrUserInterfaceTraits `protobuf:"bytes,3,opt,name=interface_traits,json=interfaceTraits,proto3" json:"interface_traits,omitempty"`
	UserType             ApiGrrUser_UserType        `protobuf:"varint,4,opt,name=user_type,json=userType,proto3,enum=proto.ApiGrrUser_UserType" json:"user_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ApiGrrUser) Reset()         { *m = ApiGrrUser{} }
func (m *ApiGrrUser) String() string { return proto.CompactTextString(m) }
func (*ApiGrrUser) ProtoMessage()    {}
func (*ApiGrrUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{4}
}

func (m *ApiGrrUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGrrUser.Unmarshal(m, b)
}
func (m *ApiGrrUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGrrUser.Marshal(b, m, deterministic)
}
func (m *ApiGrrUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGrrUser.Merge(m, src)
}
func (m *ApiGrrUser) XXX_Size() int {
	return xxx_messageInfo_ApiGrrUser.Size(m)
}
func (m *ApiGrrUser) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGrrUser.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGrrUser proto.InternalMessageInfo

func (m *ApiGrrUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ApiGrrUser) GetSettings() *GUISettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *ApiGrrUser) GetInterfaceTraits() *ApiGrrUserInterfaceTraits {
	if m != nil {
		return m.InterfaceTraits
	}
	return nil
}

func (m *ApiGrrUser) GetUserType() ApiGrrUser_UserType {
	if m != nil {
		return m.UserType
	}
	return ApiGrrUser_USER_TYPE_NONE
}

type UserNotificationCount struct {
	Count                uint64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserNotificationCount) Reset()         { *m = UserNotificationCount{} }
func (m *UserNotificationCount) String() string { return proto.CompactTextString(m) }
func (*UserNotificationCount) ProtoMessage()    {}
func (*UserNotificationCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{5}
}

func (m *UserNotificationCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserNotificationCount.Unmarshal(m, b)
}
func (m *UserNotificationCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserNotificationCount.Marshal(b, m, deterministic)
}
func (m *UserNotificationCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserNotificationCount.Merge(m, src)
}
func (m *UserNotificationCount) XXX_Size() int {
	return xxx_messageInfo_UserNotificationCount.Size(m)
}
func (m *UserNotificationCount) XXX_DiscardUnknown() {
	xxx_messageInfo_UserNotificationCount.DiscardUnknown(m)
}

var xxx_messageInfo_UserNotificationCount proto.InternalMessageInfo

func (m *UserNotificationCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type UserNotification struct {
	Username             string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	NotificationType     UserNotification_Type  `protobuf:"varint,2,opt,name=notification_type,json=notificationType,proto3,enum=proto.UserNotification_Type" json:"notification_type,omitempty"`
	State                UserNotification_State `protobuf:"varint,3,opt,name=state,proto3,enum=proto.UserNotification_State" json:"state,omitempty"`
	Timestamp            uint64                 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message              string                 `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UserNotification) Reset()         { *m = UserNotification{} }
func (m *UserNotification) String() string { return proto.CompactTextString(m) }
func (*UserNotification) ProtoMessage()    {}
func (*UserNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{6}
}

func (m *UserNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserNotification.Unmarshal(m, b)
}
func (m *UserNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserNotification.Marshal(b, m, deterministic)
}
func (m *UserNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserNotification.Merge(m, src)
}
func (m *UserNotification) XXX_Size() int {
	return xxx_messageInfo_UserNotification.Size(m)
}
func (m *UserNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_UserNotification.DiscardUnknown(m)
}

var xxx_messageInfo_UserNotification proto.InternalMessageInfo

func (m *UserNotification) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserNotification) GetNotificationType() UserNotification_Type {
	if m != nil {
		return m.NotificationType
	}
	return UserNotification_TYPE_UNSET
}

func (m *UserNotification) GetState() UserNotification_State {
	if m != nil {
		return m.State
	}
	return UserNotification_STATE_UNSET
}

func (m *UserNotification) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *UserNotification) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetUserNotificationsResponse struct {
	Items                []*UserNotification `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetUserNotificationsResponse) Reset()         { *m = GetUserNotificationsResponse{} }
func (m *GetUserNotificationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserNotificationsResponse) ProtoMessage()    {}
func (*GetUserNotificationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{7}
}

func (m *GetUserNotificationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserNotificationsResponse.Unmarshal(m, b)
}
func (m *GetUserNotificationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserNotificationsResponse.Marshal(b, m, deterministic)
}
func (m *GetUserNotificationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserNotificationsResponse.Merge(m, src)
}
func (m *GetUserNotificationsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserNotificationsResponse.Size(m)
}
func (m *GetUserNotificationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserNotificationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserNotificationsResponse proto.InternalMessageInfo

func (m *GetUserNotificationsResponse) GetItems() []*UserNotification {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetUserNotificationsRequest struct {
	ClearPending         bool     `protobuf:"varint,1,opt,name=clear_pending,json=clearPending,proto3" json:"clear_pending,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserNotificationsRequest) Reset()         { *m = GetUserNotificationsRequest{} }
func (m *GetUserNotificationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserNotificationsRequest) ProtoMessage()    {}
func (*GetUserNotificationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{8}
}

func (m *GetUserNotificationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserNotificationsRequest.Unmarshal(m, b)
}
func (m *GetUserNotificationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserNotificationsRequest.Marshal(b, m, deterministic)
}
func (m *GetUserNotificationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserNotificationsRequest.Merge(m, src)
}
func (m *GetUserNotificationsRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserNotificationsRequest.Size(m)
}
func (m *GetUserNotificationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserNotificationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserNotificationsRequest proto.InternalMessageInfo

func (m *GetUserNotificationsRequest) GetClearPending() bool {
	if m != nil {
		return m.ClearPending
	}
	return false
}

type SetGUIOptionsRequest struct {
	Options              string   `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetGUIOptionsRequest) Reset()         { *m = SetGUIOptionsRequest{} }
func (m *SetGUIOptionsRequest) String() string { return proto.CompactTextString(m) }
func (*SetGUIOptionsRequest) ProtoMessage()    {}
func (*SetGUIOptionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{9}
}

func (m *SetGUIOptionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGUIOptionsRequest.Unmarshal(m, b)
}
func (m *SetGUIOptionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGUIOptionsRequest.Marshal(b, m, deterministic)
}
func (m *SetGUIOptionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGUIOptionsRequest.Merge(m, src)
}
func (m *SetGUIOptionsRequest) XXX_Size() int {
	return xxx_messageInfo_SetGUIOptionsRequest.Size(m)
}
func (m *SetGUIOptionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGUIOptionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetGUIOptionsRequest proto.InternalMessageInfo

func (m *SetGUIOptionsRequest) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

type Users struct {
	Users                []*VelociraptorUser `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{10}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*VelociraptorUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.GUISettings_UIMode", GUISettings_UIMode_name, GUISettings_UIMode_value)
	proto.RegisterEnum("proto.ApiGrrUser_UserType", ApiGrrUser_UserType_name, ApiGrrUser_UserType_value)
	proto.RegisterEnum("proto.UserNotification_Type", UserNotification_Type_name, UserNotification_Type_value)
	proto.RegisterEnum("proto.UserNotification_State", UserNotification_State_name, UserNotification_State_value)
	proto.RegisterType((*VelociraptorUser)(nil), "proto.VelociraptorUser")
	proto.RegisterType((*GUISettings)(nil), "proto.GUISettings")
	proto.RegisterType((*UILink)(nil), "proto.UILink")
	proto.RegisterType((*ApiGrrUserInterfaceTraits)(nil), "proto.ApiGrrUserInterfaceTraits")
	proto.RegisterType((*ApiGrrUser)(nil), "proto.ApiGrrUser")
	proto.RegisterType((*UserNotificationCount)(nil), "proto.UserNotificationCount")
	proto.RegisterType((*UserNotification)(nil), "proto.UserNotification")
	proto.RegisterType((*GetUserNotificationsResponse)(nil), "proto.GetUserNotificationsResponse")
	proto.RegisterType((*GetUserNotificationsRequest)(nil), "proto.GetUserNotificationsRequest")
	proto.RegisterType((*SetGUIOptionsRequest)(nil), "proto.SetGUIOptionsRequest")
	proto.RegisterType((*Users)(nil), "proto.Users")
}

func init() {
	proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea)
}

var fileDescriptor_030765f334c86cea = []byte{
	// 1460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x72, 0x1a, 0xcb,
	0x15, 0x36, 0x08, 0x24, 0x38, 0x48, 0x08, 0xb5, 0xa4, 0x6b, 0x2c, 0xeb, 0xda, 0x9d, 0x91, 0x5d,
	0xd2, 0xbd, 0xd7, 0x1e, 0x22, 0xdd, 0xb2, 0x5d, 0xc9, 0x26, 0x35, 0xc0, 0x08, 0x4d, 0x8c, 0x06,
	0xd2, 0x30, 0x4a, 0x9c, 0xcd, 0x54, 0x7b, 0x68, 0x60, 0xca, 0xc3, 0x0c, 0x9e, 0x6e, 0x8c, 0xf5,
	0x12, 0x59, 0xe5, 0x3d, 0xb2, 0xcd, 0x03, 0xe4, 0x29, 0xb2, 0xc8, 0x22, 0x79, 0x8a, 0xa4, 0xb2,
	0x48, 0x75, 0x0f, 0x7f, 0xc2, 0xd2, 0x26, 0x1b, 0xd1, 0x7d, 0xce, 0x77, 0xbe, 0xf3, 0xdf, 0x1a,
	0x28, 0x4c, 0x38, 0x8b, 0xb9, 0x3e, 0x8e, 0x23, 0x11, 0xa1, 0xac, 0xfa, 0x39, 0x3a, 0x50, 0x3f,
	0x15, 0xce, 0x46, 0x34, 0x14, 0xbe, 0xa7, 0xcf, 0xa4, 0xd4, 0x0b, 0x78, 0x25, 0x51, 0x51, 0x2f,
	0x48, 0xa4, 0xda, 0xbf, 0x37, 0xa0, 0x74, 0xc3, 0x82, 0xc8, 0xf3, 0x63, 0x3a, 0x16, 0x51, 0xec,
	0x70, 0x16, 0xa3, 0x33, 0xc8, 0x84, 0x74, 0xc4, 0xca, 0x29, 0x9c, 0x3a, 0xcb, 0x57, 0x0f, 0xfe,
	0xf9, 0xdf, 0x7f, 0xfd, 0x2d, 0x55, 0x44, 0xdb, 0xdd, 0x21, 0xc3, 0xd2, 0x9d, 0xd4, 0x11, 0x85,
	0x40, 0x16, 0xec, 0x8c, 0x29, 0xe7, 0xd3, 0x28, 0xee, 0xb9, 0x43, 0xca, 0x87, 0xe5, 0x34, 0x4e,
	0x9d, 0x6d, 0x57, 0x5f, 0x28, 0x93, 0x67, 0xe8, 0xb8, 0x73, 0x65, 0x5c, 0xbc, 0x79, 0x8b, 0xa5,
	0x0a, 0x47, 0x7d, 0x2c, 0x86, 0x0c, 0xcf, 0xf1, 0x3a, 0xd9, 0x9e, 0x1f, 0xaf, 0x28, 0x1f, 0xa2,
	0x93, 0x15, 0x2a, 0x4e, 0x03, 0x51, 0xde, 0x90, 0x54, 0x4b, 0x50, 0x87, 0x06, 0x02, 0x1d, 0x40,
	0x96, 0x8d, 0xa8, 0x1f, 0x94, 0x33, 0x32, 0x34, 0x92, 0x5c, 0x50, 0x19, 0xb6, 0xc6, 0xbe, 0x27,
	0x26, 0x31, 0x2b, 0x67, 0x95, 0x7c, 0x7e, 0x45, 0x2f, 0xa1, 0xf8, 0x85, 0xc5, 0x7e, 0xdf, 0x67,
	0x3d, 0x37, 0x31, 0xdc, 0xc4, 0xa9, 0xb3, 0x1c, 0xd9, 0x99, 0x4b, 0x4d, 0x45, 0x10, 0x43, 0x3e,
	0x66, 0xb4, 0xe7, 0x46, 0x61, 0x70, 0x5b, 0xde, 0x92, 0x88, 0xaa, 0xa3, 0x52, 0x68, 0xa1, 0x6b,
	0x03, 0x4b, 0x15, 0x96, 0x2a, 0x95, 0x3d, 0xf6, 0x68, 0x28, 0x0f, 0x2a, 0x99, 0x86, 0x63, 0xe1,
	0x8f, 0x13, 0x81, 0x7d, 0x8e, 0xc3, 0x48, 0x60, 0x1a, 0x04, 0xd1, 0x94, 0xf5, 0xb0, 0x88, 0x70,
	0x40, 0x27, 0xa1, 0x37, 0xc4, 0xfd, 0x20, 0x9a, 0x72, 0x1c, 0xc5, 0x78, 0x38, 0x09, 0x05, 0xd7,
	0x49, 0x4e, 0x92, 0xb5, 0xc2, 0xe0, 0x16, 0x59, 0xb0, 0x19, 0x44, 0xde, 0x27, 0xd6, 0x2b, 0xe7,
	0x94, 0xc3, 0x73, 0xe5, 0xf0, 0x27, 0xf4, 0x83, 0xd5, 0xc7, 0x9c, 0x09, 0x45, 0xaf, 0xfc, 0xdd,
	0xc3, 0x1d, 0x0d, 0xb0, 0x1f, 0xea, 0x64, 0x46, 0x80, 0xde, 0x40, 0xa1, 0xcd, 0xe2, 0x91, 0xcf,
	0xb9, 0x1f, 0x85, 0xbc, 0x9c, 0xc7, 0xa9, 0xb3, 0xc2, 0xc5, 0x7e, 0xd2, 0x61, 0xdd, 0x18, 0xfb,
	0xb5, 0xc0, 0x67, 0xa1, 0x30, 0x6a, 0x4d, 0xb2, 0x8a, 0xd3, 0xfe, 0x94, 0x86, 0x42, 0xc3, 0xb1,
	0x3a, 0x4c, 0x08, 0x3f, 0x1c, 0x70, 0xf4, 0x1e, 0x32, 0xa3, 0xa8, 0x97, 0xb4, 0xbd, 0x78, 0xf1,
	0x64, 0x66, 0xbf, 0x82, 0xd0, 0x1d, 0xeb, 0x3a, 0xea, 0xb1, 0xea, 0xb1, 0x0a, 0xf5, 0x3b, 0x74,
	0xe0, 0xa8, 0xf8, 0x42, 0xc1, 0xe2, 0x3e, 0xf5, 0x18, 0x96, 0xd6, 0x3a, 0x51, 0x24, 0xe8, 0x0f,
	0x50, 0xf0, 0x68, 0x48, 0xe3, 0x5b, 0x57, 0x71, 0x6e, 0xa8, 0x1c, 0xdf, 0x29, 0xc3, 0x73, 0x54,
	0xb1, 0xfa, 0x58, 0xc4, 0x13, 0xf6, 0x0a, 0xf3, 0x61, 0x34, 0xc5, 0x7d, 0x46, 0x65, 0xb7, 0x38,
	0x16, 0x43, 0x2a, 0x30, 0x8d, 0x19, 0xfe, 0xc8, 0xfc, 0x70, 0x80, 0x15, 0x81, 0xcf, 0x7a, 0x3a,
	0x81, 0x84, 0x4b, 0x7a, 0xd7, 0x5e, 0xc1, 0x66, 0x12, 0x07, 0xca, 0x43, 0xb6, 0x6a, 0x74, 0xac,
	0x5a, 0xe9, 0x11, 0xda, 0x86, 0x9c, 0x51, 0xbf, 0x31, 0xec, 0x9a, 0x59, 0x2f, 0xa5, 0xa4, 0xa2,
	0x6e, 0x56, 0x9d, 0x46, 0x29, 0xfd, 0xeb, 0x1f, 0xff, 0x2e, 0x5d, 0xbe, 0x00, 0x4d, 0xc5, 0x2a,
	0xfb, 0xc5, 0x67, 0xd9, 0x60, 0x1a, 0xf6, 0xf0, 0x38, 0x66, 0x7d, 0x16, 0xb3, 0xd0, 0x63, 0x5c,
	0xd7, 0x74, 0xc9, 0xdc, 0xf4, 0xc3, 0x4f, 0x08, 0x41, 0x46, 0xb0, 0xaf, 0x22, 0xd9, 0x00, 0xa2,
	0xce, 0xa8, 0x04, 0x1b, 0x93, 0x38, 0x50, 0x13, 0x9e, 0x27, 0xf2, 0xa8, 0xfd, 0x23, 0x05, 0x4f,
	0x8c, 0xb1, 0xdf, 0x88, 0xd5, 0xda, 0x58, 0xf3, 0x3a, 0x74, 0x63, 0xea, 0x0b, 0xfe, 0x7f, 0x76,
	0x05, 0xfd, 0x08, 0x7b, 0x74, 0x22, 0x86, 0xee, 0x84, 0xfb, 0xe1, 0xc0, 0x1d, 0x44, 0xd1, 0x20,
	0x60, 0xe5, 0x7d, 0x35, 0xb5, 0xbb, 0x52, 0xe1, 0x48, 0x79, 0x43, 0x89, 0x57, 0x07, 0xff, 0xe0,
	0xee, 0xe0, 0x9f, 0x40, 0x36, 0xf0, 0xc3, 0x4f, 0xbc, 0x7c, 0x88, 0x37, 0xce, 0x0a, 0x17, 0x3b,
	0x33, 0xb7, 0x49, 0x7a, 0x24, 0xd1, 0xa1, 0xe7, 0x50, 0x98, 0xf8, 0xee, 0xbc, 0x1e, 0xb3, 0x64,
	0x61, 0xe2, 0xcf, 0xfb, 0xad, 0xfd, 0x65, 0x03, 0x60, 0x99, 0x20, 0xfa, 0x15, 0xe4, 0xe6, 0xfb,
	0x3f, 0x7b, 0x1b, 0xbe, 0x57, 0x0d, 0x7d, 0x8c, 0x0e, 0xe5, 0xdb, 0x20, 0xe5, 0xf3, 0x2d, 0x97,
	0x38, 0x9d, 0x2c, 0xe0, 0xc8, 0x82, 0xdc, 0xc2, 0x4f, 0x5a, 0x55, 0x02, 0x7d, 0x3b, 0x5f, 0xd5,
	0x27, 0x8a, 0x6e, 0x1f, 0xed, 0xa9, 0x66, 0xad, 0xf4, 0x4a, 0x27, 0x0b, 0x73, 0xf4, 0xe7, 0x14,
	0x94, 0x16, 0x33, 0xe7, 0x0a, 0x55, 0x6c, 0x35, 0x5f, 0x85, 0x0b, 0xbc, 0xac, 0xee, 0xfd, 0x4d,
	0xa9, 0x36, 0x94, 0x07, 0x03, 0xfd, 0x46, 0x2a, 0x4f, 0xf9, 0xca, 0xf0, 0x26, 0x44, 0xf8, 0x6c,
	0x2a, 0x67, 0x50, 0x0c, 0xd9, 0xad, 0x5a, 0x74, 0x39, 0x23, 0x1e, 0x0d, 0x4f, 0x05, 0xee, 0x45,
	0xd8, 0x0f, 0x55, 0x6e, 0x8e, 0xf5, 0x83, 0x4e, 0x76, 0xfd, 0xb5, 0x76, 0xbf, 0x83, 0xbc, 0xcc,
	0xd6, 0x15, 0xb7, 0x63, 0xa6, 0x9e, 0xa7, 0xe2, 0xc5, 0xd1, 0x37, 0xe1, 0xe8, 0xf2, 0x4f, 0xf7,
	0x76, 0xcc, 0x92, 0xd2, 0xc8, 0x93, 0xf6, 0x1e, 0x72, 0x73, 0x29, 0x42, 0x50, 0x74, 0x3a, 0x26,
	0x71, 0xbb, 0x1f, 0xda, 0xa6, 0x6b, 0xb7, 0x6c, 0xb3, 0xf4, 0x08, 0x7d, 0x07, 0x68, 0x29, 0xeb,
	0x74, 0x0d, 0xbb, 0x6e, 0x10, 0x39, 0xe4, 0xfb, 0xb0, 0xbb, 0x94, 0x1b, 0xf5, 0x6b, 0xcb, 0x2e,
	0xa5, 0xb5, 0xd7, 0x70, 0x28, 0xc9, 0xec, 0x48, 0xf8, 0x7d, 0xdf, 0xa3, 0xc2, 0x8f, 0xc2, 0x5a,
	0x34, 0x09, 0xd5, 0xcb, 0xe9, 0xc9, 0x83, 0x6a, 0x5c, 0x86, 0x24, 0x17, 0xed, 0x3f, 0x9b, 0x50,
	0x5a, 0xc7, 0xa3, 0xa3, 0xf5, 0x36, 0xdf, 0xe9, 0xe3, 0x5e, 0xb8, 0x82, 0x4d, 0xb2, 0x4d, 0xab,
	0x6c, 0x8f, 0xe7, 0x33, 0xb6, 0xc6, 0xa7, 0xab, 0x7c, 0x4b, 0xab, 0x66, 0x2a, 0xd7, 0x9f, 0x21,
	0xcb, 0x05, 0x15, 0xc9, 0xdb, 0x50, 0xbc, 0xf8, 0xfe, 0x21, 0xf3, 0x8e, 0x04, 0x91, 0x04, 0x8b,
	0xce, 0x21, 0x2f, 0xfc, 0x11, 0xe3, 0x82, 0x8e, 0xc6, 0xaa, 0xca, 0x99, 0xea, 0xbe, 0x6a, 0xe9,
	0x0e, 0x14, 0x48, 0xfd, 0xb2, 0x4e, 0x05, 0x93, 0x7a, 0xb2, 0x44, 0xc9, 0x25, 0x19, 0x31, 0xce,
	0xe9, 0x60, 0xf1, 0xdf, 0x61, 0x76, 0xd5, 0xfe, 0x9a, 0x81, 0x8c, 0x0a, 0xa5, 0x08, 0xa0, 0xaa,
	0xe8, 0xd8, 0x1d, 0xb3, 0x5b, 0x7a, 0x84, 0x8e, 0xa1, 0xac, 0xee, 0xb5, 0xa6, 0x65, 0xda, 0x5d,
	0xd7, 0xb2, 0xbb, 0x26, 0x21, 0xad, 0x86, 0xd1, 0x55, 0xaf, 0x8b, 0x06, 0xcf, 0x56, 0xb5, 0x46,
	0xbb, 0x4d, 0x5a, 0x37, 0x46, 0xd3, 0x25, 0xe6, 0xef, 0x1c, 0xb3, 0x23, 0x31, 0x69, 0x84, 0xe1,
	0x58, 0x61, 0xae, 0x9c, 0xfb, 0x11, 0x1b, 0xe8, 0x05, 0xe0, 0x84, 0x85, 0xb4, 0x6c, 0xf7, 0xb7,
	0xad, 0xea, 0x7d, 0xa8, 0xcc, 0x82, 0x67, 0xdd, 0x57, 0x83, 0x18, 0xb6, 0x44, 0x64, 0xd1, 0x33,
	0x38, 0xba, 0xc7, 0xd3, 0x5c, 0xbf, 0xb9, 0x8c, 0xf6, 0x1b, 0x3f, 0x73, 0xcc, 0x16, 0x7a, 0x0a,
	0x8f, 0x15, 0xe6, 0xe6, 0xb2, 0xe3, 0x5e, 0x5a, 0x4d, 0xd3, 0xad, 0xb5, 0x9a, 0x4d, 0xb3, 0x26,
	0x95, 0x39, 0x74, 0x02, 0xcf, 0xef, 0x55, 0x5a, 0x2d, 0xdb, 0xbd, 0x34, 0xac, 0xa6, 0x59, 0x2f,
	0xe5, 0xd1, 0x21, 0xec, 0x2d, 0xa3, 0xe8, 0x74, 0x5b, 0xed, 0xb6, 0x59, 0x2f, 0x01, 0x7a, 0x0e,
	0x4f, 0x95, 0x58, 0xd9, 0x19, 0xa4, 0x76, 0x65, 0xdd, 0x98, 0x6e, 0xc3, 0xb4, 0x4d, 0xa2, 0x6a,
	0x59, 0x40, 0xa7, 0x70, 0xf2, 0x20, 0x60, 0xc5, 0xc1, 0xf6, 0x22, 0xc4, 0xcb, 0x66, 0xeb, 0xf7,
	0x2e, 0x71, 0x6c, 0xb7, 0xd6, 0xba, 0x6e, 0x37, 0x4d, 0xc9, 0xb2, 0x83, 0xca, 0x70, 0x70, 0x57,
	0x39, 0x33, 0x2b, 0xa2, 0x97, 0xf0, 0x8b, 0x45, 0xf0, 0x4d, 0xab, 0xd3, 0x75, 0xeb, 0x16, 0x31,
	0x6b, 0xdd, 0x16, 0xf9, 0xb0, 0x42, 0xb0, 0x8b, 0x2a, 0xf0, 0xd3, 0x02, 0x46, 0xcc, 0x9a, 0x43,
	0x3a, 0x32, 0x8e, 0x07, 0x0d, 0x4a, 0x5a, 0x15, 0xb2, 0x6a, 0x2e, 0xd1, 0x2e, 0x14, 0x3a, 0x5d,
	0xa3, 0xbb, 0x9c, 0x9d, 0x3d, 0xd8, 0x49, 0x04, 0x6d, 0xd3, 0xae, 0x5b, 0x76, 0xa3, 0x94, 0x92,
	0xc5, 0x49, 0x44, 0x76, 0xab, 0xbb, 0x10, 0xa7, 0xb5, 0x6b, 0x38, 0x6e, 0x30, 0xb1, 0x3e, 0xef,
	0x9c, 0x30, 0x3e, 0x8e, 0x42, 0xce, 0xd0, 0x6b, 0xc8, 0xfa, 0x82, 0x8d, 0xe4, 0xc3, 0x2c, 0xdf,
	0xf0, 0xc7, 0x0f, 0x2c, 0x08, 0x49, 0x50, 0xda, 0x67, 0x78, 0x7a, 0x3f, 0xdd, 0xe7, 0x09, 0xe3,
	0x02, 0x11, 0xd8, 0xf1, 0x02, 0x46, 0x63, 0x77, 0xcc, 0xc2, 0x9e, 0x1f, 0x0e, 0xd4, 0x6a, 0xe7,
	0xaa, 0xaf, 0xd5, 0xf6, 0x9c, 0xa2, 0x97, 0xc9, 0x67, 0xc7, 0x2b, 0xac, 0x40, 0x1c, 0xcf, 0x50,
	0x78, 0x75, 0x6d, 0xb9, 0x4e, 0xb6, 0x95, 0xba, 0x9d, 0x28, 0xb5, 0x5f, 0xc2, 0x41, 0x87, 0x89,
	0x86, 0x63, 0xb5, 0xc6, 0x77, 0x7c, 0x95, 0x61, 0x2b, 0x4a, 0x24, 0xb3, 0x07, 0x64, 0x7e, 0xd5,
	0xde, 0x42, 0x56, 0x46, 0xc8, 0x65, 0x72, 0xea, 0xd3, 0x75, 0x2d, 0xb9, 0xf5, 0x6f, 0x51, 0x92,
	0xa0, 0xaa, 0xe7, 0x7f, 0xac, 0x4c, 0xa7, 0x53, 0xfd, 0x8b, 0x52, 0xf7, 0xd8, 0x57, 0xdd, 0x8b,
	0x46, 0x95, 0x41, 0x14, 0xd0, 0x70, 0x50, 0xf9, 0xb2, 0x62, 0x53, 0xa1, 0x63, 0x3f, 0xf9, 0xc8,
	0xfd, 0xb8, 0xa9, 0x7e, 0x7e, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x65, 0x5d, 0x2f,
	0x23, 0x0b, 0x00, 0x00,
}
