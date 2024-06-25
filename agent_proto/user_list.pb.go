// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_list.proto

package agent_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type UserListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListRequest) Reset()         { *m = UserListRequest{} }
func (m *UserListRequest) String() string { return proto.CompactTextString(m) }
func (*UserListRequest) ProtoMessage()    {}
func (*UserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_157f19e16d90e583, []int{0}
}

func (m *UserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListRequest.Unmarshal(m, b)
}
func (m *UserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListRequest.Marshal(b, m, deterministic)
}
func (m *UserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListRequest.Merge(m, src)
}
func (m *UserListRequest) XXX_Size() int {
	return xxx_messageInfo_UserListRequest.Size(m)
}
func (m *UserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserListRequest proto.InternalMessageInfo

type UserListResponse struct {
	Message              string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	List                 []*UserInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_157f19e16d90e583, []int{1}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserListResponse) GetList() []*UserInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type UserInfo struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	UserType             string   `protobuf:"bytes,3,opt,name=UserType,proto3" json:"UserType,omitempty"`
	GroupID              string   `protobuf:"bytes,4,opt,name=groupID,proto3" json:"groupID,omitempty"`
	GroupName            string   `protobuf:"bytes,5,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Comment              string   `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
	HomeDir              string   `protobuf:"bytes,7,opt,name=homeDir,proto3" json:"homeDir,omitempty"`
	Shell                string   `protobuf:"bytes,8,opt,name=shell,proto3" json:"shell,omitempty"`
	LastLoginTime        string   `protobuf:"bytes,9,opt,name=lastLoginTime,proto3" json:"lastLoginTime,omitempty"`
	Status               string   `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_157f19e16d90e583, []int{2}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfo) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserInfo) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

func (m *UserInfo) GetGroupID() string {
	if m != nil {
		return m.GroupID
	}
	return ""
}

func (m *UserInfo) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *UserInfo) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *UserInfo) GetHomeDir() string {
	if m != nil {
		return m.HomeDir
	}
	return ""
}

func (m *UserInfo) GetShell() string {
	if m != nil {
		return m.Shell
	}
	return ""
}

func (m *UserInfo) GetLastLoginTime() string {
	if m != nil {
		return m.LastLoginTime
	}
	return ""
}

func (m *UserInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*UserListRequest)(nil), "agent_proto.UserListRequest")
	proto.RegisterType((*UserListResponse)(nil), "agent_proto.UserListResponse")
	proto.RegisterType((*UserInfo)(nil), "agent_proto.UserInfo")
}

func init() {
	proto.RegisterFile("user_list.proto", fileDescriptor_157f19e16d90e583)
}

var fileDescriptor_157f19e16d90e583 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xdd, 0xff, 0xed, 0x96, 0x51, 0x0d, 0x2a, 0x61, 0x4c, 0x18, 0xc5, 0x87, 0xf9, 0x52,
	0x61, 0x7e, 0x03, 0x29, 0x48, 0xa1, 0xf8, 0x50, 0x26, 0x82, 0x3e, 0x8c, 0x2a, 0xd7, 0xae, 0xd0,
	0x34, 0xb5, 0x37, 0x7d, 0xf0, 0xab, 0xf8, 0x69, 0x25, 0x49, 0xeb, 0x3a, 0xc4, 0xa7, 0xe6, 0x77,
	0x4e, 0xce, 0xe9, 0x4d, 0x02, 0x6e, 0x4d, 0x58, 0xed, 0xf2, 0x8c, 0x94, 0x5f, 0x56, 0x52, 0x49,
	0xe6, 0x24, 0x29, 0x16, 0x6a, 0x67, 0xc0, 0x3b, 0x03, 0xf7, 0x89, 0xb0, 0x8a, 0x32, 0x52, 0x31,
	0x7e, 0xd6, 0x48, 0xca, 0x7b, 0x86, 0xd3, 0x83, 0x44, 0xa5, 0x2c, 0x08, 0x19, 0x87, 0x89, 0x40,
	0xa2, 0x24, 0x45, 0xde, 0x5b, 0xf5, 0xd6, 0xb3, 0xb8, 0x45, 0x76, 0x03, 0x43, 0xdd, 0xcd, 0xfb,
	0xab, 0xc1, 0xda, 0xd9, 0x5c, 0xf8, 0x9d, 0x72, 0x5f, 0xd7, 0x84, 0xc5, 0x87, 0x8c, 0xcd, 0x16,
	0xef, 0xbb, 0x0f, 0xd3, 0x56, 0x62, 0x0b, 0x98, 0xea, 0xc1, 0x1e, 0x13, 0xd1, 0x56, 0xfe, 0x32,
	0xbb, 0x84, 0xb1, 0x5e, 0x87, 0x01, 0xef, 0x1b, 0xa7, 0x21, 0x9d, 0xd1, 0xf9, 0xed, 0x57, 0x89,
	0x7c, 0x60, 0x33, 0x2d, 0xeb, 0x09, 0xd3, 0x4a, 0xd6, 0x65, 0x18, 0xf0, 0xa1, 0x9d, 0xb0, 0x41,
	0xb6, 0x84, 0x99, 0x59, 0x9a, 0x5f, 0x8d, 0x8c, 0x77, 0x10, 0x74, 0xee, 0x5d, 0x0a, 0x81, 0x85,
	0xe2, 0x63, 0x9b, 0x6b, 0x50, 0x3b, 0x7b, 0x29, 0x30, 0xc8, 0x2a, 0x3e, 0xb1, 0x4e, 0x83, 0xec,
	0x1c, 0x46, 0xb4, 0xc7, 0x3c, 0xe7, 0x53, 0xa3, 0x5b, 0x60, 0xd7, 0x30, 0xcf, 0x13, 0x52, 0x91,
	0x4c, 0xb3, 0x62, 0x9b, 0x09, 0xe4, 0x33, 0xe3, 0x1e, 0x8b, 0xfa, 0x6c, 0xa4, 0x12, 0x55, 0x13,
	0x07, 0x7b, 0x36, 0x4b, 0x9b, 0x57, 0x70, 0x1e, 0x50, 0xb5, 0x17, 0xcf, 0xa2, 0x63, 0x5c, 0xfe,
	0xb9, 0xd7, 0xce, 0x8b, 0x2d, 0xae, 0xfe, 0x71, 0xed, 0xe3, 0x79, 0x27, 0xf7, 0xee, 0xcb, 0xdc,
	0xbf, 0xed, 0xec, 0x79, 0x1b, 0x9b, 0xcf, 0xdd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xa4,
	0xba, 0xad, 0x1d, 0x02, 0x00, 0x00,
}
