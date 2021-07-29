package systemdconf

import (
	"strconv"
	"time"

	"github.com/gentlemanautomaton/systemdconf/unitvalue"
)

// Unit describes the properties of a systemd unit section.
type Unit struct {
	// Unit Options

	Description              string
	Documentation            []string
	Wants                    []string
	Requires                 []string
	Requisite                []string
	BindsTo                  []string
	PartOf                   []string
	Upholds                  []string
	Conflicts                []string
	Before                   []string
	After                    []string
	OnFailure                []string
	OnSuccess                []string
	PropagatesReloadTo       []string
	ReloadPropagatedFrom     []string
	PropagatesStopTo         []string
	StopPropagatedFrom       []string
	JoinsNamespaceOf         []string
	RequiresMountsFor        []string
	OnFailureJobMode         unitvalue.JobMode
	IgnoreOnIsolate          unitvalue.Bool
	StopWhenUnneeded         unitvalue.Bool
	RefuseManualStart        unitvalue.Bool
	RefuseManualStop         unitvalue.Bool
	AllowIsolate             unitvalue.Bool
	DefaultDependencies      unitvalue.YesNo
	CollectMode              unitvalue.CollectMode
	FailureAction            unitvalue.Action
	SuccessAction            unitvalue.Action
	FailureActionExitStatus  unitvalue.ExitStatus
	SuccessActionExitStatus  unitvalue.ExitStatus
	JobTimeout               time.Duration
	JobRunningTimeout        time.Duration
	JobTimeoutAction         unitvalue.Action
	JobTimeoutRebootArgument string
	StartLimitInterval       time.Duration
	StartLimitBurst          int
	StartLimitAction         unitvalue.Action
	RebootArgument           string
	SourcePath               string

	// Conditions and Assertions

	ConditionPathExists []string
}

// SectionName returns "Unit" as the name of the configuration section.
func (u Unit) SectionName() string {
	return "Unit"
}

// Directives returns the systemd directives for the [Unit] section.
func (u Unit) Directives() Directives {
	var out Directives

	// Unit Options
	out.AddOptional("Description", u.Description)
	out.AddMany("Documentation", u.Documentation)
	out.AddMany("Wants", u.Wants)
	out.AddMany("Requires", u.Requires)
	out.AddMany("Requisite", u.Requisite)
	out.AddMany("BindsTo", u.BindsTo)
	out.AddMany("PartOf", u.PartOf)
	out.AddMany("Upholds", u.Upholds)
	out.AddDelimited("Conflicts", u.Conflicts)
	out.AddMany("Before", u.Before)
	out.AddMany("After", u.After)
	out.AddDelimited("OnFailure", u.OnFailure)
	out.AddDelimited("OnSuccess", u.OnSuccess)
	out.AddDelimited("PropagatesReloadTo", u.PropagatesReloadTo)
	out.AddDelimited("ReloadPropagatedFrom", u.ReloadPropagatedFrom)
	out.AddDelimited("PropagatesStopTo", u.PropagatesStopTo)
	out.AddDelimited("StopPropagatedFrom", u.StopPropagatedFrom)
	out.AddDelimited("JoinsNamespaceOf", u.JoinsNamespaceOf)
	out.AddDelimited("RequiresMountsFor", u.RequiresMountsFor)
	out.AddOptional("OnFailureJobMode", u.OnFailureJobMode.Value())
	out.AddOptional("IgnoreOnIsolate", u.IgnoreOnIsolate.Value())
	out.AddOptional("StopWhenUnneeded", u.StopWhenUnneeded.Value())
	out.AddOptional("RefuseManualStart", u.RefuseManualStop.Value())
	out.AddOptional("RefuseManualStop", u.RefuseManualStop.Value())
	out.AddOptional("AllowIsolate", u.AllowIsolate.Value())
	out.AddOptional("DefaultDependencies", u.DefaultDependencies.Value())
	out.AddOptional("CollectMode", u.CollectMode.Value())
	out.AddOptional("FailureAction", u.FailureAction.Value())
	out.AddOptional("SuccessAction", u.SuccessAction.Value())
	out.AddOptional("FailureActionExitStatus", u.FailureActionExitStatus.Value())
	out.AddOptional("SuccessActionExitStatus", u.SuccessActionExitStatus.Value())
	if u.JobTimeout > 0 {
		out.AddOptional("JobTimeoutSec", u.JobTimeout.String())
	}
	if u.JobRunningTimeout > 0 {
		out.AddOptional("JobRunningTimeoutSec", u.JobRunningTimeout.String())
	}
	out.AddOptional("JobTimeoutAction", u.JobTimeoutAction.Value())
	out.AddOptional("JobTimeoutRebootArgument", u.JobTimeoutRebootArgument)
	if u.StartLimitInterval > 0 {
		out.AddOptional("StartLimitIntervalSec", u.StartLimitInterval.String())
	}
	if u.StartLimitBurst > 0 {
		out.AddOptional("StartLimitBurst", strconv.Itoa(u.StartLimitBurst))
	}
	out.AddOptional("StartLimitAction", u.StartLimitAction.Value())
	out.AddOptional("RebootArgument", u.RebootArgument)
	out.AddOptional("SourcePath", u.SourcePath)

	// Conditions and Assertions
	out.AddMany("ConditionPathExists", u.ConditionPathExists)

	return out
}
