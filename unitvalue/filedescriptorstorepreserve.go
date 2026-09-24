package unitvalue

// FileDescriptorStorePreserve controls the preservation of a service's file
// descriptor store.
type FileDescriptorStorePreserve string

// Possible file descriptor store preservation mechanisms.
const (
	FileDescriptorStorePreserveNo        FileDescriptorStorePreserve = "no"
	FileDescriptorStorePreserveYes       FileDescriptorStorePreserve = "yes"
	FileDescriptorStorePreserveRestart   FileDescriptorStorePreserve = "restart"
	FileDescriptorStorePreserveOnSuccess FileDescriptorStorePreserve = "on-success"
)

// Value returns a string representation of the file descriptor store
// preservation.
func (p FileDescriptorStorePreserve) Value() string {
	return string(p)
}
