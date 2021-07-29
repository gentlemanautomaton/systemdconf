package systemdconf_test

import (
	"bytes"
	"fmt"

	"github.com/gentlemanautomaton/systemdconf"
	"github.com/gentlemanautomaton/systemdconf/unitvalue"
)

func Example() {
	var buffer bytes.Buffer
	systemdconf.WriteSections(&buffer,
		systemdconf.Unit{
			Description: "Test Unit",
			ConditionPathExists: []string{
				"/usr/lib/test",
				"/etc/test/env1.conf",
			},
			Before: []string{"network-pre.target"},
			Wants:  []string{"network-pre.target"},
		},
		systemdconf.Service{
			Type:            "oneshot",
			RemainAfterExit: unitvalue.True,
			EnvironmentFiles: []string{
				"/usr/lib/test/env1.conf",
				"/etc/test/env1.conf",
			},
			ExecStart: []string{"/usr/lib/test start"},
			ExecStop:  []string{"/usr/lib/test stop"},
		},
		systemdconf.Install{
			WantedBy: []string{"multi-user.target"},
		},
	)

	fmt.Println(buffer.String())

	// Output:
	// [Unit]
	// Description=Test Unit
	// Wants=network-pre.target
	// Before=network-pre.target
	// ConditionPathExists=/usr/lib/test
	// ConditionPathExists=/etc/test/env1.conf
	//
	// [Service]
	// Type=oneshot
	// RemainAfterExit=true
	// ExecStart=/usr/lib/test start
	// ExecStop=/usr/lib/test stop
	// EnvironmentFile=/usr/lib/test/env1.conf
	// EnvironmentFile=/etc/test/env1.conf
	//
	// [Install]
	// WantedBy=multi-user.target
}
