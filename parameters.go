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

// Represents a parameter you want to include when generating signed
// authentication tokens for your clients.
type Parameter struct {
	// The field to match parameters based on.
	Field string

	// The operation to apply.
	Op ParameterOperation

	// The vlaue to compare the field against using the operation.
	Value string
}
