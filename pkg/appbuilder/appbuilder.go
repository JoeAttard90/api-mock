package appbuilder

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Builder struct {
	dirPath            string
	modName            string
	staticResponsePath string
}

func NewBuilder(dirPath, modName, staticResponsePath string) *Builder {
	return &Builder{
		dirPath:            dirPath,
		modName:            modName,
		staticResponsePath: staticResponsePath,
	}
}

func (b *Builder) ExecuteCommands(dockerRun bool, examplesDir string) (string, error) {
	dir := filepath.Join(b.dirPath, b.modName)
	log.Printf("building the mock api in %s", dir)

	// Copy the directory
	if b.staticResponsePath != "" {
		err := copyDir(b.staticResponsePath, fmt.Sprintf("%s/%s", dir, examplesDir))
		if err != nil {
			return "", err
		}
	}

	// go mod init
	log.Println("initialising the go package")
	goModInit := exec.Command("go", "mod", "init", b.modName)
	goModInit.Dir = dir
	_, err := goModInit.Output()
	if err != nil {
		return "", fmt.Errorf("failed to initialize module: %w", err)
	}

	// go mod tidy
	log.Println("running go mod tidy")
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
		return "", fmt.Errorf("failed to build the server: %w", err)
	}

	// Docker compose up
	if dockerRun {
		log.Println("spinning up mock-api-server in docker container")
		dockerComposeUp := exec.Command("docker-compose", "up", "-d", "--build")
		dockerComposeUp.Dir = dir
		_, err = dockerComposeUp.Output()
		if err != nil {
			return "", fmt.Errorf("failed to build the docker image(s): %w", err)
		}
	}

	return dir, nil
}

func copyDir(src string, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	sStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sStat.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Make sure destination directory exists
	err = os.MkdirAll(dst, sStat.Mode())
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(srcFile string, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {

		}
	}(in)

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
