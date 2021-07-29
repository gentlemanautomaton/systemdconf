package unitvalue

// Restart indicates whether a service process should be restarted when it
// exits.
type Restart string

// Possible collection modes.
const (
	RestartNo         Restart = "no"
	RestartOnSuccess  Restart = "on-success"
	RestartOnFailure  Restart = "on-failure"
	RestartOnAbnormal Restart = "on-abnormal"
	RestartOnWatchdog Restart = "on-watchdod"
	RestartOnAbort    Restart = "on-abort"
	RestartAlways     Restart = "always"
)

// Value returns a string representation of the restart value.
func (r Restart) Value() string {
	return string(r)
}
