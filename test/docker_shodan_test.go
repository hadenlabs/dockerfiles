package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestShodanLatestSuccess(t *testing.T) {
	tag := "hadenlabs/shodan:latest"
	otherOptions := []string{
		"--no-cache",
	}
	buildOptions := &docker.BuildOptions{
		Tags:         []string{tag},
		OtherOptions: otherOptions,
	}

	docker.Build(t, "../shodan", buildOptions)

	opts := &docker.RunOptions{
		Command: []string{},
	}
	output := docker.Run(t, tag, opts)
	assert.NotEmpty(t, output, output)
}
