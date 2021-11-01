package unitvalue

// FileMode is the mode of a file in octal format.
//
// TODO: Consider making this a uint32.
type FileMode string

// Value returns a string representation of the file mode in octal format.
func (s FileMode) Value() string {
	return string(s)
}
