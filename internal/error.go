package internal

// Error is constant type error.
type Error string

// Error implementation of error interface.
func (e Error) Error() string {
	return string(e)
}
