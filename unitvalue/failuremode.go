package unitvalue

// FailureMode indicates how to handle a failed unit.
type FailureMode string

// Possible collection modes.
const (
	FailureModeTerminate FailureMode = "terminate"
	FailureModeAbort     FailureMode = "abort"
	FailureModeKill      FailureMode = "kill"
)

// Value returns a string representation of the failure mode.
func (m FailureMode) Value() string {
	return string(m)
}
