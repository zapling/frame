package generate

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/zapling/frame/internal/file"
)

var ComponentHandler bool

func generateComponent(name string) error {
	if err := validateName(name); err != nil {
		return fmt.Errorf(
			"'%s' is not allowed as the component name\n"+
				"Only a-z and _ is allowed as the component name",
			name,
		)
	}

	if err := os.Mkdir(name, file.ModeDirectory); err != nil {
		return fmt.Errorf("directory '%s' could not be created: %v", name, err)
	}

	if err := createComponentTemplate(name); err != nil {
		return err
	}

	if ComponentHandler {
		if err := createComponentHandler(name); err != nil {
			return err
		}
	}

	return nil
}

func createComponentTemplate(name string) error {
	fileContent, err := file.ReadContent(skeletonFiles, "_skeleton/component/component.templ")
	if err != nil {
		return err
	}

	fileContent = bytes.ReplaceAll(
		fileContent,
		[]byte("package component"),
		[]byte("package "+name),
	)

	fileName := fmt.Sprintf("%s.templ", name)
	filePath := fmt.Sprintf("%s/%s", name, fileName)

	destinationFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, file.ModeFile)
	if err != nil {
		return fmt.Errorf("could not open new file '%s': %v", filePath, err)
	}

	defer destinationFile.Close()

	_, err = destinationFile.Write(fileContent)
	if err != nil {
		return fmt.Errorf("could not write content to file '%s': %v", filePath, err)
	}

	templGenerateCmd := exec.Command("templ", "generate", "-f", filePath)
	if err := templGenerateCmd.Run(); err != nil {
		// TODO: get stdrr from command
		return fmt.Errorf("could not run 'templ generate' on file: %v", err)
	}

	return nil
}

func createComponentHandler(name string) error {
	fileContent, err := file.ReadContent(skeletonFiles, "_skeleton/component/handler.go")
	if err != nil {
		return err
	}

	fileContent = bytes.ReplaceAll(
		fileContent,
		[]byte("package component"),
		[]byte("package "+name),
	)

	filePath := fmt.Sprintf("%s/handler.go", name)

	destinationFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, file.ModeFile)
	if err != nil {
		return fmt.Errorf("could not open new file '%s': %v", filePath, err)
	}

	defer destinationFile.Close()

	_, err = destinationFile.Write(fileContent)
	if err != nil {
		return fmt.Errorf("could not write content to file '%s': %v", filePath, err)
	}

	return nil
}

var nameRegex = regexp.MustCompile(`^[a-z\_]+$`)

func validateName(name string) error {
	isValid := nameRegex.MatchString(name)
	if !isValid {
		return fmt.Errorf("name does not match regex: %s", nameRegex.String())
	}

	return nil
}
