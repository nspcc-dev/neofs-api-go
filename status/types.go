package status

// Detail represents structure of NeoFS API V2-compatible status detail message.
type Detail struct {
	id uint32

	val []byte
}

// ID returns identifier of the Detail.
func (x *Detail) ID() uint32 {
	if x != nil {
		return x.id
	}

	return 0
}

// SetID sets identifier of the Detail.
func (x *Detail) SetID(id uint32) {
	x.id = id
}

// Value returns value of the Detail.
func (x *Detail) Value() []byte {
	if x != nil {
		return x.val
	}

	return nil
}

// SetValue sets value of the Detail.
func (x *Detail) SetValue(val []byte) {
	x.val = val
}

// Code represents NeoFS API V2-compatible status code.
type Code uint32

// EqualNumber checks if the numerical Code equals num.
func (x Code) EqualNumber(num uint32) bool {
	return uint32(x) == num
}

// Status represents structure of NeoFS API V2-compatible status return message.
type Status struct {
	code Code

	msg string

	details []Detail
}

// Code returns code of the Status.
func (x *Status) Code() Code {
	if x != nil {
		return x.code
	}

	return 0
}

// SetCode sets code of the Status.
func (x *Status) SetCode(code Code) {
	x.code = code
}

// Message sets message of the Status.
func (x *Status) Message() string {
	if x != nil {
		return x.msg
	}

	return ""
}

// SetMessage sets message of the Status.
func (x *Status) SetMessage(msg string) {
	x.msg = msg
}

// NumberOfParameters returns number of network parameters.
func (x *Status) NumberOfDetails() int {
	if x != nil {
		return len(x.details)
	}

	return 0
}

// IterateDetails iterates over details of the Status.
// Breaks iteration on f's true return.
//
// Handler must not be nil.
func (x *Status) IterateDetails(f func(*Detail) bool) {
	if x != nil {
		for i := range x.details {
			if f(&x.details[i]) {
				break
			}
		}
	}
}

// ResetDetails empties the detail list.
func (x *Status) ResetDetails() {
	if x != nil {
		x.details = x.details[:0]
	}
}

// AppendDetails appends the list of details to the Status.
func (x *Status) AppendDetails(ds ...Detail) {
	if x != nil {
		x.details = append(x.details, ds...)
	}
}

// SetStatusDetails sets Detail list of the Status.
func SetStatusDetails(dst *Status, ds []Detail) {
	dst.ResetDetails()
	dst.AppendDetails(ds...)
}
