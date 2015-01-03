package commands

import (
	"fmt"
	"strings"

	"github.com/dnotez/dnotez-cli/dsl/suggestion"
	"github.com/spf13/cobra"
	"github.com/wsxiaoys/terminal"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for `query` in the saved notes",
	Long: `Search for the query string in the saved notes.
If type flag is set, results will be filtered by the type.

Examples:
dz search elasticsearch curl
dz search -t bash for loop	`,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig()
		if len(args) < 1 {
			cmd.Help()
		} else {
			search(args)
		}
	},
}

var queryType string

func init() {
	searchCmd.Flags().StringVarP(&queryType, "type", "t", "", "Type of the note (e.g. bash, article, sof)")
}

func search(args []string) {
	response, duration, err := suggestion.Suggest(strings.Join(args, " "), queryType)
	if err != nil {
		fmt.Println("Error in search: ", err)
		return
	}
	if !Quiet {
		fmt.Printf("%d result(s) (in %0.2fs)\n", len(response.Results), duration.Seconds())
	}
	//use terminal to print colorful results
	for i, r := range response.Results {
		terminal.Stdout.Color("y").
			Print(fmt.Sprintf("\n[%d] ", i+1)).
			Colorf("@{w}" + r.ID).
			Nl().
			Colorf("@{b}" + r.Suggestion + "\n")
		//fmt.Printf("\n[%d] %s\n%s\n", i+1, r.ID, r.Suggestion)
	}
}
