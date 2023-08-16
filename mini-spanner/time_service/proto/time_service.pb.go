// modified from cs426 lab1
// https://github.com/shixiao/cs426-spring2023-labs/tree/main/lab1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/time_service.proto

package proto

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTimeRequest) Reset() {
	*x = GetTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_time_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimeRequest) ProtoMessage() {}

func (x *GetTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_time_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimeRequest.ProtoReflect.Descriptor instead.
func (*GetTimeRequest) Descriptor() ([]byte, []int) {
	return file_proto_time_service_proto_rawDescGZIP(), []int{0}
}

type GetTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GetTimeResponse) Reset() {
	*x = GetTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_time_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimeResponse) ProtoMessage() {}

func (x *GetTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_time_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimeResponse.ProtoReflect.Descriptor instead.
func (*GetTimeResponse) Descriptor() ([]byte, []int) {
	return file_proto_time_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTimeResponse) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_proto_time_service_proto protoreflect.FileDescriptor

var file_proto_time_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0x57,
	0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x6d, 0x69, 0x6e, 0x69, 0x2d,
	0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_time_service_proto_rawDescOnce sync.Once
	file_proto_time_service_proto_rawDescData = file_proto_time_service_proto_rawDesc
)

func file_proto_time_service_proto_rawDescGZIP() []byte {
	file_proto_time_service_proto_rawDescOnce.Do(func() {
		file_proto_time_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_time_service_proto_rawDescData)
	})
	return file_proto_time_service_proto_rawDescData
}

var file_proto_time_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_time_service_proto_goTypes = []interface{}{
	(*GetTimeRequest)(nil),      // 0: time_service.GetTimeRequest
	(*GetTimeResponse)(nil),     // 1: time_service.GetTimeResponse
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_time_service_proto_depIdxs = []int32{
	2, // 0: time_service.GetTimeResponse.time:type_name -> google.protobuf.Timestamp
	0, // 1: time_service.TimeService.GetTime:input_type -> time_service.GetTimeRequest
	1, // 2: time_service.TimeService.GetTime:output_type -> time_service.GetTimeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_time_service_proto_init() }
func file_proto_time_service_proto_init() {
	if File_proto_time_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_time_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimeRequest); i {
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
		file_proto_time_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimeResponse); i {
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
			RawDescriptor: file_proto_time_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_time_service_proto_goTypes,
		DependencyIndexes: file_proto_time_service_proto_depIdxs,
		MessageInfos:      file_proto_time_service_proto_msgTypes,
	}.Build()
	File_proto_time_service_proto = out.File
	file_proto_time_service_proto_rawDesc = nil
	file_proto_time_service_proto_goTypes = nil
	file_proto_time_service_proto_depIdxs = nil
}