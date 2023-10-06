package initialize

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"text/template"

	"github.com/spf13/cobra"
)

var directoryPerm fs.FileMode = 0755
var filePerm fs.FileMode = 0644

//go:embed skeleton
var files embed.FS

func ExecuteCommand(cmd *cobra.Command, args []string) {
	projectName := args[0]

	if _, err := os.Stat(projectName); !os.IsNotExist(err) {
		fmt.Println("Directory already exists!")
		os.Exit(1)
	}

	fmt.Println("Creating project dir")

	if err := os.Mkdir(projectName, directoryPerm); err != nil {
		fmt.Printf("Failed to create project directory: %v", err)
		os.Exit(1)
	}

	if err := os.Chdir(projectName); err != nil {
		fmt.Printf("Failed to change working directory to %s", projectName)
		os.Exit(1)
	}

	fmt.Println("Creating project skeleton")

	if err := createFilesFromSkeleton(files, "skeleton", projectName); err != nil {
		fmt.Printf("Failed to create file skeleton: %v", err)
		os.Exit(1)
	}

	goModInitCmd := exec.Command("go", "mod", "init", projectName)
	goModTidyCmd := exec.Command("go", "mod", "tidy")

	goModInitCmd.Run()
	goModTidyCmd.Run()
}

func createFilesFromSkeleton(files fs.FS, rootDir string, moduleName string) error {
	return fs.WalkDir(files, rootDir, func(path string, d fs.DirEntry, _ error) error {
		if path == rootDir {
			return nil
		}

		destinationPath := path[len(rootDir)+1:]

		if d.IsDir() {
			if err := os.Mkdir(destinationPath, directoryPerm); err != nil {
				return fmt.Errorf("failed to create dir '%s': %w", path, err)
			}
			return nil
		}

		sourceFile, err := files.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open source file '%s': %w", path, err)
		}

		bytes, _ := io.ReadAll(sourceFile)

		t := template.New("template-file")
		t, err = t.Parse(string(bytes))
		if err != nil {
			return fmt.Errorf("failed to parse template: %w", err)
		}

		destinationFile, err := os.OpenFile(destinationPath, os.O_CREATE|os.O_WRONLY, filePerm)
		if err != nil {
			return fmt.Errorf("failed to open destination file '%s': %w", destinationPath, err)
		}

		if err := t.Execute(destinationFile, map[string]interface{}{
			"ModuleName": moduleName,
		}); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}

		return nil
	})
}
