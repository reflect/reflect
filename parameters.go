package reflect

type ParameterOperation string

const (
	EqualsOperation              = ParameterOperation("=")
	NotEqualsOperation           = ParameterOperation("!=")
	GreaterThanOperation         = ParameterOperation(">")
	GreaterThanOrEqualsOperation = ParameterOperation(">=")
	LessThanOperation            = ParameterOperation("<")
	LessThanOrEqualsOperation    = ParameterOperation("<=")
)

type Parameter struct {
	Field string
	Op    ParameterOperation
	Value string
}
