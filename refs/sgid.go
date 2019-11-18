package refs

import (
	"github.com/pkg/errors"
)

// SGIDFromBytes parse bytes representation of SGID into new SGID value.
func SGIDFromBytes(data []byte) (sgid SGID, err error) {
	if ln := len(data); ln != SGIDSize {
		return SGID{}, errors.Wrapf(ErrWrongDataSize, "expect=%d, actual=%d", SGIDSize, ln)
	}
	copy(sgid[:], data)
	return
}
