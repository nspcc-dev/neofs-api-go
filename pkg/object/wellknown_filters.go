package object

const (
	valFalse = "false"
	valTrue  = "true"
)

// ReservedFilterPrefix is a reserved prefix for system search filter keys.
const ReservedFilterPrefix = "$Object:"

const (
	// KeyRoot is a reserved search filter key to source objects.
	KeyRoot = ReservedFilterPrefix + "ROOT"

	// ValRoot is a value of root object filter.
	ValRoot = valTrue

	// ValNonRoot is a value of non-root object filter.
	ValNonRoot = valFalse
)
