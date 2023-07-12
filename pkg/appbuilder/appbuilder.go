package appbuilder

import (
	"fmt"
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

	// Command 1: go mod init
	goModInit := exec.Command("go", "mod", "init", b.modName)
	goModInit.Dir = dir
	_, err := goModInit.Output()
	if err != nil {
		return "", fmt.Errorf("failed to initialize module: %w", err)
	}

	// Command 2: go mod tidy
	goModTidy := exec.Command("go", "mod", "tidy")
	goModTidy.Dir = dir
	_, err = goModTidy.Output()
	if err != nil {
		return "", fmt.Errorf("failed to tidy module dependencies: %w", err)
	}

	// Run "go fmt ./..."
	goFmt := exec.Command("go", "fmt", "./...")
	goFmt.Dir = dir

	_, err = goFmt.Output()
	if err != nil {
		return "", fmt.Errorf("failed to format go code: %w", err)
	}

	// Go build
	goBuild := exec.Command("go", "build", "./cmd/server/main.go")
	goBuild.Dir = dir
	_, err = goBuild.Output()
	if err != nil {
		return "", fmt.Errorf("failed to build the packages: %w", err)
	}
	return dir, nil
}
