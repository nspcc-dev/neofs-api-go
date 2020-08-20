// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v2/accounting/grpc/service.proto

package accounting

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc1 "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
	grpc2 "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Message defines the request body of Balance method.
//
// To indicate the account for which the balance is requested, it's identifier
// is used.
//
// To gain access to the requested information, the request body must be formed
// according to the requirements from the system specification.
type BalanceRequest struct {
	// Body of the balance request message.
	Body *BalanceRequest_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries request meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *grpc.RequestMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries request verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness
	// of transmission.
	VerifyHeader         *grpc.RequestVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *BalanceRequest) Reset()         { *m = BalanceRequest{} }
func (m *BalanceRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceRequest) ProtoMessage()    {}
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9dd5af2ff2bbb25, []int{0}
}
func (m *BalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceRequest.Merge(m, src)
}
func (m *BalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *BalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceRequest proto.InternalMessageInfo

func (m *BalanceRequest) GetBody() *BalanceRequest_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *BalanceRequest) GetMetaHeader() *grpc.RequestMetaHeader {
	if m != nil {
		return m.MetaHeader
	}
	return nil
}

func (m *BalanceRequest) GetVerifyHeader() *grpc.RequestVerificationHeader {
	if m != nil {
		return m.VerifyHeader
	}
	return nil
}

//Request body
type BalanceRequest_Body struct {
	// Carries user identifier in NeoFS system for which the balance
	// is requested.
	OwnerId              *grpc1.OwnerID `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BalanceRequest_Body) Reset()         { *m = BalanceRequest_Body{} }
func (m *BalanceRequest_Body) String() string { return proto.CompactTextString(m) }
func (*BalanceRequest_Body) ProtoMessage()    {}
func (*BalanceRequest_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9dd5af2ff2bbb25, []int{0, 0}
}
func (m *BalanceRequest_Body) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalanceRequest_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalanceRequest_Body.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalanceRequest_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceRequest_Body.Merge(m, src)
}
func (m *BalanceRequest_Body) XXX_Size() int {
	return m.Size()
}
func (m *BalanceRequest_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceRequest_Body.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceRequest_Body proto.InternalMessageInfo

func (m *BalanceRequest_Body) GetOwnerId() *grpc1.OwnerID {
	if m != nil {
		return m.OwnerId
	}
	return nil
}

// Message defines the response body of Balance method.
//
// The amount of funds is calculated in decimal numbers.
type BalanceResponse struct {
	// Body of the balance response message.
	Body *BalanceResponse_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries response meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *grpc.ResponseMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries response verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness
	// of transmission.
	VerifyHeader         *grpc.ResponseVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *BalanceResponse) Reset()         { *m = BalanceResponse{} }
func (m *BalanceResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceResponse) ProtoMessage()    {}
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9dd5af2ff2bbb25, []int{1}
}
func (m *BalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceResponse.Merge(m, src)
}
func (m *BalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *BalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceResponse proto.InternalMessageInfo

func (m *BalanceResponse) GetBody() *BalanceResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *BalanceResponse) GetMetaHeader() *grpc.ResponseMetaHeader {
	if m != nil {
		return m.MetaHeader
	}
	return nil
}

func (m *BalanceResponse) GetVerifyHeader() *grpc.ResponseVerificationHeader {
	if m != nil {
		return m.VerifyHeader
	}
	return nil
}

//Request body
type BalanceResponse_Body struct {
	// Carries the amount of funds on the account.
	Balance              *Decimal `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceResponse_Body) Reset()         { *m = BalanceResponse_Body{} }
func (m *BalanceResponse_Body) String() string { return proto.CompactTextString(m) }
func (*BalanceResponse_Body) ProtoMessage()    {}
func (*BalanceResponse_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9dd5af2ff2bbb25, []int{1, 0}
}
func (m *BalanceResponse_Body) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalanceResponse_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalanceResponse_Body.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalanceResponse_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceResponse_Body.Merge(m, src)
}
func (m *BalanceResponse_Body) XXX_Size() int {
	return m.Size()
}
func (m *BalanceResponse_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceResponse_Body.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceResponse_Body proto.InternalMessageInfo

func (m *BalanceResponse_Body) GetBalance() *Decimal {
	if m != nil {
		return m.Balance
	}
	return nil
}

func init() {
	proto.RegisterType((*BalanceRequest)(nil), "neo.fs.v2.accounting.BalanceRequest")
	proto.RegisterType((*BalanceRequest_Body)(nil), "neo.fs.v2.accounting.BalanceRequest.Body")
	proto.RegisterType((*BalanceResponse)(nil), "neo.fs.v2.accounting.BalanceResponse")
	proto.RegisterType((*BalanceResponse_Body)(nil), "neo.fs.v2.accounting.BalanceResponse.Body")
}

func init() { proto.RegisterFile("v2/accounting/grpc/service.proto", fileDescriptor_d9dd5af2ff2bbb25) }

var fileDescriptor_d9dd5af2ff2bbb25 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdb, 0x6a, 0xd4, 0x40,
	0x1c, 0xc6, 0x4d, 0x2c, 0xae, 0x4c, 0x3d, 0xe0, 0x20, 0x74, 0x89, 0x18, 0x4a, 0x69, 0x41, 0xc5,
	0xcc, 0x40, 0xbc, 0x10, 0x14, 0x2d, 0x5d, 0x6a, 0x71, 0x2f, 0x3c, 0xa5, 0xd0, 0x0b, 0x6f, 0xca,
	0x64, 0xf2, 0x4f, 0x3a, 0xb4, 0x99, 0x89, 0x99, 0xd9, 0x48, 0xde, 0xc4, 0x17, 0xf0, 0xc6, 0x0b,
	0x9f, 0xc3, 0x4b, 0x1f, 0x41, 0xd6, 0x17, 0x91, 0x1c, 0xf6, 0xc4, 0x66, 0xdd, 0xbd, 0xcb, 0xf0,
	0xff, 0xbe, 0x2f, 0xff, 0xef, 0xc7, 0x0c, 0xda, 0x2d, 0x7c, 0xca, 0x38, 0x57, 0x23, 0x69, 0x84,
	0x4c, 0x68, 0x92, 0x67, 0x9c, 0x6a, 0xc8, 0x0b, 0xc1, 0x81, 0x64, 0xb9, 0x32, 0x0a, 0xdf, 0x97,
	0xa0, 0x48, 0xac, 0x49, 0xe1, 0x93, 0x99, 0xd0, 0x71, 0x3b, 0x7c, 0xa6, 0xcc, 0x40, 0x37, 0x2e,
	0xa7, 0x5f, 0xf8, 0x34, 0x87, 0x58, 0x2f, 0x4f, 0x1e, 0x14, 0x3e, 0xd5, 0xa0, 0xb5, 0x50, 0x72,
	0x69, 0xb8, 0xf7, 0xdd, 0x46, 0x77, 0x06, 0xec, 0x8a, 0x49, 0x0e, 0x01, 0x7c, 0x19, 0x81, 0x36,
	0xf8, 0x15, 0xda, 0x0a, 0x55, 0x54, 0xf6, 0xad, 0x5d, 0xeb, 0xd1, 0xb6, 0xff, 0x98, 0x74, 0xad,
	0x43, 0x16, 0x3d, 0x64, 0xa0, 0xa2, 0x32, 0xa8, 0x6d, 0xf8, 0x0d, 0xda, 0x4e, 0xc1, 0xb0, 0xf3,
	0x0b, 0x60, 0x11, 0xe4, 0x7d, 0xbb, 0x4e, 0xd9, 0x9f, 0x4b, 0x69, 0x77, 0x21, 0xad, 0xf7, 0x1d,
	0x18, 0xf6, 0xb6, 0xd6, 0x06, 0x28, 0x9d, 0x7e, 0xe3, 0x4f, 0xe8, 0x76, 0x01, 0xb9, 0x88, 0xcb,
	0x49, 0xd0, 0xf5, 0x3a, 0xe8, 0xe9, 0xea, 0xa0, 0xb3, 0x4a, 0x2e, 0x38, 0x33, 0x42, 0xc9, 0x36,
	0xf0, 0x56, 0x13, 0xd1, 0x9c, 0x9c, 0x17, 0x68, 0xab, 0xda, 0x13, 0xfb, 0xe8, 0xa6, 0xfa, 0x2a,
	0x21, 0x3f, 0x17, 0x51, 0x5b, 0x72, 0x67, 0x2e, 0xb5, 0x82, 0x48, 0x3e, 0x54, 0xf3, 0xe1, 0x71,
	0xd0, 0xab, 0x85, 0xc3, 0x68, 0xef, 0xa7, 0x8d, 0xee, 0x4e, 0x3b, 0xeb, 0x4c, 0x49, 0x0d, 0xf8,
	0xf5, 0x02, 0xa8, 0x27, 0x6b, 0x40, 0x35, 0xa6, 0x79, 0x52, 0x27, 0x5d, 0xa4, 0x0e, 0x3a, 0x0b,
	0x36, 0xe6, 0x15, 0xa8, 0x82, 0x6e, 0x54, 0xde, 0x7f, 0x92, 0xd6, 0xb2, 0x3a, 0x6c, 0x59, 0x3d,
	0x47, 0xbd, 0xb0, 0x69, 0xd0, 0xd6, 0x7c, 0xd8, 0x5d, 0xf3, 0x18, 0xb8, 0x48, 0xd9, 0x55, 0x30,
	0x51, 0xfb, 0x97, 0xe8, 0xde, 0xd1, 0x74, 0x7c, 0xda, 0x5c, 0x70, 0x7c, 0x86, 0x7a, 0x2d, 0x0f,
	0xbc, 0xbf, 0xc9, 0xbd, 0x72, 0x0e, 0x36, 0x82, 0x3a, 0xb8, 0xfc, 0x35, 0x76, 0xad, 0xdf, 0x63,
	0xd7, 0xfa, 0x33, 0x76, 0xad, 0x6f, 0x7f, 0xdd, 0x6b, 0x9f, 0x0f, 0x13, 0x61, 0x2e, 0x46, 0x21,
	0xe1, 0x2a, 0xa5, 0x52, 0x67, 0x9c, 0x7b, 0x11, 0x14, 0x54, 0x82, 0x8a, 0xb5, 0xc7, 0x32, 0xe1,
	0x25, 0x8a, 0x2e, 0x3f, 0xa8, 0x97, 0xb3, 0xf3, 0x0f, 0x7b, 0xe7, 0x3d, 0xa8, 0x93, 0x53, 0x72,
	0xf4, 0x71, 0x58, 0xfd, 0x7c, 0xd6, 0x25, 0xbc, 0x51, 0xbf, 0x9c, 0x67, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x54, 0x53, 0x6d, 0x41, 0xca, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc2.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc2.SupportPackageIsVersion4

// AccountingServiceClient is the client API for AccountingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountingServiceClient interface {
	// Returns the amount of funds for the requested NeoFS account.
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc2.CallOption) (*BalanceResponse, error)
}

type accountingServiceClient struct {
	cc *grpc2.ClientConn
}

func NewAccountingServiceClient(cc *grpc2.ClientConn) AccountingServiceClient {
	return &accountingServiceClient{cc}
}

func (c *accountingServiceClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc2.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/neo.fs.v2.accounting.AccountingService/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountingServiceServer is the server API for AccountingService service.
type AccountingServiceServer interface {
	// Returns the amount of funds for the requested NeoFS account.
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
}

// UnimplementedAccountingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountingServiceServer struct {
}

func (*UnimplementedAccountingServiceServer) Balance(ctx context.Context, req *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}

func RegisterAccountingServiceServer(s *grpc2.Server, srv AccountingServiceServer) {
	s.RegisterService(&_AccountingService_serviceDesc, srv)
}

func _AccountingService_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc2.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServiceServer).Balance(ctx, in)
	}
	info := &grpc2.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neo.fs.v2.accounting.AccountingService/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServiceServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountingService_serviceDesc = grpc2.ServiceDesc{
	ServiceName: "neo.fs.v2.accounting.AccountingService",
	HandlerType: (*AccountingServiceServer)(nil),
	Methods: []grpc2.MethodDesc{
		{
			MethodName: "Balance",
			Handler:    _AccountingService_Balance_Handler,
		},
	},
	Streams:  []grpc2.StreamDesc{},
	Metadata: "v2/accounting/grpc/service.proto",
}

func (m *BalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.VerifyHeader != nil {
		{
			size, err := m.VerifyHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MetaHeader != nil {
		{
			size, err := m.MetaHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Body != nil {
		{
			size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BalanceRequest_Body) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalanceRequest_Body) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalanceRequest_Body) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.OwnerId != nil {
		{
			size, err := m.OwnerId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.VerifyHeader != nil {
		{
			size, err := m.VerifyHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MetaHeader != nil {
		{
			size, err := m.MetaHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Body != nil {
		{
			size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BalanceResponse_Body) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalanceResponse_Body) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalanceResponse_Body) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Balance != nil {
		{
			size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Body != nil {
		l = m.Body.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.MetaHeader != nil {
		l = m.MetaHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.VerifyHeader != nil {
		l = m.VerifyHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BalanceRequest_Body) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OwnerId != nil {
		l = m.OwnerId.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Body != nil {
		l = m.Body.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.MetaHeader != nil {
		l = m.MetaHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.VerifyHeader != nil {
		l = m.VerifyHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BalanceResponse_Body) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Balance != nil {
		l = m.Balance.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BalanceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Body == nil {
				m.Body = &BalanceRequest_Body{}
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaHeader == nil {
				m.MetaHeader = &grpc.RequestMetaHeader{}
			}
			if err := m.MetaHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VerifyHeader == nil {
				m.VerifyHeader = &grpc.RequestVerificationHeader{}
			}
			if err := m.VerifyHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BalanceRequest_Body) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Body: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Body: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OwnerId == nil {
				m.OwnerId = &grpc1.OwnerID{}
			}
			if err := m.OwnerId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BalanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Body == nil {
				m.Body = &BalanceResponse_Body{}
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaHeader == nil {
				m.MetaHeader = &grpc.ResponseMetaHeader{}
			}
			if err := m.MetaHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VerifyHeader == nil {
				m.VerifyHeader = &grpc.ResponseVerificationHeader{}
			}
			if err := m.VerifyHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BalanceResponse_Body) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Body: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Body: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Balance == nil {
				m.Balance = &Decimal{}
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
