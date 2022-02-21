package status

const sectionBitSize = 10

// InSections checks if the Code is in [i,j] section list.
func (x Code) InSections(i, j uint32) bool {
	return uint32(x) >= i<<sectionBitSize && uint32(x) < (j+1)<<sectionBitSize
}

// LocalizeSection localizes the Code to the sec-th section.
//
// Does not make sense if the Code is outside the section.
func (x *Code) LocalizeSection(sec uint32) {
	*x = *x - Code(sec<<sectionBitSize)
}

// GlobalizeSection globalizes the Code of the sec-th section.
//
// Does not make sense if the Code is outside the section.
func (x *Code) GlobalizeSection(sec uint32) {
	*x = *x + Code(sec<<sectionBitSize)
}

// IsInSection returns true if the Code belongs to sec-th section.
func IsInSection(code Code, sec uint32) bool {
	return code.InSections(sec, sec)
}

const successSections = 1

// IsSuccess checks if the Code is a success code.
func IsSuccess(c Code) bool {
	return c.InSections(0, successSections-1)
}

// LocalizeSuccess localizes the Code to the success section.
func LocalizeSuccess(c *Code) {
	c.LocalizeSection(0)
}

// GlobalizeSuccess globalizes the Code to the success section.
func GlobalizeSuccess(c *Code) {
	c.GlobalizeSection(0)
}

func sectionAfterSuccess(sec uint32) uint32 {
	return successSections + sec
}

// Success codes.
const (
	// OK is a local status Code value for default success.
	OK Code = iota
)

// Common failure codes.
const (
	// Internal is a local Code value for INTERNAL failure status.
	Internal Code = iota
	// WrongMagicNumber is a local Code value for WRONG_MAGIC_NUMBER failure status.
	WrongMagicNumber
)

const (
	_ = iota - 1
	sectionCommon
)

// IsCommonFail checks if the Code is a common failure code.
func IsCommonFail(c Code) bool {
	return IsInSection(c, sectionAfterSuccess(sectionCommon))
}

// LocalizeCommonFail localizes the Code to the common fail section.
func LocalizeCommonFail(c *Code) {
	c.LocalizeSection(sectionAfterSuccess(sectionCommon))
}

// GlobalizeCommonFail globalizes the Code to the common fail section.
func GlobalizeCommonFail(c *Code) {
	c.GlobalizeSection(sectionAfterSuccess(sectionCommon))
}

// LocalizeIfInSection checks if passed global status.Code belongs to the section and:
//   then localizes the code and returns true,
//   else leaves the code unchanged and returns false.
//
// Arg must not be nil.
func LocalizeIfInSection(c *Code, sec uint32) bool {
	if IsInSection(*c, sec) {
		c.LocalizeSection(sec)
		return true
	}

	return false
}
