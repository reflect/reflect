package reflect

// Different operations to match parameters with.
type ParameterOperation string

const (
	EqualsOperation              = ParameterOperation("=")
	NotEqualsOperation           = ParameterOperation("!=")
	GreaterThanOperation         = ParameterOperation(">")
	GreaterThanOrEqualsOperation = ParameterOperation(">=")
	LessThanOperation            = ParameterOperation("<")
	LessThanOrEqualsOperation    = ParameterOperation("<=")
)

// A parameter you want to include when generating signed authentication tokens
// for your clients.
type Parameter struct {
	// The name of the field you want to enforce authentication against.
	Field string

	// The operation to apply when enforcing authentication.
	Op ParameterOperation

	// The value to use when enforcing authentication.
	Value string
}
