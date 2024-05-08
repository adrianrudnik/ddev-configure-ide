package ddev

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"os/exec"
	"strings"
)

func MustDescribeConfig(workingDirectory string) *DescribeResult {
	// Can we find anything that resembles a ddev executable?
	_, err := exec.LookPath("ddev")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not find ddev executable.")
	}

	// Execute ddev describe with JSON as output, so we can unmarshal it
	describeCmd := exec.Command("ddev", "describe", "-j")
	describeCmd.Dir = workingDirectory

	var cmdOut strings.Builder
	describeCmd.Stdout = &cmdOut

	err = describeCmd.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to describe config")
	}

	var d DescribeResult
	err = json.Unmarshal([]byte(cmdOut.String()), &d)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to process `ddev describe -j` JSON output")
	}

	if d.Raw.Status != "running" {
		log.Fatal().Str("router-status", d.Raw.Status).Msg("DDEV status is not running, aborting.")
	}

	return &d
}
