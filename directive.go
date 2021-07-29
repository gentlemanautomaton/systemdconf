package systemdconf

import "strings"

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

// AddOptional appends a directive to the list with the given name and value.
//
// If no values is supplied, the directive will not be added.
func (list *Directives) AddOptional(name, value string) {
	if value != "" {
		(*list) = append(*list, Directive{Name: name, Value: value})
	}
}

// AddDelimited appends a single directive to the list with the given name
// and values.
//
// If more than one value is supplied, the values will be space delimited.
//
// If no values are supplied, the directive will not be added.
func (list *Directives) AddDelimited(name string, values []string) {
	if len(values) > 0 {
		list.Add(name, strings.Join(values, " "))
	}
}

// AddMany appends a separate directive for each value in the values list.
func (list *Directives) AddMany(name string, values []string) {
	for _, value := range values {
		list.Add(name, value)
	}
}
