package container

import (
	acl "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	service "github.com/nspcc-dev/neofs-api-go/v2/service/grpc"
)

// SetContainer sets container of the request.
func (m *PutRequest_Body) SetContainer(v *Container) {
	if m != nil {
		m.Container = v
	}
}

// SetPublicKey sets public key of the container owner.
func (m *PutRequest_Body) SetPublicKey(v []byte) {
	if m != nil {
		m.PublicKey = v
	}
}

// SetSignature sets signature of the container structure.
func (m *PutRequest_Body) SetSignature(v []byte) {
	if m != nil {
		m.Signature = v
	}
}

// SetBody sets body of the request.
func (m *PutRequest) SetBody(v *PutRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *PutRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *PutRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetContainerId sets identifier of the container.
func (m *PutResponse_Body) SetContainerId(v *refs.ContainerID) {
	if m != nil {
		m.ContainerId = v
	}
}

// SetBody sets body of the response.
func (m *PutResponse) SetBody(v *PutResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *PutResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *PutResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetContainerId sets identifier of the container.
func (m *DeleteRequest_Body) SetContainerId(v *refs.ContainerID) {
	if m != nil {
		m.ContainerId = v
	}
}

// SetSignature sets signature of the container identifier.
func (m *DeleteRequest_Body) SetSignature(v []byte) {
	if m != nil {
		m.Signature = v
	}
}

// SetBody sets body of the request.
func (m *DeleteRequest) SetBody(v *DeleteRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *DeleteRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *DeleteRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetBody sets body of the response.
func (m *DeleteResponse) SetBody(v *DeleteResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *DeleteResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *DeleteResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetContainerId sets identifier of the container.
func (m *GetRequest_Body) SetContainerId(v *refs.ContainerID) {
	m.ContainerId = v
}

// SetBody sets body of the request.
func (m *GetRequest) SetBody(v *GetRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *GetRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *GetRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetContainer sets the container structure.
func (m *GetResponse_Body) SetContainer(v *Container) {
	if m != nil {
		m.Container = v
	}
}

// SetBody sets body of the response.
func (m *GetResponse) SetBody(v *GetResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *GetResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *GetResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetOwnerId sets identifier of the container owner.
func (m *ListRequest_Body) SetOwnerId(v *refs.OwnerID) {
	if m != nil {
		m.OwnerId = v
	}
}

// SetBody sets body of the request.
func (m *ListRequest) SetBody(v *ListRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *ListRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *ListRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetContainerIds sets list of the container identifiers.
func (m *ListResponse_Body) SetContainerIds(v []*refs.ContainerID) {
	if m != nil {
		m.ContainerIds = v
	}
}

// SetBody sets body of the response.
func (m *ListResponse) SetBody(v *ListResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *ListResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *ListResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetEacl sets eACL table structure.
func (m *SetExtendedACLRequest_Body) SetEacl(v *acl.EACLTable) {
	if m != nil {
		m.Eacl = v
	}
}

// SetSignature sets signature of the eACL table.
func (m *SetExtendedACLRequest_Body) SetSignature(v []byte) {
	if m != nil {
		m.Signature = v
	}
}

// SetBody sets body of the request.
func (m *SetExtendedACLRequest) SetBody(v *SetExtendedACLRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *SetExtendedACLRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *SetExtendedACLRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetBody sets body of the response.
func (m *SetExtendedACLResponse) SetBody(v *SetExtendedACLResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *SetExtendedACLResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *SetExtendedACLResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetContainerId sets identifier of the container.
func (m *GetExtendedACLRequest_Body) SetContainerId(v *refs.ContainerID) {
	if m != nil {
		m.ContainerId = v
	}
}

// SetBody sets body of the request.
func (m *GetExtendedACLRequest) SetBody(v *GetExtendedACLRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *GetExtendedACLRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *GetExtendedACLRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetEacl sets eACL table structure.
func (m *GetExtendedACLResponse_Body) SetEacl(v *acl.EACLTable) {
	if m != nil {
		m.Eacl = v
	}
}

// SetSignature sets signature of the eACL table.
func (m *GetExtendedACLResponse_Body) SetSignature(v []byte) {
	if m != nil {
		m.Signature = v
	}
}

// SetBody sets body of the response.
func (m *GetExtendedACLResponse) SetBody(v *GetExtendedACLResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *GetExtendedACLResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *GetExtendedACLResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}
