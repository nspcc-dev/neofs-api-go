package acl

// RequestInfo is an interface of request information needed for extended ACL check.
type RequestInfo interface {
	TypedHeaderSource

	// Must return the binary representation of request initiator's key.
	Key() []byte

	// Must return true if request corresponds to operation type.
	TypeOf(OperationType) bool

	// Must return true if request has passed target.
	TargetOf(Target) bool
}

// ExtendedACLChecker is an interface of extended ACL checking tool.
type ExtendedACLChecker interface {
	// Must return an action according to the results of applying the ACL table rules to request.
	//
	// Must return ActionUndefined if it is unable to explicitly calculate the action.
	Action(ExtendedACLTable, RequestInfo) ExtendedACLAction
}

type extendedACLChecker struct{}

const (
	// ActionUndefined is ExtendedACLAction used to mark value as undefined.
	// Most of the tools consider ActionUndefined as incalculable.
	// Using ActionUndefined in ExtendedACLRecord is unsafe.
	ActionUndefined ExtendedACLAction = iota

	// ActionAllow is ExtendedACLAction used to mark an applicability of ACL rule.
	ActionAllow

	// ActionDeny is ExtendedACLAction used to mark an inapplicability of ACL rule.
	ActionDeny
)
