package commands

import (
	"fmt"
	"github.com/dnotez/dnotez-cli/dsl/article"
	"github.com/spf13/cobra"
	//"github.com/wsxiaoys/terminal"
)

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a note",
	Long: `Remove a note by id.
Examples:
dz rm f5cda30e-3d0d-4d4c-b593-4c0a8befa4ef`,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig()
		if len(args) < 1 {
			cmd.Help()
		} else {
			remove(args)
		}
	},
}

func remove(args []string) {
	_, duration, err := article.Remove(args[0])
	if err != nil {
		fmt.Println("Error in remove: ", err)
		return
	}
	if !Quiet {
		fmt.Printf("Removed in %0.2fs\n", duration.Seconds())
	}

}
