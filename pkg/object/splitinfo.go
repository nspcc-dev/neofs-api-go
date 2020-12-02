package object

import (
	"github.com/nspcc-dev/neofs-api-go/v2/object"
)

type SplitInfo object.SplitInfo

func NewSplitInfoFromV2(v2 *object.SplitInfo) *SplitInfo {
	return (*SplitInfo)(v2)
}

func NewSplitInfo() *SplitInfo {
	return NewSplitInfoFromV2(new(object.SplitInfo))
}

func (s *SplitInfo) ToV2() *object.SplitInfo {
	return (*object.SplitInfo)(s)
}

func (s *SplitInfo) SplitID() *SplitID {
	return NewSplitIDFromV2(
		(*object.SplitInfo)(s).GetSplitID(),
	)
}

func (s *SplitInfo) SetSplitID(v *SplitID) {
	(*object.SplitInfo)(s).SetSplitID(v.ToV2())
}

func (s *SplitInfo) LastPart() *ID {
	return NewIDFromV2(
		(*object.SplitInfo)(s).GetLastPart(),
	)
}

func (s *SplitInfo) SetLastPart(v *ID) {
	(*object.SplitInfo)(s).SetLastPart(v.ToV2())
}

func (s *SplitInfo) Link() *ID {
	return NewIDFromV2(
		(*object.SplitInfo)(s).GetLink(),
	)
}

func (s *SplitInfo) SetLink(v *ID) {
	(*object.SplitInfo)(s).SetLink(v.ToV2())
}
