package main

import (
	"fmt"
	flag "github.com/dnotez/cli/mflag"
	"os"
)

var (
	flVersion = flag.Bool([]string{"version", "-version"}, false, "Print version information")
	flVerbose = flag.Bool([]string{"v", "-verbose"}, false, "Show more details")
)

func init() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage: pl [OPTIONS] command\nYour knowledge base, built by you & your team.\n\nOptions:\n")
		flag.PrintDefaults()

		help := "\nCommands:\n"
		for _, command := range [][]string{
			{"save", "Save the current command/resource"},
			{"search", "Search for a saved command/resource"},
			{"share", "Share a saved resource with colleages."},
			{"get", "Get an resource."},
			{"rm", "Remove a resource."},
		} {
			help += fmt.Sprintf("    %-10.10s%s\n", command[0], command[1])
		}
		help += "\nRun 'pl COMMAND --help' for more information on a command."
		fmt.Fprintf(os.Stderr, "%s\n", help)
	}
}
