package generate

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var Command = &cobra.Command{
	Use:     "generate",
	Short:   "Generate files",
	Aliases: []string{"g"},
	Args:    cobra.ExactArgs(2),
	Run:     executeCommand,
}

//go:embed _skeleton
var skeletonFiles embed.FS

func executeCommand(cmd *cobra.Command, args []string) {
	generationType := args[0]
	name := args[1]

	var err error

	switch generationType {
	case "c", "component":
		err = generateComponent(name)
		if err != nil {
			break
		}
		fmt.Printf("Component '%s' successfully generated\n", name)
		return
	default:
		err = fmt.Errorf("'%s' is not a valid target to generate", generationType)
	}

	fmt.Printf("%sError: %s%s\n", chalk.Red, err.Error(), chalk.Reset)
	os.Exit(1)
}
