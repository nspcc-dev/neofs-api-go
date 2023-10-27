// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: session/grpc/service.proto

package session

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SessionService_Create_FullMethodName = "/neo.fs.v2.session.SessionService/Create"
)

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionServiceClient interface {
	// Open a new session between two peers.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// session has been successfully opened;
	// - Common failures (SECTION_FAILURE_COMMON).
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type sessionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionServiceClient(cc grpc.ClientConnInterface) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, SessionService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
// All implementations should embed UnimplementedSessionServiceServer
// for forward compatibility
type SessionServiceServer interface {
	// Open a new session between two peers.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// session has been successfully opened;
	// - Common failures (SECTION_FAILURE_COMMON).
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
}

// UnimplementedSessionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSessionServiceServer struct {
}

func (UnimplementedSessionServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

// UnsafeSessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServiceServer will
// result in compilation errors.
type UnsafeSessionServiceServer interface {
	mustEmbedUnimplementedSessionServiceServer()
}

func RegisterSessionServiceServer(s grpc.ServiceRegistrar, srv SessionServiceServer) {
	s.RegisterService(&SessionService_ServiceDesc, srv)
}

func _SessionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SessionService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SessionService_ServiceDesc is the grpc.ServiceDesc for SessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "neo.fs.v2.session.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SessionService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session/grpc/service.proto",
}
