package unitvalue

// RestartMode defines a unit's state transition behavior when it is
// restarted.
type RestartMode string

// Possible restart modes.
const (
	RestartModeNormal RestartMode = "normal"
	RestartModeDirect RestartMode = "direct"
	RestartModeDebug  RestartMode = "debug"
)

// Value returns a string representation of the restart mode.
func (m RestartMode) Value() string {
	return string(m)
}
