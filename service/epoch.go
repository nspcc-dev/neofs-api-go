package service

// EpochRequest interface gives possibility to get or set epoch in RPC Requests.
type EpochRequest interface {
	GetEpoch() uint64
	SetEpoch(v uint64)
}
