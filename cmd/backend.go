package main

import (
	"github.com/grafana/grafana-build/pipelines"
	"github.com/urfave/cli/v2"
)

var TestBackendUnit = &cli.Command{
	Name:   "test",
	Action: PipelineAction(pipelines.GrafanaBackendTests),
}

var TestBackendIntegration = &cli.Command{
	Name:   "test-integration",
	Action: PipelineAction(pipelines.GrafanaBackendTestIntegration),
}

var BuildBackend = &cli.Command{
	Name:   "build",
	Action: PipelineAction(pipelines.GrafanaBackendBuild),
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "version",
			Required: false,
		},
	},
}
