// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.6
// source: reputation/grpc/types.proto

package reputation

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

// NeoFS unique peer identifier is a 33 byte long compressed public key of the
// node, the same as the one stored in the network map.
//
// String presentation is a
// [base58](https://tools.ietf.org/html/draft-msporny-base58-02) encoded string.
//
// JSON value will be data encoded as a string using standard base64
// encoding with paddings. Either
// [standard](https://tools.ietf.org/html/rfc4648#section-4) or
// [URL-safe](https://tools.ietf.org/html/rfc4648#section-5) base64 encoding
// with/without paddings are accepted.
type PeerID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Peer node's public key
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *PeerID) Reset() {
	*x = PeerID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerID) ProtoMessage() {}

func (x *PeerID) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerID.ProtoReflect.Descriptor instead.
func (*PeerID) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_types_proto_rawDescGZIP(), []int{0}
}

func (x *PeerID) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// Trust level to a NeoFS network peer.
type Trust struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of the trusted peer
	Peer *PeerID `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	// Trust level in [0:1] range
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Trust) Reset() {
	*x = Trust{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trust) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trust) ProtoMessage() {}

func (x *Trust) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trust.ProtoReflect.Descriptor instead.
func (*Trust) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_types_proto_rawDescGZIP(), []int{1}
}

func (x *Trust) GetPeer() *PeerID {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *Trust) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Trust level of a peer to a peer.
type PeerToPeerTrust struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of the trusting peer
	TrustingPeer *PeerID `protobuf:"bytes,1,opt,name=trusting_peer,json=trustingPeer,proto3" json:"trusting_peer,omitempty"`
	// Trust level
	Trust *Trust `protobuf:"bytes,2,opt,name=trust,proto3" json:"trust,omitempty"`
}

func (x *PeerToPeerTrust) Reset() {
	*x = PeerToPeerTrust{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerToPeerTrust) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerToPeerTrust) ProtoMessage() {}

func (x *PeerToPeerTrust) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerToPeerTrust.ProtoReflect.Descriptor instead.
func (*PeerToPeerTrust) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_types_proto_rawDescGZIP(), []int{2}
}

func (x *PeerToPeerTrust) GetTrustingPeer() *PeerID {
	if x != nil {
		return x.TrustingPeer
	}
	return nil
}

func (x *PeerToPeerTrust) GetTrust() *Trust {
	if x != nil {
		return x.Trust
	}
	return nil
}

// Global trust level to NeoFS node.
type GlobalTrust struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message format version. Effectively, the version of API library used to create
	// the message.
	Version *grpc.Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Message body
	Body *GlobalTrust_Body `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// Signature of the binary `body` field by the manager.
	Signature *grpc.Signature `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *GlobalTrust) Reset() {
	*x = GlobalTrust{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalTrust) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalTrust) ProtoMessage() {}

func (x *GlobalTrust) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalTrust.ProtoReflect.Descriptor instead.
func (*GlobalTrust) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_types_proto_rawDescGZIP(), []int{3}
}

func (x *GlobalTrust) GetVersion() *grpc.Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *GlobalTrust) GetBody() *GlobalTrust_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *GlobalTrust) GetSignature() *grpc.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// Message body structure.
type GlobalTrust_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node manager ID
	Manager *PeerID `protobuf:"bytes,1,opt,name=manager,proto3" json:"manager,omitempty"`
	// Global trust level
	Trust *Trust `protobuf:"bytes,2,opt,name=trust,proto3" json:"trust,omitempty"`
}

func (x *GlobalTrust_Body) Reset() {
	*x = GlobalTrust_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalTrust_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalTrust_Body) ProtoMessage() {}

func (x *GlobalTrust_Body) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalTrust_Body.ProtoReflect.Descriptor instead.
func (*GlobalTrust_Body) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_types_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GlobalTrust_Body) GetManager() *PeerID {
	if x != nil {
		return x.Manager
	}
	return nil
}

func (x *GlobalTrust_Body) GetTrust() *Trust {
	if x != nil {
		return x.Trust
	}
	return nil
}

var File_reputation_grpc_types_proto protoreflect.FileDescriptor

var file_reputation_grpc_types_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e,
	0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x15, 0x72, 0x65, 0x66, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x06, 0x50, 0x65,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x22, 0x4f, 0x0a, 0x05, 0x54, 0x72, 0x75, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04,
	0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x6f,
	0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x50, 0x65, 0x65, 0x72, 0x54, 0x6f, 0x50,
	0x65, 0x65, 0x72, 0x54, 0x72, 0x75, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0d, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x52, 0x0c, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x05, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x6f,
	0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x05, 0x74, 0x72, 0x75, 0x73, 0x74, 0x22, 0xa8,
	0x02, 0x0a, 0x0b, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x54, 0x72, 0x75, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x66, 0x73,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x3a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x37, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x66,
	0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x71, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x36,
	0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x52, 0x07, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x05, 0x74, 0x72, 0x75, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x52, 0x05, 0x74, 0x72, 0x75, 0x73, 0x74, 0x42, 0x62, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x73, 0x70, 0x63, 0x63, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x76,
	0x32, 0x2f, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x3b, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xaa, 0x02, 0x1e, 0x4e,
	0x65, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41,
	0x50, 0x49, 0x2e, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reputation_grpc_types_proto_rawDescOnce sync.Once
	file_reputation_grpc_types_proto_rawDescData = file_reputation_grpc_types_proto_rawDesc
)

func file_reputation_grpc_types_proto_rawDescGZIP() []byte {
	file_reputation_grpc_types_proto_rawDescOnce.Do(func() {
		file_reputation_grpc_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_reputation_grpc_types_proto_rawDescData)
	})
	return file_reputation_grpc_types_proto_rawDescData
}

var file_reputation_grpc_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_reputation_grpc_types_proto_goTypes = []interface{}{
	(*PeerID)(nil),           // 0: neo.fs.v2.reputation.PeerID
	(*Trust)(nil),            // 1: neo.fs.v2.reputation.Trust
	(*PeerToPeerTrust)(nil),  // 2: neo.fs.v2.reputation.PeerToPeerTrust
	(*GlobalTrust)(nil),      // 3: neo.fs.v2.reputation.GlobalTrust
	(*GlobalTrust_Body)(nil), // 4: neo.fs.v2.reputation.GlobalTrust.Body
	(*grpc.Version)(nil),     // 5: neo.fs.v2.refs.Version
	(*grpc.Signature)(nil),   // 6: neo.fs.v2.refs.Signature
}
var file_reputation_grpc_types_proto_depIdxs = []int32{
	0, // 0: neo.fs.v2.reputation.Trust.peer:type_name -> neo.fs.v2.reputation.PeerID
	0, // 1: neo.fs.v2.reputation.PeerToPeerTrust.trusting_peer:type_name -> neo.fs.v2.reputation.PeerID
	1, // 2: neo.fs.v2.reputation.PeerToPeerTrust.trust:type_name -> neo.fs.v2.reputation.Trust
	5, // 3: neo.fs.v2.reputation.GlobalTrust.version:type_name -> neo.fs.v2.refs.Version
	4, // 4: neo.fs.v2.reputation.GlobalTrust.body:type_name -> neo.fs.v2.reputation.GlobalTrust.Body
	6, // 5: neo.fs.v2.reputation.GlobalTrust.signature:type_name -> neo.fs.v2.refs.Signature
	0, // 6: neo.fs.v2.reputation.GlobalTrust.Body.manager:type_name -> neo.fs.v2.reputation.PeerID
	1, // 7: neo.fs.v2.reputation.GlobalTrust.Body.trust:type_name -> neo.fs.v2.reputation.Trust
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_reputation_grpc_types_proto_init() }
func file_reputation_grpc_types_proto_init() {
	if File_reputation_grpc_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reputation_grpc_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerID); i {
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
		file_reputation_grpc_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trust); i {
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
		file_reputation_grpc_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerToPeerTrust); i {
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
		file_reputation_grpc_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalTrust); i {
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
		file_reputation_grpc_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalTrust_Body); i {
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
			RawDescriptor: file_reputation_grpc_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reputation_grpc_types_proto_goTypes,
		DependencyIndexes: file_reputation_grpc_types_proto_depIdxs,
		MessageInfos:      file_reputation_grpc_types_proto_msgTypes,
	}.Build()
	File_reputation_grpc_types_proto = out.File
	file_reputation_grpc_types_proto_rawDesc = nil
	file_reputation_grpc_types_proto_goTypes = nil
	file_reputation_grpc_types_proto_depIdxs = nil
}
