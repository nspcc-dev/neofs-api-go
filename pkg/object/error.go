package object

type SplitInfoError struct {
	si *SplitInfo
}

const splitInfoErrorMsg = "object not found, split info has been provided"

func (s *SplitInfoError) Error() string {
	return splitInfoErrorMsg
}

func (s *SplitInfoError) SplitInfo() *SplitInfo {
	return s.si
}

func NewSplitInfoError(v *SplitInfo) *SplitInfoError {
	return &SplitInfoError{si: v}
}
