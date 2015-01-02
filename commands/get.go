package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/dnotez/dnotez-cli/dsl/article"
	"github.com/spf13/cobra"
	//"github.com/wsxiaoys/terminal"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a note by a key",
	Long: `Fetch a note by id or label.
Output fields can be filtered by --field flag.

Examples:
dz get 577ae747-951b-41d5-a7a6-1f0298a5766f
dz get -f text 577ae747-951b-41d5-a7a6-1f0298a5766f
dz get -l update docker`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 && !label {
			cmd.Help()
		} else {
			get(args)
		}
	},
}

var field string
var label bool

func init() {
	getCmd.Flags().BoolVarP(&label, "label", "l", false, "Get note by the nabel instead of id")
	getCmd.Flags().StringVarP(&field, "field", "f", "all", "Output field (can be id, html, text, title or all)")
}

func get(args []string) {
	var key string
	if key = "id"; label {
		key = "label"
	}

	articles, duration, err := article.Get(strings.Join(args, " "), key, 1)
	if err != nil {
		fmt.Println("Error in get: ", err)
		return
	}
	if !Quiet {
		fmt.Printf("Fetched in %0.2fs\n\n", duration.Seconds())
	}

	for _, article := range *articles {
		switch field {
		case "text":
			fmt.Printf("%s\n", article.Text)
		case "title":
			fmt.Printf("%s\n", article.Title)
		case "id":
			fmt.Printf("%s\n", article.Id)
		case "label":
			fmt.Printf("%s\n", article.Label)
		default:
			fmt.Printf("%s\n", article.Id)
			fmt.Printf("%s\n", article.Title)
			fmt.Printf("%s\n", article.Text)
			fmt.Printf("%s\n", article.Label)
			fmt.Printf("%s\n", time.Unix(article.SaveDate/1000, 0).Local())
			fmt.Printf("\n")
		}
	}

}
