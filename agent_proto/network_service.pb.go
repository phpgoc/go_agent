// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: network_service.proto

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

type NetworkInterfaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NetworkInterfaceRequest) Reset() {
	*x = NetworkInterfaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInterfaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterfaceRequest) ProtoMessage() {}

func (x *NetworkInterfaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterfaceRequest.ProtoReflect.Descriptor instead.
func (*NetworkInterfaceRequest) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{0}
}

type NetworkInterfaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message           string              `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	NetworkInterfaces []*NetworkInterface `protobuf:"bytes,2,rep,name=network_interfaces,json=networkInterfaces,proto3" json:"network_interfaces,omitempty"`
}

func (x *NetworkInterfaceResponse) Reset() {
	*x = NetworkInterfaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInterfaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterfaceResponse) ProtoMessage() {}

func (x *NetworkInterfaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterfaceResponse.ProtoReflect.Descriptor instead.
func (*NetworkInterfaceResponse) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkInterfaceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *NetworkInterfaceResponse) GetNetworkInterfaces() []*NetworkInterface {
	if x != nil {
		return x.NetworkInterfaces
	}
	return nil
}

type NetworkInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mac   string   `protobuf:"bytes,2,opt,name=mac,proto3" json:"mac,omitempty"`
	Ipv4  []string `protobuf:"bytes,3,rep,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6  []string `protobuf:"bytes,4,rep,name=ipv6,proto3" json:"ipv6,omitempty"`
	Flags string   `protobuf:"bytes,7,opt,name=flags,proto3" json:"flags,omitempty"`
}

func (x *NetworkInterface) Reset() {
	*x = NetworkInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterface) ProtoMessage() {}

func (x *NetworkInterface) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterface.ProtoReflect.Descriptor instead.
func (*NetworkInterface) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkInterface) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NetworkInterface) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *NetworkInterface) GetIpv4() []string {
	if x != nil {
		return x.Ipv4
	}
	return nil
}

func (x *NetworkInterface) GetIpv6() []string {
	if x != nil {
		return x.Ipv6
	}
	return nil
}

func (x *NetworkInterface) GetFlags() string {
	if x != nil {
		return x.Flags
	}
	return ""
}

var File_network_service_proto protoreflect.FileDescriptor

var file_network_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x82, 0x01, 0x0a, 0x18, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x52, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x73, 0x22, 0x76, 0x0a, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x61, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70,
	0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x36, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x70, 0x76, 0x36, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x32, 0x6f, 0x0a, 0x07,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x64, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x24,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_service_proto_rawDescOnce sync.Once
	file_network_service_proto_rawDescData = file_network_service_proto_rawDesc
)

func file_network_service_proto_rawDescGZIP() []byte {
	file_network_service_proto_rawDescOnce.Do(func() {
		file_network_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_service_proto_rawDescData)
	})
	return file_network_service_proto_rawDescData
}

var file_network_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_network_service_proto_goTypes = []any{
	(*NetworkInterfaceRequest)(nil),  // 0: agent_proto.NetworkInterfaceRequest
	(*NetworkInterfaceResponse)(nil), // 1: agent_proto.NetworkInterfaceResponse
	(*NetworkInterface)(nil),         // 2: agent_proto.NetworkInterface
}
var file_network_service_proto_depIdxs = []int32{
	2, // 0: agent_proto.NetworkInterfaceResponse.network_interfaces:type_name -> agent_proto.NetworkInterface
	0, // 1: agent_proto.Network.GetNetworkInterface:input_type -> agent_proto.NetworkInterfaceRequest
	1, // 2: agent_proto.Network.GetNetworkInterface:output_type -> agent_proto.NetworkInterfaceResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_network_service_proto_init() }
func file_network_service_proto_init() {
	if File_network_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NetworkInterfaceRequest); i {
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
		file_network_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NetworkInterfaceResponse); i {
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
		file_network_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*NetworkInterface); i {
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
			RawDescriptor: file_network_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_service_proto_goTypes,
		DependencyIndexes: file_network_service_proto_depIdxs,
		MessageInfos:      file_network_service_proto_msgTypes,
	}.Build()
	File_network_service_proto = out.File
	file_network_service_proto_rawDesc = nil
	file_network_service_proto_goTypes = nil
	file_network_service_proto_depIdxs = nil
}
