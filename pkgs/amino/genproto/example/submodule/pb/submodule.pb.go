// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: submodule.proto

package pb

import (
	pb "github.com/tendermint/tendermint2/pkgs/amino/genproto/example/submodule2/pb"
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

// messages
type StructSM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldA int64         `protobuf:"zigzag64,1,opt,name=FieldA,proto3" json:"FieldA,omitempty"`
	FieldB string        `protobuf:"bytes,2,opt,name=FieldB,proto3" json:"FieldB,omitempty"`
	FieldC *pb.StructSM2 `protobuf:"bytes,3,opt,name=FieldC,proto3" json:"FieldC,omitempty"`
}

func (x *StructSM) Reset() {
	*x = StructSM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructSM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructSM) ProtoMessage() {}

func (x *StructSM) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructSM.ProtoReflect.Descriptor instead.
func (*StructSM) Descriptor() ([]byte, []int) {
	return file_submodule_proto_rawDescGZIP(), []int{0}
}

func (x *StructSM) GetFieldA() int64 {
	if x != nil {
		return x.FieldA
	}
	return 0
}

func (x *StructSM) GetFieldB() string {
	if x != nil {
		return x.FieldB
	}
	return ""
}

func (x *StructSM) GetFieldC() *pb.StructSM2 {
	if x != nil {
		return x.FieldC
	}
	return nil
}

var File_submodule_proto protoreflect.FileDescriptor

var file_submodule_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x1a, 0x4e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6e, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x67, 0x6e, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x73, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x32, 0x2f, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x08,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x53, 0x4d, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x41,
	0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x12, 0x2d, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x43, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x32, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x53, 0x4d, 0x32, 0x52,
	0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6e, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6e,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x73, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x75,
	0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_submodule_proto_rawDescOnce sync.Once
	file_submodule_proto_rawDescData = file_submodule_proto_rawDesc
)

func file_submodule_proto_rawDescGZIP() []byte {
	file_submodule_proto_rawDescOnce.Do(func() {
		file_submodule_proto_rawDescData = protoimpl.X.CompressGZIP(file_submodule_proto_rawDescData)
	})
	return file_submodule_proto_rawDescData
}

var file_submodule_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_submodule_proto_goTypes = []interface{}{
	(*StructSM)(nil),     // 0: submodule.StructSM
	(*pb.StructSM2)(nil), // 1: submodule2.StructSM2
}
var file_submodule_proto_depIdxs = []int32{
	1, // 0: submodule.StructSM.FieldC:type_name -> submodule2.StructSM2
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_submodule_proto_init() }
func file_submodule_proto_init() {
	if File_submodule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_submodule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructSM); i {
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
			RawDescriptor: file_submodule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_submodule_proto_goTypes,
		DependencyIndexes: file_submodule_proto_depIdxs,
		MessageInfos:      file_submodule_proto_msgTypes,
	}.Build()
	File_submodule_proto = out.File
	file_submodule_proto_rawDesc = nil
	file_submodule_proto_goTypes = nil
	file_submodule_proto_depIdxs = nil
}
