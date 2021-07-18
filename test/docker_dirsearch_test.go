package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDirsearchLatestSuccess(t *testing.T) {

	tag := "hadenlabs/dirsearch:latest"
	otherOptions := []string{
		"--no-cache",
	}
	buildOptions := &docker.BuildOptions{
		Tags:         []string{tag},
		OtherOptions: otherOptions,
	}

	docker.Build(t, "../dirsearch", buildOptions)

	opts := &docker.RunOptions{
		Command: []string{},
	}
	output := docker.Run(t, tag, opts)
	assert.NotEmpty(t, output, output)
}
