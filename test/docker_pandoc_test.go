package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestPandocLatestSuccess(t *testing.T) {

	tag := "hadenlabs/pandoc:latest"
	otherOptions := []string{
		"--no-cache",
	}
	buildOptions := &docker.BuildOptions{
		Tags:         []string{tag},
		OtherOptions: otherOptions,
	}

	docker.Build(t, "../pandoc", buildOptions)

	opts := &docker.RunOptions{
		Command: []string{
			"--version",
		},
	}
	output := docker.Run(t, tag, opts)
	assert.NotEmpty(t, output, output)
}
