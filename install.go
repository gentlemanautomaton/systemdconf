package systemdconf

// Install describes the properties of a systemd installation section.
type Install struct {
	Alias           []string
	WantedBy        []string
	RequiredBy      []string
	UpheldBy        []string
	Also            []string
	DefaultInstance string
}

// SectionName returns "Install" as the name of the configuration section.
func (i Install) SectionName() string {
	return "Install"
}

// Directives returns the systemd directives for the [Install] section.
func (i Install) Directives() Directives {
	var out Directives
	for _, alias := range i.Alias {
		out.Add("Alias", alias)
	}
	for _, wantedBy := range i.WantedBy {
		out.Add("WantedBy", wantedBy)
	}
	for _, requiredBy := range i.RequiredBy {
		out.Add("RequiredBy", requiredBy)
	}
	for _, upheldBy := range i.UpheldBy {
		out.Add("UpheldBy", upheldBy)
	}
	for _, also := range i.Also {
		out.Add("Also", also)
	}
	if i.DefaultInstance != "" {
		out.Add("DefaultInstance", i.DefaultInstance)
	}
	return out
}
