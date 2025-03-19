// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: hashq.proto

package gen

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sequence      int32                  `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Key           string                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Gen           string                 `protobuf:"bytes,4,opt,name=gen,proto3" json:"gen,omitempty"`
	Owner         string                 `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	mi := &file_hashq_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hashq_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_hashq_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *AddRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AddRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AddRequest) GetGen() string {
	if x != nil {
		return x.Gen
	}
	return ""
}

func (x *AddRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type AddReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Verified      bool                   `protobuf:"varint,1,opt,name=verified,proto3" json:"verified,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddReply) Reset() {
	*x = AddReply{}
	mi := &file_hashq_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReply) ProtoMessage() {}

func (x *AddReply) ProtoReflect() protoreflect.Message {
	mi := &file_hashq_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReply.ProtoReflect.Descriptor instead.
func (*AddReply) Descriptor() ([]byte, []int) {
	return file_hashq_proto_rawDescGZIP(), []int{1}
}

func (x *AddReply) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

var File_hashq_proto protoreflect.FileDescriptor

var file_hashq_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x68, 0x61, 0x73, 0x68, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68,
	0x61, 0x73, 0x68, 0x71, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x78, 0x0a, 0x0a, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x67,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x65, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x32, 0x3d, 0x0a, 0x04, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x35, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x16, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x71, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x71, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x2f,
	0x68, 0x61, 0x73, 0x68, 0x71, 0x5f, 0x6d, 0x6f, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_hashq_proto_rawDescOnce sync.Once
	file_hashq_proto_rawDescData []byte
)

func file_hashq_proto_rawDescGZIP() []byte {
	file_hashq_proto_rawDescOnce.Do(func() {
		file_hashq_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_hashq_proto_rawDesc), len(file_hashq_proto_rawDesc)))
	})
	return file_hashq_proto_rawDescData
}

var file_hashq_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hashq_proto_goTypes = []any{
	(*AddRequest)(nil), // 0: hashq_grpc.AddRequest
	(*AddReply)(nil),   // 1: hashq_grpc.AddReply
}
var file_hashq_proto_depIdxs = []int32{
	0, // 0: hashq_grpc.Hash.Add:input_type -> hashq_grpc.AddRequest
	1, // 1: hashq_grpc.Hash.Add:output_type -> hashq_grpc.AddReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hashq_proto_init() }
func file_hashq_proto_init() {
	if File_hashq_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_hashq_proto_rawDesc), len(file_hashq_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hashq_proto_goTypes,
		DependencyIndexes: file_hashq_proto_depIdxs,
		MessageInfos:      file_hashq_proto_msgTypes,
	}.Build()
	File_hashq_proto = out.File
	file_hashq_proto_goTypes = nil
	file_hashq_proto_depIdxs = nil
}
