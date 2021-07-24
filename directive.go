package systemdconf

// Directive holds a single systemd directive within a unit file.
type Directive struct {
	Name  string
	Value string
}

// String returns a string representation of the directive.
func (d Directive) String() string {
	return d.Name + "=" + d.Value
}

// Directives is a slice of systemd directives.
type Directives []Directive

// Add appends a directive to the list with the given name and value.
func (list *Directives) Add(name, value string) {
	(*list) = append(*list, Directive{Name: name, Value: value})
}
