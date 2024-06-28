// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: system.proto

package agent_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetSysInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSysInfoRequest) Reset() {
	*x = GetSysInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSysInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysInfoRequest) ProtoMessage() {}

func (x *GetSysInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysInfoRequest.ProtoReflect.Descriptor instead.
func (*GetSysInfoRequest) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{0}
}

type GetSysInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message       string         `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Platform      *PlatformModel `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	CpuModel      *CpuModel      `protobuf:"bytes,3,opt,name=cpuModel,proto3" json:"cpuModel,omitempty"`
	Timezone      string         `protobuf:"bytes,4,opt,name=timezone,proto3" json:"timezone,omitempty"`
	KernelVersion string         `protobuf:"bytes,5,opt,name=kernelVersion,proto3" json:"kernelVersion,omitempty"`
	Os            string         `protobuf:"bytes,6,opt,name=os,proto3" json:"os,omitempty"`
	Arch          string         `protobuf:"bytes,7,opt,name=arch,proto3" json:"arch,omitempty"`
	Hostname      string         `protobuf:"bytes,8,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Uptime        string         `protobuf:"bytes,9,opt,name=uptime,proto3" json:"uptime,omitempty"`
	BootTime      string         `protobuf:"bytes,10,opt,name=bootTime,proto3" json:"bootTime,omitempty"`
	LoadAverage   *LoadAverage   `protobuf:"bytes,11,opt,name=loadAverage,proto3" json:"loadAverage,omitempty"`
	VirtualMemory *MemoryStat    `protobuf:"bytes,12,opt,name=virtualMemory,proto3" json:"virtualMemory,omitempty"`
	SwapMemory    *MemoryStat    `protobuf:"bytes,13,opt,name=swapMemory,proto3" json:"swapMemory,omitempty"`
}

func (x *GetSysInfoResponse) Reset() {
	*x = GetSysInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSysInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysInfoResponse) ProtoMessage() {}

func (x *GetSysInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysInfoResponse.ProtoReflect.Descriptor instead.
func (*GetSysInfoResponse) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{1}
}

func (x *GetSysInfoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetSysInfoResponse) GetPlatform() *PlatformModel {
	if x != nil {
		return x.Platform
	}
	return nil
}

func (x *GetSysInfoResponse) GetCpuModel() *CpuModel {
	if x != nil {
		return x.CpuModel
	}
	return nil
}

func (x *GetSysInfoResponse) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *GetSysInfoResponse) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

func (x *GetSysInfoResponse) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *GetSysInfoResponse) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *GetSysInfoResponse) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *GetSysInfoResponse) GetUptime() string {
	if x != nil {
		return x.Uptime
	}
	return ""
}

func (x *GetSysInfoResponse) GetBootTime() string {
	if x != nil {
		return x.BootTime
	}
	return ""
}

func (x *GetSysInfoResponse) GetLoadAverage() *LoadAverage {
	if x != nil {
		return x.LoadAverage
	}
	return nil
}

func (x *GetSysInfoResponse) GetVirtualMemory() *MemoryStat {
	if x != nil {
		return x.VirtualMemory
	}
	return nil
}

func (x *GetSysInfoResponse) GetSwapMemory() *MemoryStat {
	if x != nil {
		return x.SwapMemory
	}
	return nil
}

type PlatformModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform string `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Family   string `protobuf:"bytes,2,opt,name=family,proto3" json:"family,omitempty"`
	Version  string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *PlatformModel) Reset() {
	*x = PlatformModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformModel) ProtoMessage() {}

func (x *PlatformModel) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformModel.ProtoReflect.Descriptor instead.
func (*PlatformModel) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{2}
}

func (x *PlatformModel) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *PlatformModel) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

func (x *PlatformModel) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type CpuModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName     string `protobuf:"bytes,1,opt,name=modelName,proto3" json:"modelName,omitempty"`
	PhysicalCount int32  `protobuf:"varint,2,opt,name=physicalCount,proto3" json:"physicalCount,omitempty"`
	LogicalCount  int32  `protobuf:"varint,3,opt,name=logicalCount,proto3" json:"logicalCount,omitempty"`
}

func (x *CpuModel) Reset() {
	*x = CpuModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuModel) ProtoMessage() {}

func (x *CpuModel) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuModel.ProtoReflect.Descriptor instead.
func (*CpuModel) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{3}
}

func (x *CpuModel) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *CpuModel) GetPhysicalCount() int32 {
	if x != nil {
		return x.PhysicalCount
	}
	return 0
}

func (x *CpuModel) GetLogicalCount() int32 {
	if x != nil {
		return x.LogicalCount
	}
	return 0
}

type MemoryStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     string `protobuf:"bytes,1,opt,name=total,proto3" json:"total,omitempty"`
	Free      string `protobuf:"bytes,2,opt,name=free,proto3" json:"free,omitempty"`
	Used      string `protobuf:"bytes,3,opt,name=used,proto3" json:"used,omitempty"`
	Available string `protobuf:"bytes,4,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *MemoryStat) Reset() {
	*x = MemoryStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryStat) ProtoMessage() {}

func (x *MemoryStat) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryStat.ProtoReflect.Descriptor instead.
func (*MemoryStat) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{4}
}

func (x *MemoryStat) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *MemoryStat) GetFree() string {
	if x != nil {
		return x.Free
	}
	return ""
}

func (x *MemoryStat) GetUsed() string {
	if x != nil {
		return x.Used
	}
	return ""
}

func (x *MemoryStat) GetAvailable() string {
	if x != nil {
		return x.Available
	}
	return ""
}

type LoadAverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	One     float64 `protobuf:"fixed64,1,opt,name=one,proto3" json:"one,omitempty"`
	Five    float64 `protobuf:"fixed64,2,opt,name=five,proto3" json:"five,omitempty"`
	Fifteen float64 `protobuf:"fixed64,3,opt,name=fifteen,proto3" json:"fifteen,omitempty"`
}

func (x *LoadAverage) Reset() {
	*x = LoadAverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadAverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadAverage) ProtoMessage() {}

func (x *LoadAverage) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadAverage.ProtoReflect.Descriptor instead.
func (*LoadAverage) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{5}
}

func (x *LoadAverage) GetOne() float64 {
	if x != nil {
		return x.One
	}
	return 0
}

func (x *LoadAverage) GetFive() float64 {
	if x != nil {
		return x.Five
	}
	return 0
}

func (x *LoadAverage) GetFifteen() float64 {
	if x != nil {
		return x.Fifteen
	}
	return 0
}

type UserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserListRequest) Reset() {
	*x = UserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListRequest) ProtoMessage() {}

func (x *UserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListRequest.ProtoReflect.Descriptor instead.
func (*UserListRequest) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{6}
}

type UserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	List    []*UserInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResponse.ProtoReflect.Descriptor instead.
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{7}
}

func (x *UserListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserListResponse) GetList() []*UserInfo {
	if x != nil {
		return x.List
	}
	return nil
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName      string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	UserID        string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	UserType      string `protobuf:"bytes,3,opt,name=UserType,proto3" json:"UserType,omitempty"`
	GroupID       string `protobuf:"bytes,4,opt,name=groupID,proto3" json:"groupID,omitempty"`
	GroupName     string `protobuf:"bytes,5,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Comment       string `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
	HomeDir       string `protobuf:"bytes,7,opt,name=homeDir,proto3" json:"homeDir,omitempty"`
	Shell         string `protobuf:"bytes,8,opt,name=shell,proto3" json:"shell,omitempty"`
	LastLoginTime string `protobuf:"bytes,9,opt,name=lastLoginTime,proto3" json:"lastLoginTime,omitempty"`
	Status        string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{8}
}

func (x *UserInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserInfo) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

func (x *UserInfo) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *UserInfo) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *UserInfo) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *UserInfo) GetHomeDir() string {
	if x != nil {
		return x.HomeDir
	}
	return ""
}

func (x *UserInfo) GetShell() string {
	if x != nil {
		return x.Shell
	}
	return ""
}

func (x *UserInfo) GetLastLoginTime() string {
	if x != nil {
		return x.LastLoginTime
	}
	return ""
}

func (x *UserInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_system_proto protoreflect.FileDescriptor

var file_system_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x83, 0x04, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x70, 0x75,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x70, 0x75, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x08, 0x63, 0x70, 0x75, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6b, 0x65, 0x72, 0x6e,
	0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x72, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x3d,
	0x0a, 0x0d, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0d,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x37, 0x0a,
	0x0a, 0x73, 0x77, 0x61, 0x70, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0a, 0x73, 0x77, 0x61, 0x70,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x5d, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x72, 0x0a, 0x08, 0x43, 0x70, 0x75, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x6f, 0x67,
	0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x0a, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x65,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x22, 0x4d, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x66, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x66, 0x74,
	0x65, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x66, 0x69, 0x66, 0x74, 0x65,
	0x65, 0x6e, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x57, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x9a,
	0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x68, 0x6f, 0x6d, 0x65, 0x44, 0x69, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x68, 0x6f, 0x6d, 0x65, 0x44, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x12, 0x24, 0x0a,
	0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xae, 0x01, 0x0a, 0x0d,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_proto_rawDescOnce sync.Once
	file_system_proto_rawDescData = file_system_proto_rawDesc
)

func file_system_proto_rawDescGZIP() []byte {
	file_system_proto_rawDescOnce.Do(func() {
		file_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_proto_rawDescData)
	})
	return file_system_proto_rawDescData
}

var file_system_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_system_proto_goTypes = []any{
	(*GetSysInfoRequest)(nil),  // 0: agent_proto.GetSysInfoRequest
	(*GetSysInfoResponse)(nil), // 1: agent_proto.GetSysInfoResponse
	(*PlatformModel)(nil),      // 2: agent_proto.PlatformModel
	(*CpuModel)(nil),           // 3: agent_proto.CpuModel
	(*MemoryStat)(nil),         // 4: agent_proto.MemoryStat
	(*LoadAverage)(nil),        // 5: agent_proto.LoadAverage
	(*UserListRequest)(nil),    // 6: agent_proto.UserListRequest
	(*UserListResponse)(nil),   // 7: agent_proto.UserListResponse
	(*UserInfo)(nil),           // 8: agent_proto.UserInfo
}
var file_system_proto_depIdxs = []int32{
	2, // 0: agent_proto.GetSysInfoResponse.platform:type_name -> agent_proto.PlatformModel
	3, // 1: agent_proto.GetSysInfoResponse.cpuModel:type_name -> agent_proto.CpuModel
	5, // 2: agent_proto.GetSysInfoResponse.loadAverage:type_name -> agent_proto.LoadAverage
	4, // 3: agent_proto.GetSysInfoResponse.virtualMemory:type_name -> agent_proto.MemoryStat
	4, // 4: agent_proto.GetSysInfoResponse.swapMemory:type_name -> agent_proto.MemoryStat
	8, // 5: agent_proto.UserListResponse.list:type_name -> agent_proto.UserInfo
	0, // 6: agent_proto.SystemService.GetSysInfo:input_type -> agent_proto.GetSysInfoRequest
	6, // 7: agent_proto.SystemService.GetUserList:input_type -> agent_proto.UserListRequest
	1, // 8: agent_proto.SystemService.GetSysInfo:output_type -> agent_proto.GetSysInfoResponse
	7, // 9: agent_proto.SystemService.GetUserList:output_type -> agent_proto.UserListResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_system_proto_init() }
func file_system_proto_init() {
	if File_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetSysInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetSysInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PlatformModel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CpuModel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MemoryStat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LoadAverage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UserListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UserListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UserInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_proto_goTypes,
		DependencyIndexes: file_system_proto_depIdxs,
		MessageInfos:      file_system_proto_msgTypes,
	}.Build()
	File_system_proto = out.File
	file_system_proto_rawDesc = nil
	file_system_proto_goTypes = nil
	file_system_proto_depIdxs = nil
}
