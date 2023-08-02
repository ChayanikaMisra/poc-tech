// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: hello_world.proto

package pb

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

type HelloWorldServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloWorldServiceRequest) Reset() {
	*x = HelloWorldServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldServiceRequest) ProtoMessage() {}

func (x *HelloWorldServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldServiceRequest.ProtoReflect.Descriptor instead.
func (*HelloWorldServiceRequest) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{0}
}

func (x *HelloWorldServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloWorldServiceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloWorldServiceReply) Reset() {
	*x = HelloWorldServiceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldServiceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldServiceReply) ProtoMessage() {}

func (x *HelloWorldServiceReply) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldServiceReply.ProtoReflect.Descriptor instead.
func (*HelloWorldServiceReply) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{1}
}

func (x *HelloWorldServiceReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_hello_world_proto protoreflect.FileDescriptor

var file_hello_world_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x18, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x55, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x08,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x05,
	0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_world_proto_rawDescOnce sync.Once
	file_hello_world_proto_rawDescData = file_hello_world_proto_rawDesc
)

func file_hello_world_proto_rawDescGZIP() []byte {
	file_hello_world_proto_rawDescOnce.Do(func() {
		file_hello_world_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_world_proto_rawDescData)
	})
	return file_hello_world_proto_rawDescData
}

var file_hello_world_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hello_world_proto_goTypes = []interface{}{
	(*HelloWorldServiceRequest)(nil), // 0: HelloWorldServiceRequest
	(*HelloWorldServiceReply)(nil),   // 1: HelloWorldServiceReply
}
var file_hello_world_proto_depIdxs = []int32{
	0, // 0: HelloWorldService.Greeting:input_type -> HelloWorldServiceRequest
	1, // 1: HelloWorldService.Greeting:output_type -> HelloWorldServiceReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_world_proto_init() }
func file_hello_world_proto_init() {
	if File_hello_world_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_world_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldServiceRequest); i {
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
		file_hello_world_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldServiceReply); i {
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
			RawDescriptor: file_hello_world_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_world_proto_goTypes,
		DependencyIndexes: file_hello_world_proto_depIdxs,
		MessageInfos:      file_hello_world_proto_msgTypes,
	}.Build()
	File_hello_world_proto = out.File
	file_hello_world_proto_rawDesc = nil
	file_hello_world_proto_goTypes = nil
	file_hello_world_proto_depIdxs = nil
}
