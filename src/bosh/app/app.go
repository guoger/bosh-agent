package app

import (
	boshagent "bosh/agent"
	boshboot "bosh/bootstrap"
	boshinf "bosh/infrastructure"
	boshmbus "bosh/mbus"
	boshplatform "bosh/platform"
	boshdisk "bosh/platform/disk"
	boshsys "bosh/system"
	"flag"
	"io/ioutil"
)

type App struct {
}

type options struct {
	InfrastructureName string
	PlatformName       string
}

func New() (app App) {
	return
}

func (app App) Run(args []string) (err error) {
	fs := boshsys.OsFileSystem{}
	runner := boshsys.ExecCmdRunner{}
	partitioner := boshdisk.NewSfdiskPartitioner(runner)

	opts, err := parseOptions(args)
	if err != nil {
		return
	}

	infProvider := boshinf.NewProvider()
	infrastructure, err := infProvider.Get(opts.InfrastructureName)
	if err != nil {
		return
	}

	platformProvider := boshplatform.NewProvider(fs, runner, partitioner)
	platform, err := platformProvider.Get(opts.PlatformName)
	if err != nil {
		return
	}

	boot := boshboot.New(fs, infrastructure, platform)
	settings, err := boot.Run()
	if err != nil {
		return
	}

	mbusHandlerProvider := boshmbus.NewHandlerProvider(settings)
	mbusHandler, err := mbusHandlerProvider.Get()
	if err != nil {
		return
	}

	agent := boshagent.New(settings, mbusHandler, platform)
	err = agent.Run()
	return
}

func parseOptions(args []string) (opts options, err error) {
	flagSet := flag.NewFlagSet("bosh-agent-args", flag.ContinueOnError)
	flagSet.SetOutput(ioutil.Discard)
	flagSet.StringVar(&opts.InfrastructureName, "I", "", "Set Infrastructure")
	flagSet.StringVar(&opts.PlatformName, "P", "", "Set Platform")

	err = flagSet.Parse(args[1:])
	return
}