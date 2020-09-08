// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: payload.proto

package papertrailgo

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string               `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Tag      string               `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	LogTime  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=log_time,json=logTime,proto3" json:"log_time,omitempty"`
	Log      string               `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{0}
}

func (x *Payload) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Payload) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Payload) GetLogTime() *timestamp.Timestamp {
	if x != nil {
		return x.LogTime
	}
	return nil
}

func (x *Payload) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

var File_payload_proto protoreflect.FileDescriptor

var file_payload_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x73, 0x6f, 0x6c, 0x61, 0x72, 0x77, 0x69, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x6c, 0x6f, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x3b,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payload_proto_rawDescOnce sync.Once
	file_payload_proto_rawDescData = file_payload_proto_rawDesc
)

func file_payload_proto_rawDescGZIP() []byte {
	file_payload_proto_rawDescOnce.Do(func() {
		file_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_payload_proto_rawDescData)
	})
	return file_payload_proto_rawDescData
}

var file_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_payload_proto_goTypes = []interface{}{
	(*Payload)(nil),             // 0: solarwinds.papertrail.Payload
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_payload_proto_depIdxs = []int32{
	1, // 0: solarwinds.papertrail.Payload.log_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payload_proto_init() }
func file_payload_proto_init() {
	if File_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
			RawDescriptor: file_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payload_proto_goTypes,
		DependencyIndexes: file_payload_proto_depIdxs,
		MessageInfos:      file_payload_proto_msgTypes,
	}.Build()
	File_payload_proto = out.File
	file_payload_proto_rawDesc = nil
	file_payload_proto_goTypes = nil
	file_payload_proto_depIdxs = nil
}
