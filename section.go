package systemdconf

import "io"

// Section is a systemd configuration section.
type Section interface {
	SectionName() string
	Directives() Directives
}

// WriteSection marshals the section configuration and writes it to w.
func WriteSection(w io.Writer, s Section) (n int, err error) {
	// Write the section header
	if written, err := io.WriteString(w, "["+s.SectionName()+"]\n"); err != nil {
		return n + written, err
	} else {
		n += written
	}

	// Write each of the unit directives
	for _, directive := range s.Directives() {
		if written, err := io.WriteString(w, directive.String()+"\n"); err != nil {
			return n + written, err
		} else {
			n += written
		}
	}

	return n, nil
}

// WriteSections marshals a set of configuration sections and writes their
// configuration to w.
func WriteSections(w io.Writer, sections ...Section) (n int, err error) {
	for i, section := range sections {
		// Write an extra newline between sections
		if i > 0 {
			if written, err := io.WriteString(w, "\n"); err != nil {
				return n + written, err
			} else {
				n += written
			}
		}

		// Ask the section to write its contents
		if written, err := WriteSection(w, section); err != nil {
			return n + written, err
		} else {
			n += written
		}
	}
	return n, nil
}
