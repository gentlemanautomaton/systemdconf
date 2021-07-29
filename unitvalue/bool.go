package unitvalue

// Bool is a boolean value that can be "true" or "false".
type Bool string

// Possible boolean values.
const (
	True  Bool = "true"
	False Bool = "false"
)

// Value returns a string representation of the bool.
func (b Bool) Value() string {
	return string(b)
}

// YesNo is a boolean value that can be "yes" or "no".
type YesNo string

// Possible yes/no values.
const (
	Yes YesNo = "yes"
	No  YesNo = "no"
)

// Value returns a string representation of the bool.
func (b YesNo) Value() string {
	return string(b)
}
