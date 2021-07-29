package unitvalue

// JobMode indicates the job mode used by units activated on failure.
type JobMode string

// Possible job modes.
const (
	JobModeFail                = "fail"
	JobModeReplace             = "replace"
	JobModeReplaceIrreversibly = "replace-irreversibly"
	JobModeIsolate             = "isolate"
	JobModeFlush               = "flush"
	JobModeIgnoreDependencies  = "ignore-dependencies"
	JobModeIgnoreRequirements  = "ignore-requirements"
)

// Value returns a string representation of the job mode.
func (m JobMode) Value() string {
	return string(m)
}
