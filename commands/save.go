package commands

import (
	"fmt"
	"strings"

	"github.com/dnotez/dnotez-cli/dsl/cmd"
	"github.com/spf13/cobra"
	//"github.com/wsxiaoys/terminal"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save a note",
	Long: `Save the command line as a note.
Input hash will be used for finding duplicates.

Examples:
dz save "apt-get update; apt-get upgrade"
dz save -l update-docker "curl -sSL https://get.docker.com/ubuntu/ | sudo sh"`,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig()
		if len(args) < 1 {
			cmd.Help()
		} else {
			save(args)
		}
	},
}

var saveLabel string

func init() {
	saveCmd.Flags().StringVarP(&saveLabel, "label", "l", "", "Set note label for quick fetch")
}

func save(args []string) {
	body := strings.Join(args, " ")
	request := cmd.SaveCmdRequest{Body: body, Label: saveLabel}
	response, duration, err := request.Submit()
	if err != nil {
		fmt.Println("Error in save: ", err)
		return
	}
	if !Quiet {
		fmt.Printf("Saved in %0.2fs\n\n", duration.Seconds())
	}

	fmt.Printf("%s\n", (*response).URL)
}
