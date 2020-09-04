package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
)

// SetBody sets body of the request.
func (m *LocalNodeInfoRequest) SetBody(v *LocalNodeInfoRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *LocalNodeInfoRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *LocalNodeInfoRequest) SetVerifyHeader(v *session.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetVersion sets version of response body.
func (m *LocalNodeInfoResponse_Body) SetVersion(v *refs.Version) {
	if m != nil {
		m.Version = v
	}
}

// SetNodeInfo sets node info of response body.
func (m *LocalNodeInfoResponse_Body) SetNodeInfo(v *NodeInfo) {
	if m != nil {
		m.NodeInfo = v
	}
}

// SetBody sets body of the response.
func (m *LocalNodeInfoResponse) SetBody(v *LocalNodeInfoResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *LocalNodeInfoResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *LocalNodeInfoResponse) SetVerifyHeader(v *session.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}
