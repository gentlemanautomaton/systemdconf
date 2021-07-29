package unitvalue

// Action is an action that can be taken on a unit when it fails or becomes
// inactive.
type Action string

// Possible actions.
const (
	ActionNone              Action = "none"
	ActionReboot            Action = "reboot"
	ActionRebootForce       Action = "reboot-force"
	ActionRebootImmediate   Action = "reboot-immediate"
	ActionPoweroff          Action = "poweroff"
	ActionPoweroffForce     Action = "poweroff-force"
	ActionPoweroffImmediate Action = "poweroff-immediate"
	ActionExit              Action = "exit"
	ActionExitForce         Action = "exit-force"
)

// Value returns a string representation of the action.
func (a Action) Value() string {
	return string(a)
}
