package systemdconf

// Unit describes the properties of a systemd unit section.
type Unit struct {
	Description         string
	ConditionPathExists []string
	Before              []string
	Wants               []string
}

// SectionName returns "Unit" as the name of the configuration section.
func (u Unit) SectionName() string {
	return "Unit"
}

// Directives returns the systemd directives for the [Unit] section.
func (u Unit) Directives() Directives {
	var out Directives
	if u.Description != "" {
		out.Add("Description", u.Description)
	}
	for _, condition := range u.ConditionPathExists {
		out.Add("ConditionPathExists", condition)
	}
	for _, before := range u.Before {
		out.Add("Before", before)
	}
	for _, wants := range u.Wants {
		out.Add("Wants", wants)
	}
	return out
}
