package unitvalue

// ExitType determines how the service manager detects that a service has
// exited.
type ExitType string

// Possible exit types.
const (
	ExitTypeMain   ExitType = "main"
	ExitTypeCgroup ExitType = "cgroup"
)

// Value returns a string representation of the exit type.
func (t ExitType) Value() string {
	return string(t)
}
