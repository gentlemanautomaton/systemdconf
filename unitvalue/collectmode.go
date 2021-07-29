package unitvalue

// CollectMode influences a unit's garbage collection.
type CollectMode string

// Possible collection modes.
const (
	CollectModeInactive         CollectMode = "inactive"
	CollectModeInactiveOrFailed CollectMode = "inactive-or-failed"
)

// Value returns a string representation of the collect mode.
func (m CollectMode) Value() string {
	return string(m)
}
