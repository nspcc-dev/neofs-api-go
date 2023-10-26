// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: lock/grpc/types.proto

package lock

import (
	grpc "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
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

// Lock objects protects a list of objects from being deleted. The lifetime of a
// lock object is limited similar to regular objects in
// `__NEOFS__EXPIRATION_EPOCH` attribute. Lock object MUST have expiration epoch.
// It is impossible to delete a lock object via ObjectService.Delete RPC call.
// Deleting a container containing lock/locked objects results in their removal
// too, regardless of their expiration epochs.
type Lock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of objects to lock. Must not be empty or carry empty IDs.
	// All members must be of the `REGULAR` type.
	Members []*grpc.ObjectID `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *Lock) Reset() {
	*x = Lock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lock_grpc_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lock) ProtoMessage() {}

func (x *Lock) ProtoReflect() protoreflect.Message {
	mi := &file_lock_grpc_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lock.ProtoReflect.Descriptor instead.
func (*Lock) Descriptor() ([]byte, []int) {
	return file_lock_grpc_types_proto_rawDescGZIP(), []int{0}
}

func (x *Lock) GetMembers() []*grpc.ObjectID {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_lock_grpc_types_proto protoreflect.FileDescriptor

var file_lock_grpc_types_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x15, 0x72, 0x65, 0x66, 0x73, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a,
	0x0a, 0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x66, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x44, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x42, 0x50, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x73, 0x70, 0x63, 0x63, 0x2d, 0x64,
	0x65, 0x76, 0x2f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x76, 0x32, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x6c, 0x6f, 0x63,
	0x6b, 0xaa, 0x02, 0x18, 0x4e, 0x65, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lock_grpc_types_proto_rawDescOnce sync.Once
	file_lock_grpc_types_proto_rawDescData = file_lock_grpc_types_proto_rawDesc
)

func file_lock_grpc_types_proto_rawDescGZIP() []byte {
	file_lock_grpc_types_proto_rawDescOnce.Do(func() {
		file_lock_grpc_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_lock_grpc_types_proto_rawDescData)
	})
	return file_lock_grpc_types_proto_rawDescData
}

var file_lock_grpc_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lock_grpc_types_proto_goTypes = []interface{}{
	(*Lock)(nil),          // 0: neo.fs.v2.lock.Lock
	(*grpc.ObjectID)(nil), // 1: neo.fs.v2.refs.ObjectID
}
var file_lock_grpc_types_proto_depIdxs = []int32{
	1, // 0: neo.fs.v2.lock.Lock.members:type_name -> neo.fs.v2.refs.ObjectID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_lock_grpc_types_proto_init() }
func file_lock_grpc_types_proto_init() {
	if File_lock_grpc_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lock_grpc_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lock); i {
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
			RawDescriptor: file_lock_grpc_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lock_grpc_types_proto_goTypes,
		DependencyIndexes: file_lock_grpc_types_proto_depIdxs,
		MessageInfos:      file_lock_grpc_types_proto_msgTypes,
	}.Build()
	File_lock_grpc_types_proto = out.File
	file_lock_grpc_types_proto_rawDesc = nil
	file_lock_grpc_types_proto_goTypes = nil
	file_lock_grpc_types_proto_depIdxs = nil
}
