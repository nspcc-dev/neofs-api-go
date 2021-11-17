package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

type LocalNodeInfoRequest struct {
	body *LocalNodeInfoRequestBody

	session.RequestHeaders
}

type LocalNodeInfoResponse struct {
	body *LocalNodeInfoResponseBody

	session.ResponseHeaders
}

// NetworkInfoRequest is a structure of NetworkInfo request.
type NetworkInfoRequest struct {
	body *NetworkInfoRequestBody

	session.RequestHeaders
}

// NetworkInfoResponse is a structure of NetworkInfo response.
type NetworkInfoResponse struct {
	body *NetworkInfoResponseBody

	session.ResponseHeaders
}

type Filter struct {
	name    string
	key     string
	op      Operation
	value   string
	filters []*Filter
}

type Selector struct {
	name      string
	count     uint32
	clause    Clause
	attribute string
	filter    string
}

type Replica struct {
	count    uint32
	selector string
}

type Operation uint32

type PlacementPolicy struct {
	replicas     []*Replica
	backupFactor uint32
	selectors    []*Selector
	filters      []*Filter
}

// Attribute of storage node.
type Attribute struct {
	key     string
	value   string
	parents []string
}

// NodeInfo of storage node.
type NodeInfo struct {
	publicKey  []byte
	addresses  []string
	attributes []*Attribute
	state      NodeState
}

// NodeState of storage node.
type NodeState uint32

// Clause of placement selector.
type Clause uint32

type LocalNodeInfoRequestBody struct{}

type LocalNodeInfoResponseBody struct {
	version  *refs.Version
	nodeInfo *NodeInfo
}

const (
	UnspecifiedState NodeState = iota
	Online
	Offline
)

const (
	UnspecifiedOperation Operation = iota
	EQ
	NE
	GT
	GE
	LT
	LE
	OR
	AND
)

const (
	UnspecifiedClause Clause = iota
	Same
	Distinct
)

func (f *Filter) GetFilters() []*Filter {
	if f != nil {
		return f.filters
	}

	return nil
}

func (f *Filter) SetFilters(filters []*Filter) {
	if f != nil {
		f.filters = filters
	}
}

func (f *Filter) GetValue() string {
	if f != nil {
		return f.value
	}

	return ""
}

func (f *Filter) SetValue(value string) {
	if f != nil {
		f.value = value
	}
}

func (f *Filter) GetOp() Operation {
	if f != nil {
		return f.op
	}
	return UnspecifiedOperation
}

func (f *Filter) SetOp(op Operation) {
	if f != nil {
		f.op = op
	}
}

func (f *Filter) GetKey() string {
	if f != nil {
		return f.key
	}

	return ""
}

func (f *Filter) SetKey(key string) {
	if f != nil {
		f.key = key
	}
}

func (f *Filter) GetName() string {
	if f != nil {
		return f.name
	}

	return ""
}

func (f *Filter) SetName(name string) {
	if f != nil {
		f.name = name
	}
}

func (s *Selector) GetFilter() string {
	if s != nil {
		return s.filter
	}

	return ""
}

func (s *Selector) SetFilter(filter string) {
	if s != nil {
		s.filter = filter
	}
}

func (s *Selector) GetAttribute() string {
	if s != nil {
		return s.attribute
	}

	return ""
}

func (s *Selector) SetAttribute(attribute string) {
	if s != nil {
		s.attribute = attribute
	}
}

func (s *Selector) GetClause() Clause {
	if s != nil {
		return s.clause
	}

	return UnspecifiedClause
}

func (s *Selector) SetClause(clause Clause) {
	if s != nil {
		s.clause = clause
	}
}

func (s *Selector) GetCount() uint32 {
	if s != nil {
		return s.count
	}

	return 0
}

func (s *Selector) SetCount(count uint32) {
	if s != nil {
		s.count = count
	}
}

func (s *Selector) GetName() string {
	if s != nil {
		return s.name
	}

	return ""
}

func (s *Selector) SetName(name string) {
	if s != nil {
		s.name = name
	}
}

func (r *Replica) GetSelector() string {
	if r != nil {
		return r.selector
	}

	return ""
}

func (r *Replica) SetSelector(selector string) {
	if r != nil {
		r.selector = selector
	}
}

func (r *Replica) GetCount() uint32 {
	if r != nil {
		return r.count
	}

	return 0
}

func (r *Replica) SetCount(count uint32) {
	if r != nil {
		r.count = count
	}
}

func (p *PlacementPolicy) GetFilters() []*Filter {
	if p != nil {
		return p.filters
	}

	return nil
}

func (p *PlacementPolicy) SetFilters(filters []*Filter) {
	if p != nil {
		p.filters = filters
	}
}

func (p *PlacementPolicy) GetSelectors() []*Selector {
	if p != nil {
		return p.selectors
	}

	return nil
}

func (p *PlacementPolicy) SetSelectors(selectors []*Selector) {
	if p != nil {
		p.selectors = selectors
	}
}

func (p *PlacementPolicy) GetContainerBackupFactor() uint32 {
	if p != nil {
		return p.backupFactor
	}

	return 0
}

func (p *PlacementPolicy) SetContainerBackupFactor(backupFactor uint32) {
	if p != nil {
		p.backupFactor = backupFactor
	}
}

func (p *PlacementPolicy) GetReplicas() []*Replica {
	return p.replicas
}

func (p *PlacementPolicy) SetReplicas(replicas []*Replica) {
	p.replicas = replicas
}

func (a *Attribute) GetKey() string {
	if a != nil {
		return a.key
	}

	return ""
}

func (a *Attribute) SetKey(v string) {
	if a != nil {
		a.key = v
	}
}

func (a *Attribute) GetValue() string {
	if a != nil {
		return a.value
	}

	return ""
}

func (a *Attribute) SetValue(v string) {
	if a != nil {
		a.value = v
	}
}

func (a *Attribute) GetParents() []string {
	if a != nil {
		return a.parents
	}

	return nil
}

func (a *Attribute) SetParents(parent []string) {
	if a != nil {
		a.parents = parent
	}
}

func (ni *NodeInfo) GetPublicKey() []byte {
	if ni != nil {
		return ni.publicKey
	}

	return nil
}

func (ni *NodeInfo) SetPublicKey(v []byte) {
	if ni != nil {
		ni.publicKey = v
	}
}

// GetAddress returns node's network address.
//
// Deprecated: use IterateAddresses.
func (ni *NodeInfo) GetAddress() (addr string) {
	ni.IterateAddresses(func(s string) bool {
		addr = s
		return true
	})

	return
}

// SetAddress sets node's network address.
//
// Deprecated: use SetAddresses.
func (ni *NodeInfo) SetAddress(v string) {
	ni.SetAddresses(v)
}

// SetAddresses sets list of network addresses of the node.
func (ni *NodeInfo) SetAddresses(v ...string) {
	if ni != nil {
		ni.addresses = v
	}
}

// NumberOfAddresses returns number of network addresses of the node.
func (ni *NodeInfo) NumberOfAddresses() int {
	if ni != nil {
		return len(ni.addresses)
	}

	return 0
}

// IterateAddresses iterates over network addresses of the node.
// Breaks iteration on f's true return.
//
// Handler should not be nil.
func (ni *NodeInfo) IterateAddresses(f func(string) bool) {
	if ni != nil {
		for i := range ni.addresses {
			if f(ni.addresses[i]) {
				break
			}
		}
	}
}

func (ni *NodeInfo) GetAttributes() []*Attribute {
	if ni != nil {
		return ni.attributes
	}

	return nil
}

func (ni *NodeInfo) SetAttributes(v []*Attribute) {
	if ni != nil {
		ni.attributes = v
	}
}

func (ni *NodeInfo) GetState() NodeState {
	if ni != nil {
		return ni.state
	}

	return UnspecifiedState
}

func (ni *NodeInfo) SetState(state NodeState) {
	if ni != nil {
		ni.state = state
	}
}

func (l *LocalNodeInfoResponseBody) GetVersion() *refs.Version {
	if l != nil {
		return l.version
	}

	return nil
}

func (l *LocalNodeInfoResponseBody) SetVersion(version *refs.Version) {
	if l != nil {
		l.version = version
	}
}

func (l *LocalNodeInfoResponseBody) GetNodeInfo() *NodeInfo {
	if l != nil {
		return l.nodeInfo
	}

	return nil
}

func (l *LocalNodeInfoResponseBody) SetNodeInfo(nodeInfo *NodeInfo) {
	if l != nil {
		l.nodeInfo = nodeInfo
	}
}

func (l *LocalNodeInfoRequest) GetBody() *LocalNodeInfoRequestBody {
	if l != nil {
		return l.body
	}
	return nil
}

func (l *LocalNodeInfoRequest) SetBody(body *LocalNodeInfoRequestBody) {
	if l != nil {
		l.body = body
	}
}

func (l *LocalNodeInfoResponse) GetBody() *LocalNodeInfoResponseBody {
	if l != nil {
		return l.body
	}
	return nil
}

func (l *LocalNodeInfoResponse) SetBody(body *LocalNodeInfoResponseBody) {
	if l != nil {
		l.body = body
	}
}

// NetworkParameter represents NeoFS network parameter.
type NetworkParameter struct {
	k, v []byte
}

// GetKey returns parameter key.
func (x *NetworkParameter) GetKey() []byte {
	if x != nil {
		return x.k
	}

	return nil
}

// SetKey sets parameter key.
func (x *NetworkParameter) SetKey(k []byte) {
	if x != nil {
		x.k = k
	}
}

// GetValue returns parameter value.
func (x *NetworkParameter) GetValue() []byte {
	if x != nil {
		return x.v
	}

	return nil
}

// SetValue sets parameter value.
func (x *NetworkParameter) SetValue(v []byte) {
	if x != nil {
		x.v = v
	}
}

// NetworkConfig represents NeoFS network configuration.
type NetworkConfig struct {
	ps []*NetworkParameter
}

// NumberOfParameters returns number of network parameters.
func (x *NetworkConfig) NumberOfParameters() int {
	if x != nil {
		return len(x.ps)
	}

	return 0
}

// IterateParameters iterates over network parameters.
// Breaks iteration on f's true return.
//
// Handler must not be nil.
func (x *NetworkConfig) IterateParameters(f func(*NetworkParameter) bool) {
	if x != nil {
		for i := range x.ps {
			if f(x.ps[i]) {
				break
			}
		}
	}
}

// SetParameters sets list of network parameters.
func (x *NetworkConfig) SetParameters(v ...*NetworkParameter) {
	if x != nil {
		x.ps = v
	}
}

// NetworkInfo groups information about
// NeoFS network.
type NetworkInfo struct {
	curEpoch, magicNum uint64

	msPerBlock int64

	netCfg *NetworkConfig
}

// GetCurrentEpoch returns number of the current epoch.
func (i *NetworkInfo) GetCurrentEpoch() uint64 {
	if i != nil {
		return i.curEpoch
	}

	return 0
}

// SetCurrentEpoch sets number of the current epoch.
func (i *NetworkInfo) SetCurrentEpoch(epoch uint64) {
	if i != nil {
		i.curEpoch = epoch
	}
}

// GetMagicNumber returns magic number of the sidechain.
func (i *NetworkInfo) GetMagicNumber() uint64 {
	if i != nil {
		return i.magicNum
	}

	return 0
}

// SetMagicNumber sets magic number of the sidechain.
func (i *NetworkInfo) SetMagicNumber(magic uint64) {
	if i != nil {
		i.magicNum = magic
	}
}

// GetMsPerBlock returns MillisecondsPerBlock network parameter.
func (i *NetworkInfo) GetMsPerBlock() int64 {
	if i != nil {
		return i.msPerBlock
	}

	return 0
}

// SetMsPerBlock sets MillisecondsPerBlock network parameter.
func (i *NetworkInfo) SetMsPerBlock(v int64) {
	if i != nil {
		i.msPerBlock = v
	}
}

// GetNetworkConfig returns NeoFS network configuration.
func (i *NetworkInfo) GetNetworkConfig() *NetworkConfig {
	if i != nil {
		return i.netCfg
	}

	return nil
}

// SetNetworkConfig sets NeoFS network configuration.
func (i *NetworkInfo) SetNetworkConfig(v *NetworkConfig) {
	if i != nil {
		i.netCfg = v
	}
}

// NetworkInfoRequestBody is a structure of NetworkInfo request body.
type NetworkInfoRequestBody struct{}

// NetworkInfoResponseBody is a structure of NetworkInfo response body.
type NetworkInfoResponseBody struct {
	netInfo *NetworkInfo
}

// GetNetworkInfo returns information about the NeoFS network.
func (i *NetworkInfoResponseBody) GetNetworkInfo() *NetworkInfo {
	if i != nil {
		return i.netInfo
	}

	return nil
}

// SetNetworkInfo sets information about the NeoFS network.
func (i *NetworkInfoResponseBody) SetNetworkInfo(netInfo *NetworkInfo) {
	if i != nil {
		i.netInfo = netInfo
	}
}

func (l *NetworkInfoRequest) GetBody() *NetworkInfoRequestBody {
	if l != nil {
		return l.body
	}
	return nil
}

func (l *NetworkInfoRequest) SetBody(body *NetworkInfoRequestBody) {
	if l != nil {
		l.body = body
	}
}

func (l *NetworkInfoResponse) GetBody() *NetworkInfoResponseBody {
	if l != nil {
		return l.body
	}
	return nil
}

func (l *NetworkInfoResponse) SetBody(body *NetworkInfoResponseBody) {
	if l != nil {
		l.body = body
	}
}
