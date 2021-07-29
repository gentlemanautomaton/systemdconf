package unitvalue

// ExitStatus is a program exit status code.
//
// TODO: Consider making this a uint8, or perhaps a struct with a hidden
// integer field.
type ExitStatus string

// Value returns a string representation of the exit status.
func (s ExitStatus) Value() string {
	return string(s)
}
