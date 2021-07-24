package systemdconf

// Service describes the properties of a systemd service section.
type Service struct {
	Type             string
	RemainAfterExit  bool
	EnvironmentFiles []string
	ExecStart        string
	ExecStop         string
}

// SectionName returns "Service" as the name of the configuration section.
func (s Service) SectionName() string {
	return "Service"
}

// Directives returns the systemd directives for the [Service] section.
func (s Service) Directives() Directives {
	var out Directives
	if s.Type != "" {
		out.Add("Type", s.Type)
	}
	if s.RemainAfterExit {
		out.Add("RemainAfterExit", "true")
	}
	for _, path := range s.EnvironmentFiles {
		out.Add("EnvironmentFile", path)
	}
	if s.ExecStart != "" {
		out.Add("ExecStart", s.ExecStart)
	}
	if s.ExecStop != "" {
		out.Add("ExecStop", s.ExecStop)
	}
	return out
}
