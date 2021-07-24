package systemdconf

// Install describes the properties of a systemd installation section.
type Install struct {
	WantedBy []string
}

// SectionName returns "Install" as the name of the configuration section.
func (i Install) SectionName() string {
	return "Install"
}

// Directives returns the systemd directives for the [Install] section.
func (i Install) Directives() Directives {
	var out Directives
	for _, wantedby := range i.WantedBy {
		out.Add("WantedBy", wantedby)
	}
	return out
}
