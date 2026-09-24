package systemdconf

import (
	"strconv"
	"time"

	"github.com/gentlemanautomaton/systemdconf/unitvalue"
)

// Service describes the properties of a systemd service section.
type Service struct {
	// Service Options

	Type                          string
	ExitType                      unitvalue.ExitType
	RemainAfterExit               unitvalue.Bool
	GuessMainPID                  unitvalue.Bool
	PIDFile                       string
	BusName                       string
	ExecStart                     []string
	ExecStartPre                  []string
	ExecStartPost                 []string
	ExecCondition                 []string
	ExecReload                    []string
	ExecReloadPost                []string
	ExecStop                      []string
	ExecStopPost                  []string
	RestartDelay                  time.Duration
	RestartSteps                  int
	RestartMaxDelay               time.Duration
	RestartRandomizedDelay        time.Duration
	TimeoutStart                  time.Duration
	TimeoutStop                   time.Duration
	TimeoutAbort                  time.Duration
	Timeout                       time.Duration
	TimeoutStartFailureMode       unitvalue.FailureMode
	TimeoutStopFailureMode        unitvalue.FailureMode
	RuntimeMax                    time.Duration
	RuntimeRandomizedExtra        time.Duration
	WatchdogTimeout               time.Duration
	Restart                       unitvalue.Restart
	RestartMode                   unitvalue.RestartMode
	SuccessExitStatus             []string
	RestartPreventExitStatus      []string
	RestartForceExitStatus        []string
	RootDirectoryStartOnly        unitvalue.Bool
	NonBlocking                   unitvalue.Bool
	NotifyAccess                  unitvalue.NotifyAccess
	SocketUnits                   []string
	FileDescriptorStoreMax        int
	FileDescriptorStorePreserve   unitvalue.FileDescriptorStorePreserve
	LiveUpdateOrchestratorSession []string
	USBFunctionDescriptorsFile    string
	USBFunctionStringsFile        string
	OutOfMemoryPolicy             unitvalue.OutOfMemoryPolicy
	OpenFile                      []string
	ReloadSignal                  string
	RefreshOnReload               string

	// Kill Options

	KillMode          string
	KillSignal        string
	RestartKillSignal string
	SendSIGHUP        unitvalue.YesNo
	SendSIGKILL       unitvalue.YesNo
	FinalKillSignal   string
	WatchdogSignal    string

	// Environment Options

	Environment      []string
	EnvironmentFiles []string
	PassEnvironment  []string
	UnsetEnvironment []string

	// Execution Options

	WorkingDirectory string
	RootDirectory    string

	// Sandboxing

	ProtectSystem              unitvalue.SystemProtection
	ProtectHome                unitvalue.HomeProtection
	RuntimeDirectories         []string
	StateDirectories           []string
	CacheDirectories           []string
	LogDirectories             []string
	ConfigurationDirectories   []string
	RuntimeDirectoryMode       unitvalue.FileMode
	StateDirectoryMode         unitvalue.FileMode
	CacheDirectoryMode         unitvalue.FileMode
	LogDirectoryMode           unitvalue.FileMode
	ConfigurationDirectoryMode unitvalue.FileMode
}

// SectionName returns "Service" as the name of the configuration section.
func (s Service) SectionName() string {
	return "Service"
}

// Directives returns the systemd directives for the [Service] section.
func (s Service) Directives() Directives {
	var out Directives

	// Service Options
	out.AddOptional("Type", s.Type)
	out.AddOptional("ExitType", s.ExitType.Value())
	out.AddOptional("RemainAfterExit", s.RemainAfterExit.Value())
	out.AddOptional("GuessMainPID", s.GuessMainPID.Value())
	out.AddOptional("PIDFile", s.PIDFile)
	out.AddOptional("BusName", s.BusName)
	out.AddMany("ExecStart", s.ExecStart)
	out.AddMany("ExecStartPre", s.ExecStartPre)
	out.AddMany("ExecStartPost", s.ExecStartPost)
	out.AddMany("ExecCondition", s.ExecCondition)
	out.AddMany("ExecReload", s.ExecReload)
	out.AddMany("ExecReloadPost", s.ExecReloadPost)
	out.AddMany("ExecStop", s.ExecStop)
	out.AddMany("ExecStopPost", s.ExecStopPost)
	if s.RestartDelay > 0 {
		out.Add("RestartSec", s.RestartDelay.String())
	}
	if s.RestartSteps > 0 {
		out.Add("RestartSteps", strconv.Itoa(s.RestartSteps))
	}
	if s.RestartMaxDelay > 0 {
		out.Add("RestartMaxDelaySec", s.RestartMaxDelay.String())
	}
	if s.RestartRandomizedDelay > 0 {
		out.Add("RestartRandomizedDelaySec", s.RestartRandomizedDelay.String())
	}
	if s.TimeoutStart > 0 {
		out.Add("TimeoutStartSec", s.TimeoutStart.String())
	}
	if s.TimeoutStop > 0 {
		out.Add("TimeoutStopSec", s.TimeoutStop.String())
	}
	if s.TimeoutAbort > 0 {
		out.Add("TimeoutAbortSec", s.TimeoutAbort.String())
	}
	if s.Timeout > 0 {
		out.Add("TimeoutSec", s.Timeout.String())
	}
	out.AddOptional("TimeoutStartFailureMode", s.TimeoutStartFailureMode.Value())
	out.AddOptional("TimeoutStopFailureMode", s.TimeoutStopFailureMode.Value())
	if s.RuntimeMax > 0 {
		out.Add("RuntimeMaxSec", s.RuntimeMax.String())
	}
	if s.RuntimeRandomizedExtra > 0 {
		out.Add("RuntimeRandomizedExtraSec", s.RuntimeRandomizedExtra.String())
	}
	if s.WatchdogTimeout > 0 {
		out.Add("WatchdogSec", s.WatchdogTimeout.String())
	}
	out.AddOptional("Restart", s.Restart.Value())
	out.AddOptional("RestartMode", s.RestartMode.Value())
	out.AddMany("SuccessExitStatus", s.SuccessExitStatus)
	out.AddMany("RestartPreventExitStatus", s.RestartPreventExitStatus)
	out.AddMany("RestartForceExitStatus", s.RestartForceExitStatus)
	out.AddOptional("RootDirectoryStartOnly", s.RootDirectoryStartOnly.Value())
	out.AddOptional("NonBlocking", s.NonBlocking.Value())
	out.AddOptional("NotifyAccess", s.NotifyAccess.Value())
	out.AddMany("Sockets", s.SocketUnits)
	if s.FileDescriptorStoreMax > 0 {
		out.Add("FileDescriptorStoreMax", strconv.Itoa(s.FileDescriptorStoreMax))
	}
	out.AddOptional("FileDescriptorStorePreserve", s.FileDescriptorStorePreserve.Value())
	out.AddDelimited("LUOSession", s.LiveUpdateOrchestratorSession)
	out.AddOptional("USBFunctionDescriptors", s.USBFunctionDescriptorsFile)
	out.AddOptional("USBFunctionStrings", s.USBFunctionStringsFile)
	out.AddOptional("OOMPolicy", s.OutOfMemoryPolicy.Value())
	out.AddMany("OpenFile", s.OpenFile)
	out.AddOptional("ReloadSignal", s.ReloadSignal)
	out.AddOptional("RefreshOnReload", s.RefreshOnReload)

	// Kill Options
	out.AddOptional("KillMode", s.KillMode)
	out.AddOptional("KillSignal", s.KillSignal)
	out.AddOptional("RestartKillSignal", s.RestartKillSignal)
	out.AddOptional("SendSIGHUP", s.SendSIGHUP.Value())
	out.AddOptional("SendSIGKILL", s.SendSIGKILL.Value())
	out.AddOptional("FinalKillSignal", s.FinalKillSignal)
	out.AddOptional("WatchdogSignal", s.WatchdogSignal)

	// Environment Options
	out.AddMany("Environment", s.Environment)
	out.AddMany("EnvironmentFile", s.EnvironmentFiles)
	out.AddMany("PassEnvironment", s.PassEnvironment)
	out.AddMany("UnsetEnvironment", s.UnsetEnvironment)

	// Execution Options
	out.AddOptional("WorkingDirectory", s.WorkingDirectory)
	out.AddOptional("RootDirectory", s.RootDirectory)

	// Sandboxing Options
	out.AddOptional("ProtectSystem", s.ProtectSystem.Value())
	out.AddOptional("ProtectHome", s.ProtectHome.Value())
	out.AddMany("RuntimeDirectory", s.RuntimeDirectories)
	out.AddMany("StateDirectory", s.StateDirectories)
	out.AddMany("CacheDirectory", s.CacheDirectories)
	out.AddMany("LogsDirectory", s.LogDirectories)
	out.AddMany("ConfigurationDirectory", s.ConfigurationDirectories)
	out.AddOptional("RuntimeDirectoryMode", s.RuntimeDirectoryMode.Value())
	out.AddOptional("StateDirectoryMode", s.StateDirectoryMode.Value())
	out.AddOptional("CacheDirectoryMode", s.CacheDirectoryMode.Value())
	out.AddOptional("LogsDirectoryMode", s.LogDirectoryMode.Value())
	out.AddOptional("ConfigurationDirectoryMode", s.ConfigurationDirectoryMode.Value())

	return out
}
