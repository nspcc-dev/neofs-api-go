package common

type callType uint8

const (
	_ callType = iota
	callUnary
	callClientStream
	callServerStream
	callBidirStream
)

// CallMethodInfo is an information about the RPC.
type CallMethodInfo struct {
	// Name of the service.
	Service string

	// Name of the RPC.
	Name string

	t callType
}

// ServerStream checks if CallMethodInfo contains
// information about the server-side streaming RPC.
func (c CallMethodInfo) ServerStream() bool {
	return c.t == callServerStream || c.t == callBidirStream
}

// ClientStream checks if CallMethodInfo contains
// information about the client-side streaming RPC.
func (c CallMethodInfo) ClientStream() bool {
	return c.t == callClientStream || c.t == callBidirStream
}

func (c *CallMethodInfo) setCommon(service, name string) {
	c.Service = service
	c.Name = name
}

// CallMethodInfoUnary returns CallMethodInfo structure
// initialized for the unary RPC.
func CallMethodInfoUnary(service, name string) (info CallMethodInfo) {
	info.setCommon(service, name)
	info.t = callUnary

	return
}

// CallMethodInfoClientStream returns CallMethodInfo structure
// initialized for the client-side streaming RPC.
func CallMethodInfoClientStream(service, name string) (info CallMethodInfo) {
	info.setCommon(service, name)
	info.t = callClientStream

	return
}

// CallMethodInfoServerStream returns CallMethodInfo structure
// initialized for the server-side streaming RPC.
func CallMethodInfoServerStream(service, name string) (info CallMethodInfo) {
	info.setCommon(service, name)
	info.t = callServerStream

	return
}

// CallMethodInfoBidirectionalStream returns CallMethodInfo structure
// initialized for the bidirectional streaming RPC.
func CallMethodInfoBidirectionalStream(service, name string) (info CallMethodInfo) {
	info.setCommon(service, name)
	info.t = callBidirStream

	return
}
