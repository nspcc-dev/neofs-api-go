package hash

import (
	"bytes"
)

// HashesSlice is a collection that satisfies sort.Interface and can be
// sorted by the routines in sort package.
type HashesSlice []Hash

// -- HashesSlice -- an inner type to sort Objects
// Len is the number of elements in the collection.
func (hs HashesSlice) Len() int { return len(hs) }

// Less reports whether the element with
// index i should be sorted before the element with index j.
func (hs HashesSlice) Less(i, j int) bool { return bytes.Compare(hs[i].Bytes(), hs[j].Bytes()) == -1 }

// Swap swaps the elements with indexes i and j.
func (hs HashesSlice) Swap(i, j int) { hs[i], hs[j] = hs[j], hs[i] }
