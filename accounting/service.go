package accounting

import (
	"github.com/nspcc-dev/neofs-api-go/decimal"
	"github.com/nspcc-dev/neofs-api-go/internal"
	"github.com/nspcc-dev/neofs-api-go/refs"
)

type (
	// OwnerID type alias.
	OwnerID = refs.OwnerID

	// Decimal type alias.
	Decimal = decimal.Decimal

	// Filter is used to filter accounts by criteria.
	Filter func(acc *Account) bool
)

const (
	// ErrEmptyAddress is raised when passed Address is empty.
	ErrEmptyAddress = internal.Error("empty address")

	// ErrEmptyLockTarget is raised when passed LockTarget is empty.
	ErrEmptyLockTarget = internal.Error("empty lock target")

	// ErrEmptyContainerID is raised when passed CID is empty.
	ErrEmptyContainerID = internal.Error("empty container ID")

	// ErrEmptyParentAddress is raised when passed ParentAddress is empty.
	ErrEmptyParentAddress = internal.Error("empty parent address")
)

// SumFunds goes through all accounts and sums up active funds.
func SumFunds(accounts []*Account) (res *decimal.Decimal) {
	res = decimal.Zero.Copy()

	for i := range accounts {
		if accounts[i] == nil {
			continue
		}

		res = res.Add(accounts[i].ActiveFunds)
	}
	return
}
