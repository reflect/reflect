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
	// The name of the field this parameter applies to.
	Field string

	// The operation to apply to this field and value.
	Op ParameterOperation

	// The value to compare against.
	Value string
}
