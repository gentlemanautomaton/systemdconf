package systemdconf_test

import (
	"fmt"

	"github.com/gentlemanautomaton/systemdconf"
)

func ExampleEscape() {
	names := []string{"Hallöchen, Meister", ".One/Two"}
	for _, name := range names {
		fmt.Println(systemdconf.Escape(name))
	}

	// Output:
	// Hall\xc3\xb6chen\x2c\x20Meister
	// \x2eOne-Two
}
