package unitvalue

// SystemProtection indicates the protection level for system files.
type SystemProtection string

// Possible system protection modes.
const (
	SystemProtectionBasic  SystemProtection = "on"
	SystemProtectionFull   SystemProtection = "full"
	SystemProtectionStrict SystemProtection = "strict"
)

// Value returns a string representation of the system protection mode.
func (m SystemProtection) Value() string {
	return string(m)
}

// HomeProtection indicates the protection level for home files.
type HomeProtection string

// Possible home protection modes.
const (
	HomeProtectionRestricted HomeProtection = "on"
	HomeProtectionReadOnly   HomeProtection = "read-only"
	HomeProtectionTmpFS      HomeProtection = "tmpfs"
)

// Value returns a string representation of the home protection mode.
func (m HomeProtection) Value() string {
	return string(m)
}
