package drain

import (
	"strconv"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-agent/internal/github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-agent/internal/github.com/cloudfoundry/bosh-utils/system"
)

type ConcreteScript struct {
	fs              boshsys.FileSystem
	runner          boshsys.CmdRunner
	drainScriptPath string
}

func NewConcreteScript(
	fs boshsys.FileSystem,
	runner boshsys.CmdRunner,
	drainScriptPath string,
) (script ConcreteScript) {
	script = ConcreteScript{
		fs:              fs,
		runner:          runner,
		drainScriptPath: drainScriptPath,
	}
	return
}

func (script ConcreteScript) Exists() bool {
	return script.fs.FileExists(script.drainScriptPath)
}

func (script ConcreteScript) Path() string {
	return script.drainScriptPath
}

func (script ConcreteScript) Run(params ScriptParams) (int, error) {
	jobChange := params.JobChange()
	hashChange := params.HashChange()
	updatedPkgs := params.UpdatedPackages()

	command := boshsys.Command{
		Name: script.drainScriptPath,
		Env: map[string]string{
			"PATH": "/usr/sbin:/usr/bin:/sbin:/bin",
		},
	}

	jobState, err := params.JobState()
	if err != nil {
		return 0, bosherr.WrapError(err, "Getting job state")
	}

	if jobState != "" {
		command.Env["BOSH_JOB_STATE"] = jobState
	}

	jobNextState, err := params.JobNextState()
	if err != nil {
		return 0, bosherr.WrapError(err, "Getting job next state")
	}

	if jobNextState != "" {
		command.Env["BOSH_JOB_NEXT_STATE"] = jobNextState
	}

	command.Args = append(command.Args, jobChange, hashChange)
	command.Args = append(command.Args, updatedPkgs...)

	stdout, _, _, err := script.runner.RunComplexCommand(command)
	if err != nil {
		return 0, bosherr.WrapError(err, "Running drain script")
	}

	value, err := strconv.Atoi(strings.TrimSpace(stdout))
	if err != nil {
		return 0, bosherr.WrapError(err, "Script did not return a signed integer")
	}

	return value, nil
}
