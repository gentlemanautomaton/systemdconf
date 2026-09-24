package unitvalue

// OutOfMemoryPolicy defines the desired service behavior when the system is
// low on memory.
type OutOfMemoryPolicy string

// Possible out-of-memory policies.
const (
	OutOfMemoryPolicyContinue OutOfMemoryPolicy = "continue"
	OutOfMemoryPolicyStop     OutOfMemoryPolicy = "stop"
	OutOfMemoryPolicyKill     OutOfMemoryPolicy = "kill"
)

// Value returns a string representation of the out-of-memory policy.
func (p OutOfMemoryPolicy) Value() string {
	return string(p)
}
