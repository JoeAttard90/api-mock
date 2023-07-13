package dockerutils

import "api-mock/pkg/templateutils"

type DockerFileGenerator struct {
	port                      string
	dockerFileTemplatePath    string
	dockerFileOutputPath      string
	dockerComposeTemplatePath string
	dockerComposeOutputPath   string
}

func NewDockerFileGenerator(
	port,
	dockerFileTemplatePath,
	dockerFileOutputPath,
	dockerComposeTemplatePath,
	dockerComposeOutputPath string,
) *DockerFileGenerator {
	return &DockerFileGenerator{
		port:                      port,
		dockerFileTemplatePath:    dockerFileTemplatePath,
		dockerFileOutputPath:      dockerFileOutputPath,
		dockerComposeTemplatePath: dockerComposeTemplatePath,
		dockerComposeOutputPath:   dockerComposeOutputPath,
	}
}

func (dfg *DockerFileGenerator) Generate() error {
	// Generate Dockerfile
	err := templateutils.CreateTemplate(
		dfg.dockerFileTemplatePath,
		dfg.dockerFileOutputPath,
		dfg.port,
	)
	if err != nil {
		return err
	}
	// Generate Docker compose
	err = templateutils.CreateTemplate(
		dfg.dockerComposeTemplatePath,
		dfg.dockerComposeOutputPath,
		dfg.port,
	)
	if err != nil {
		return err
	}
	return nil
}
