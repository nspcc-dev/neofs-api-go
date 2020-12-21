package audit

import (
	audit "github.com/nspcc-dev/neofs-api-go/v2/audit/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// DataAuditResultToGRPCMessage converts unified DataAuditResult structure
// into gRPC DataAuditResult message.
func DataAuditResultToGRPCMessage(a *DataAuditResult) *audit.DataAuditResult {
	if a == nil {
		return nil
	}

	m := new(audit.DataAuditResult)

	m.SetVersion(
		refs.VersionToGRPCMessage(a.GetVersion()),
	)

	m.SetAuditEpoch(a.GetAuditEpoch())

	m.SetContainerId(
		refs.ContainerIDToGRPCMessage(a.GetContainerID()),
	)

	m.SetPublicKey(a.GetPublicKey())

	m.SetComplete(a.GetComplete())

	m.SetPassSg(
		refs.ObjectIDListToGRPCMessage(a.GetPassSG()),
	)

	m.SetFailSg(
		refs.ObjectIDListToGRPCMessage(a.GetFailSG()),
	)

	m.SetHit(a.GetHit())
	m.SetMiss(a.GetMiss())
	m.SetFail(a.GetFail())

	m.SetPassNodes(a.GetPassNodes())
	m.SetFailNodes(a.GetFailNodes())

	return m
}

// DataAuditResultFromGRPCMessage converts gRPC message DataAuditResult
// into unified DataAuditResult structure.
func DataAuditResultFromGRPCMessage(m *audit.DataAuditResult) *DataAuditResult {
	if m == nil {
		return nil
	}

	a := new(DataAuditResult)

	a.SetVersion(
		refs.VersionFromGRPCMessage(m.GetVersion()),
	)

	a.SetAuditEpoch(m.GetAuditEpoch())

	a.SetContainerID(
		refs.ContainerIDFromGRPCMessage(m.GetContainerId()),
	)

	a.SetPublicKey(m.GetPublicKey())

	a.SetComplete(m.GetComplete())

	a.SetPassSG(
		refs.ObjectIDListFromGRPCMessage(m.GetPassSg()),
	)

	a.SetFailSG(
		refs.ObjectIDListFromGRPCMessage(m.GetFailSg()),
	)

	a.SetHit(m.GetHit())
	a.SetMiss(m.GetMiss())
	a.SetFail(m.GetFail())

	a.SetPassNodes(m.GetPassNodes())
	a.SetFailNodes(m.GetFailNodes())

	return a
}
