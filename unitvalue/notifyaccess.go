package unitvalue

// NotifyAccess defines the level of access to a unit's notification socket.
type NotifyAccess string

// Possible notification socket access levels.
const (
	NotifyAccessNone NotifyAccess = "none"
	NotifyAccessMain NotifyAccess = "main"
	NotifyAccessExec NotifyAccess = "exec"
	NotifyAccessAll  NotifyAccess = "all"
)

// Value returns a string representation of the notification socket access
// level.
func (a NotifyAccess) Value() string {
	return string(a)
}
