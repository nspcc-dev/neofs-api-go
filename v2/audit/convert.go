package audit

import (
	neofsgrpc "github.com/nspcc-dev/neofs-api-go/rpc/grpc"
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	audit "github.com/nspcc-dev/neofs-api-go/v2/audit/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
)

func (a *DataAuditResult) ToGRPCMessage() neofsgrpc.Message {
	var m *audit.DataAuditResult

	if a != nil {
		m = new(audit.DataAuditResult)

		m.SetAuditEpoch(a.auditEpoch)
		m.SetPublicKey(a.pubKey)
		m.SetContainerId(a.cid.ToGRPCMessage().(*refsGRPC.ContainerID))
		m.SetComplete(a.complete)
		m.SetVersion(a.version.ToGRPCMessage().(*refsGRPC.Version))
		m.SetPassNodes(a.passNodes)
		m.SetFailNodes(a.failNodes)
		m.SetRetries(a.retries)
		m.SetRequests(a.requests)
		m.SetHit(a.hit)
		m.SetMiss(a.miss)
		m.SetFail(a.fail)
		m.SetPassSg(refs.ObjectIDListToGRPCMessage(a.passSG))
		m.SetFailSg(refs.ObjectIDListToGRPCMessage(a.failSG))
	}

	return m
}

func (a *DataAuditResult) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*audit.DataAuditResult)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	cid := v.GetContainerId()
	if cid == nil {
		a.cid = nil
	} else {
		if a.cid == nil {
			a.cid = new(refs.ContainerID)
		}

		err = a.cid.FromGRPCMessage(cid)
		if err != nil {
			return err
		}
	}

	version := v.GetVersion()
	if version == nil {
		a.version = nil
	} else {
		if a.version == nil {
			a.version = new(refs.Version)
		}

		err = a.version.FromGRPCMessage(version)
		if err != nil {
			return err
		}
	}

	a.passSG, err = refs.ObjectIDListFromGRPCMessage(v.GetPassSg())
	if err != nil {
		return err
	}

	a.failSG, err = refs.ObjectIDListFromGRPCMessage(v.GetFailSg())
	if err != nil {
		return err
	}

	a.auditEpoch = v.GetAuditEpoch()
	a.pubKey = v.GetPublicKey()
	a.complete = v.GetComplete()
	a.passNodes = v.GetPassNodes()
	a.failNodes = v.GetFailNodes()
	a.retries = v.GetRetries()
	a.requests = v.GetRequests()
	a.hit = v.GetHit()
	a.miss = v.GetMiss()
	a.fail = v.GetFail()

	return err
}
