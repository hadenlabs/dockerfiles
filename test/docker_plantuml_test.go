package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestPlantumlLatestSuccess(t *testing.T) {
	tag := "hadenlabs/plantuml:latest"
	otherOptions := []string{
		"--no-cache",
	}
	buildOptions := &docker.BuildOptions{
		Tags:         []string{tag},
		OtherOptions: otherOptions,
	}

	docker.Build(t, "../plantuml", buildOptions)

	opts := &docker.RunOptions{
		Command: []string{},
	}
	output := docker.Run(t, tag, opts)
	assert.NotEmpty(t, output, output)
}
