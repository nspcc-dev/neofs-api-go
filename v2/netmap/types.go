package netmap

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
	address    string
	attributes []*Attribute
	state      NodeState
}

// NodeState of storage node.
type NodeState uint32

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

	}
	p.filters = filters
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

func (ni *NodeInfo) GetAddress() string {
	if ni != nil {
		return ni.address
	}

	return ""
}

func (ni *NodeInfo) SetAddress(v string) {
	if ni != nil {
		ni.address = v
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
