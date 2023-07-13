package appbuilder

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

type Builder struct {
	dirPath string
	modName string
}

func NewBuilder(dirPath, modName string) *Builder {
	return &Builder{
		dirPath: dirPath,
		modName: modName,
	}
}

func (b *Builder) ExecuteCommands() (string, error) {
	dir := filepath.Join(b.dirPath, b.modName)
	log.Printf("building the mock api in %s", dir)

	// Command 1: go mod init
	log.Println("initialising the go package")
	goModInit := exec.Command("go", "mod", "init", b.modName)
	goModInit.Dir = dir
	_, err := goModInit.Output()
	if err != nil {
		return "", fmt.Errorf("failed to initialize module: %w", err)
	}

	// Command 2: go mod tidy
	log.Println("fetching dependencies")
	goModTidy := exec.Command("go", "mod", "tidy")
	goModTidy.Dir = dir
	_, err = goModTidy.Output()
	if err != nil {
		return "", fmt.Errorf("failed to tidy module dependencies: %w", err)
	}

	// Run "go fmt ./..."
	log.Println("tidying generated code")
	goFmt := exec.Command("go", "fmt", "./...")
	goFmt.Dir = dir

	_, err = goFmt.Output()
	if err != nil {
		return "", fmt.Errorf("failed to format go code: %w", err)
	}

	// Go build
	log.Println("running go build to ensure everything compiles")
	goBuild := exec.Command("go", "build", "./cmd/server/main.go")
	goBuild.Dir = dir
	_, err = goBuild.Output()
	if err != nil {
		return "", fmt.Errorf("failed to build the packages: %w", err)
	}

	// Docker compose up
	log.Println("spinning up mock-api-server in docker container")
	dockerComposeUp := exec.Command("docker-compose", "up", "-d", "--build")
	dockerComposeUp.Dir = dir
	_, err = dockerComposeUp.Output()
	if err != nil {
		return "", fmt.Errorf("failed to build the docker image(s): %w", err)
	}
	return dir, nil
}
