package common

// CallMethodInfo is an information about the RPC.
//
// Describes unary RPC by default.
type CallMethodInfo struct {
	svcName string

	mName string

	sClient, sServer bool
}

// SetServerStream makes CallMethodInfo to describe
// server-side streaming RPC.
func (x *CallMethodInfo) SetServerStream() {
	x.sServer = true
}

// ServerStream checks if CallMethodInfo contains
// information about the server-side streaming RPC.
func (x CallMethodInfo) ServerStream() bool {
	return x.sServer
}

// SetClientStream makes CallMethodInfo to describe
// client-side streaming RPC.
func (x *CallMethodInfo) SetClientStream() {
	x.sClient = true
}

// ClientStream checks if CallMethodInfo contains
// information about the client-side streaming RPC.
func (x CallMethodInfo) ClientStream() bool {
	return x.sClient
}

// SetServiceName sets service name.
func (x *CallMethodInfo) SetServiceName(service string) {
	x.svcName = service
}

// ServiceName returns service name.
func (x CallMethodInfo) ServiceName() string {
	return x.svcName
}

// SetMethodName sets method name.
func (x *CallMethodInfo) SetMethodName(method string) {
	x.mName = method
}

// MethodName returns method name.
func (x CallMethodInfo) MethodName() string {
	return x.mName
}
