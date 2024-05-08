package config

import "github.com/adrianrudnik/ddev-configure-ide/internal/ddev"

type RuntimeConfig struct {
	DryRun           bool
	WorkingDirectory string
	DDEVConfig       *ddev.DescribeResult
}
