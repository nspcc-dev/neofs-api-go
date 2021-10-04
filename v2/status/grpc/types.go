package status

// SetId sets identifier of the Status_Detail.
func (x *Status_Detail) SetId(v uint32) {
	if x != nil {
		x.Id = v
	}
}

// SetValue sets value of the Status_Detail.
func (x *Status_Detail) SetValue(v []byte) {
	if x != nil {
		x.Value = v
	}
}

// SetCode sets code of the Status.
func (x *Status) SetCode(v uint32) {
	if x != nil {
		x.Code = v
	}
}

// SetMessage sets message about the Status.
func (x *Status) SetMessage(v string) {
	if x != nil {
		x.Message = v
	}
}

// SetDetails sets details of the Status.
func (x *Status) SetDetails(v []*Status_Detail) {
	if x != nil {
		x.Details = v
	}
}
