// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: reputation/grpc/service.proto

package reputation

import (
	grpc "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
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

// Announce node's local trust information.
type AnnounceLocalTrustRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of the request message.
	Body *AnnounceLocalTrustRequest_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries request meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *grpc.RequestMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries request verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *grpc.RequestVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *AnnounceLocalTrustRequest) Reset() {
	*x = AnnounceLocalTrustRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceLocalTrustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceLocalTrustRequest) ProtoMessage() {}

func (x *AnnounceLocalTrustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceLocalTrustRequest.ProtoReflect.Descriptor instead.
func (*AnnounceLocalTrustRequest) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *AnnounceLocalTrustRequest) GetBody() *AnnounceLocalTrustRequest_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *AnnounceLocalTrustRequest) GetMetaHeader() *grpc.RequestMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *AnnounceLocalTrustRequest) GetVerifyHeader() *grpc.RequestVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// Node's local trust information announcement response.
type AnnounceLocalTrustResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of the response message.
	Body *AnnounceLocalTrustResponse_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries response meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *grpc.ResponseMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries response verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *grpc.ResponseVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *AnnounceLocalTrustResponse) Reset() {
	*x = AnnounceLocalTrustResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceLocalTrustResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceLocalTrustResponse) ProtoMessage() {}

func (x *AnnounceLocalTrustResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceLocalTrustResponse.ProtoReflect.Descriptor instead.
func (*AnnounceLocalTrustResponse) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *AnnounceLocalTrustResponse) GetBody() *AnnounceLocalTrustResponse_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *AnnounceLocalTrustResponse) GetMetaHeader() *grpc.ResponseMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *AnnounceLocalTrustResponse) GetVerifyHeader() *grpc.ResponseVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// Announce intermediate global trust information.
type AnnounceIntermediateResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of the request message.
	Body *AnnounceIntermediateResultRequest_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries request meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *grpc.RequestMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries request verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *grpc.RequestVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *AnnounceIntermediateResultRequest) Reset() {
	*x = AnnounceIntermediateResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceIntermediateResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceIntermediateResultRequest) ProtoMessage() {}

func (x *AnnounceIntermediateResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceIntermediateResultRequest.ProtoReflect.Descriptor instead.
func (*AnnounceIntermediateResultRequest) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *AnnounceIntermediateResultRequest) GetBody() *AnnounceIntermediateResultRequest_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *AnnounceIntermediateResultRequest) GetMetaHeader() *grpc.RequestMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *AnnounceIntermediateResultRequest) GetVerifyHeader() *grpc.RequestVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// Intermediate global trust information announcement response.
type AnnounceIntermediateResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of the response message.
	Body *AnnounceIntermediateResultResponse_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries response meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *grpc.ResponseMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries response verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *grpc.ResponseVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *AnnounceIntermediateResultResponse) Reset() {
	*x = AnnounceIntermediateResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceIntermediateResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceIntermediateResultResponse) ProtoMessage() {}

func (x *AnnounceIntermediateResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceIntermediateResultResponse.ProtoReflect.Descriptor instead.
func (*AnnounceIntermediateResultResponse) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *AnnounceIntermediateResultResponse) GetBody() *AnnounceIntermediateResultResponse_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *AnnounceIntermediateResultResponse) GetMetaHeader() *grpc.ResponseMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *AnnounceIntermediateResultResponse) GetVerifyHeader() *grpc.ResponseVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// Announce node's local trust information.
type AnnounceLocalTrustRequest_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Trust assessment Epoch number
	Epoch uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// List of normalized local trust values to other NeoFS nodes. The value
	// is calculated according to EigenTrust++ algorithm and must be a
	// floating point number in [0;1] range.
	Trusts []*Trust `protobuf:"bytes,2,rep,name=trusts,proto3" json:"trusts,omitempty"`
}

func (x *AnnounceLocalTrustRequest_Body) Reset() {
	*x = AnnounceLocalTrustRequest_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceLocalTrustRequest_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceLocalTrustRequest_Body) ProtoMessage() {}

func (x *AnnounceLocalTrustRequest_Body) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceLocalTrustRequest_Body.ProtoReflect.Descriptor instead.
func (*AnnounceLocalTrustRequest_Body) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AnnounceLocalTrustRequest_Body) GetEpoch() uint64 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *AnnounceLocalTrustRequest_Body) GetTrusts() []*Trust {
	if x != nil {
		return x.Trusts
	}
	return nil
}

// Response to the node's local trust information announcement has an empty body
// because the trust exchange operation is asynchronous. If Trust information
// does not pass sanity checks, it is silently ignored.
type AnnounceLocalTrustResponse_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnnounceLocalTrustResponse_Body) Reset() {
	*x = AnnounceLocalTrustResponse_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceLocalTrustResponse_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceLocalTrustResponse_Body) ProtoMessage() {}

func (x *AnnounceLocalTrustResponse_Body) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceLocalTrustResponse_Body.ProtoReflect.Descriptor instead.
func (*AnnounceLocalTrustResponse_Body) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_service_proto_rawDescGZIP(), []int{1, 0}
}

// Announce intermediate global trust information.
type AnnounceIntermediateResultRequest_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Iteration execution Epoch number
	Epoch uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// Iteration sequence number
	Iteration uint32 `protobuf:"varint,2,opt,name=iteration,proto3" json:"iteration,omitempty"`
	// Current global trust value calculated at the specified iteration
	Trust *PeerToPeerTrust `protobuf:"bytes,3,opt,name=trust,proto3" json:"trust,omitempty"`
}

func (x *AnnounceIntermediateResultRequest_Body) Reset() {
	*x = AnnounceIntermediateResultRequest_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceIntermediateResultRequest_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceIntermediateResultRequest_Body) ProtoMessage() {}

func (x *AnnounceIntermediateResultRequest_Body) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceIntermediateResultRequest_Body.ProtoReflect.Descriptor instead.
func (*AnnounceIntermediateResultRequest_Body) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *AnnounceIntermediateResultRequest_Body) GetEpoch() uint64 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *AnnounceIntermediateResultRequest_Body) GetIteration() uint32 {
	if x != nil {
		return x.Iteration
	}
	return 0
}

func (x *AnnounceIntermediateResultRequest_Body) GetTrust() *PeerToPeerTrust {
	if x != nil {
		return x.Trust
	}
	return nil
}

// Response to the node's intermediate global trust information announcement has
// an empty body because the trust exchange operation is asynchronous. If
// Trust information does not pass sanity checks, it is silently ignored.
type AnnounceIntermediateResultResponse_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnnounceIntermediateResultResponse_Body) Reset() {
	*x = AnnounceIntermediateResultResponse_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reputation_grpc_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnounceIntermediateResultResponse_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnounceIntermediateResultResponse_Body) ProtoMessage() {}

func (x *AnnounceIntermediateResultResponse_Body) ProtoReflect() protoreflect.Message {
	mi := &file_reputation_grpc_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnounceIntermediateResultResponse_Body.ProtoReflect.Descriptor instead.
func (*AnnounceIntermediateResultResponse_Body) Descriptor() ([]byte, []int) {
	return file_reputation_grpc_service_proto_rawDescGZIP(), []int{3, 0}
}

var File_reputation_grpc_service_proto protoreflect.FileDescriptor

var file_reputation_grpc_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x02, 0x0a,
	0x19, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x72,
	0x75, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x45, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x6f, 0x2e,
	0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x0a, 0x6d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x0d, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x51,
	0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x33, 0x0a, 0x06,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e,
	0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x06, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x73, 0x22, 0x8b, 0x02, 0x0a, 0x1a, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x49, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x46, 0x0a, 0x0b, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x65, 0x6f,
	0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x06, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22,
	0x88, 0x03, 0x0a, 0x21, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x45, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e,
	0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x51,
	0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x1a, 0x77, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a,
	0x05, 0x74, 0x72, 0x75, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e,
	0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x65, 0x65, 0x72, 0x54, 0x72,
	0x75, 0x73, 0x74, 0x52, 0x05, 0x74, 0x72, 0x75, 0x73, 0x74, 0x22, 0x9b, 0x02, 0x0a, 0x22, 0x41,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x51, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x46, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x65, 0x6f, 0x2e,
	0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0d,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x1a, 0x06, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x32, 0x9e, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x70,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77,
	0x0a, 0x12, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54,
	0x72, 0x75, 0x73, 0x74, 0x12, 0x2f, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x1a, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x38, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x70, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x62, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x73, 0x70, 0x63, 0x63, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x76,
	0x32, 0x2f, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x3b, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xaa, 0x02, 0x1e, 0x4e,
	0x65, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41,
	0x50, 0x49, 0x2e, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reputation_grpc_service_proto_rawDescOnce sync.Once
	file_reputation_grpc_service_proto_rawDescData = file_reputation_grpc_service_proto_rawDesc
)

func file_reputation_grpc_service_proto_rawDescGZIP() []byte {
	file_reputation_grpc_service_proto_rawDescOnce.Do(func() {
		file_reputation_grpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_reputation_grpc_service_proto_rawDescData)
	})
	return file_reputation_grpc_service_proto_rawDescData
}

var file_reputation_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_reputation_grpc_service_proto_goTypes = []interface{}{
	(*AnnounceLocalTrustRequest)(nil),               // 0: neo.fs.v2.reputation.AnnounceLocalTrustRequest
	(*AnnounceLocalTrustResponse)(nil),              // 1: neo.fs.v2.reputation.AnnounceLocalTrustResponse
	(*AnnounceIntermediateResultRequest)(nil),       // 2: neo.fs.v2.reputation.AnnounceIntermediateResultRequest
	(*AnnounceIntermediateResultResponse)(nil),      // 3: neo.fs.v2.reputation.AnnounceIntermediateResultResponse
	(*AnnounceLocalTrustRequest_Body)(nil),          // 4: neo.fs.v2.reputation.AnnounceLocalTrustRequest.Body
	(*AnnounceLocalTrustResponse_Body)(nil),         // 5: neo.fs.v2.reputation.AnnounceLocalTrustResponse.Body
	(*AnnounceIntermediateResultRequest_Body)(nil),  // 6: neo.fs.v2.reputation.AnnounceIntermediateResultRequest.Body
	(*AnnounceIntermediateResultResponse_Body)(nil), // 7: neo.fs.v2.reputation.AnnounceIntermediateResultResponse.Body
	(*grpc.RequestMetaHeader)(nil),                  // 8: neo.fs.v2.session.RequestMetaHeader
	(*grpc.RequestVerificationHeader)(nil),          // 9: neo.fs.v2.session.RequestVerificationHeader
	(*grpc.ResponseMetaHeader)(nil),                 // 10: neo.fs.v2.session.ResponseMetaHeader
	(*grpc.ResponseVerificationHeader)(nil),         // 11: neo.fs.v2.session.ResponseVerificationHeader
	(*Trust)(nil),                                   // 12: neo.fs.v2.reputation.Trust
	(*PeerToPeerTrust)(nil),                         // 13: neo.fs.v2.reputation.PeerToPeerTrust
}
var file_reputation_grpc_service_proto_depIdxs = []int32{
	4,  // 0: neo.fs.v2.reputation.AnnounceLocalTrustRequest.body:type_name -> neo.fs.v2.reputation.AnnounceLocalTrustRequest.Body
	8,  // 1: neo.fs.v2.reputation.AnnounceLocalTrustRequest.meta_header:type_name -> neo.fs.v2.session.RequestMetaHeader
	9,  // 2: neo.fs.v2.reputation.AnnounceLocalTrustRequest.verify_header:type_name -> neo.fs.v2.session.RequestVerificationHeader
	5,  // 3: neo.fs.v2.reputation.AnnounceLocalTrustResponse.body:type_name -> neo.fs.v2.reputation.AnnounceLocalTrustResponse.Body
	10, // 4: neo.fs.v2.reputation.AnnounceLocalTrustResponse.meta_header:type_name -> neo.fs.v2.session.ResponseMetaHeader
	11, // 5: neo.fs.v2.reputation.AnnounceLocalTrustResponse.verify_header:type_name -> neo.fs.v2.session.ResponseVerificationHeader
	6,  // 6: neo.fs.v2.reputation.AnnounceIntermediateResultRequest.body:type_name -> neo.fs.v2.reputation.AnnounceIntermediateResultRequest.Body
	8,  // 7: neo.fs.v2.reputation.AnnounceIntermediateResultRequest.meta_header:type_name -> neo.fs.v2.session.RequestMetaHeader
	9,  // 8: neo.fs.v2.reputation.AnnounceIntermediateResultRequest.verify_header:type_name -> neo.fs.v2.session.RequestVerificationHeader
	7,  // 9: neo.fs.v2.reputation.AnnounceIntermediateResultResponse.body:type_name -> neo.fs.v2.reputation.AnnounceIntermediateResultResponse.Body
	10, // 10: neo.fs.v2.reputation.AnnounceIntermediateResultResponse.meta_header:type_name -> neo.fs.v2.session.ResponseMetaHeader
	11, // 11: neo.fs.v2.reputation.AnnounceIntermediateResultResponse.verify_header:type_name -> neo.fs.v2.session.ResponseVerificationHeader
	12, // 12: neo.fs.v2.reputation.AnnounceLocalTrustRequest.Body.trusts:type_name -> neo.fs.v2.reputation.Trust
	13, // 13: neo.fs.v2.reputation.AnnounceIntermediateResultRequest.Body.trust:type_name -> neo.fs.v2.reputation.PeerToPeerTrust
	0,  // 14: neo.fs.v2.reputation.ReputationService.AnnounceLocalTrust:input_type -> neo.fs.v2.reputation.AnnounceLocalTrustRequest
	2,  // 15: neo.fs.v2.reputation.ReputationService.AnnounceIntermediateResult:input_type -> neo.fs.v2.reputation.AnnounceIntermediateResultRequest
	1,  // 16: neo.fs.v2.reputation.ReputationService.AnnounceLocalTrust:output_type -> neo.fs.v2.reputation.AnnounceLocalTrustResponse
	3,  // 17: neo.fs.v2.reputation.ReputationService.AnnounceIntermediateResult:output_type -> neo.fs.v2.reputation.AnnounceIntermediateResultResponse
	16, // [16:18] is the sub-list for method output_type
	14, // [14:16] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_reputation_grpc_service_proto_init() }
func file_reputation_grpc_service_proto_init() {
	if File_reputation_grpc_service_proto != nil {
		return
	}
	file_reputation_grpc_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_reputation_grpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceLocalTrustRequest); i {
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
		file_reputation_grpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceLocalTrustResponse); i {
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
		file_reputation_grpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceIntermediateResultRequest); i {
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
		file_reputation_grpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceIntermediateResultResponse); i {
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
		file_reputation_grpc_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceLocalTrustRequest_Body); i {
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
		file_reputation_grpc_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceLocalTrustResponse_Body); i {
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
		file_reputation_grpc_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceIntermediateResultRequest_Body); i {
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
		file_reputation_grpc_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnounceIntermediateResultResponse_Body); i {
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
			RawDescriptor: file_reputation_grpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reputation_grpc_service_proto_goTypes,
		DependencyIndexes: file_reputation_grpc_service_proto_depIdxs,
		MessageInfos:      file_reputation_grpc_service_proto_msgTypes,
	}.Build()
	File_reputation_grpc_service_proto = out.File
	file_reputation_grpc_service_proto_rawDesc = nil
	file_reputation_grpc_service_proto_goTypes = nil
	file_reputation_grpc_service_proto_depIdxs = nil
}
