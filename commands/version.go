package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"bitbucket.org/kardianos/osext"
	"github.com/spf13/cobra"
)

var timeLayout string // the layout for time.Time

var (
	commitHash string
	buildDate  string
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Prints build date and commit hash details`,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig()
		if buildDate == "" {
			setBuildDate() // set the build date from executable's mdate
		} else {
			formatBuildDate() // format the compile time
		}
		if commitHash == "" {
			fmt.Printf("dz v0.1 build date: %s\n", buildDate)
		} else {
			fmt.Printf("dz v0.1 commit hash:%s build date: %s\n", strings.ToUpper(commitHash), buildDate)
		}
	},
}

// setBuildDate checks the ModTime of the Hugo executable and returns it as a
// formatted string.  This assumes that the executable name is Hugo, if it does
// not exist, an empty string will be returned.  This is only called if the
// buildDate wasn't set during compile time.
//
// osext is used for cross-platform.
func setBuildDate() {
	fname, _ := osext.Executable()
	dir, err := filepath.Abs(filepath.Dir(fname))
	if err != nil {
		fmt.Println(err)
		return
	}
	fi, err := os.Lstat(filepath.Join(dir, "dz"))
	if err != nil {
		fmt.Println(err)
		return
	}
	t := fi.ModTime()
	buildDate = t.Format(time.RFC3339)
}

// formatBuildDate formats the buildDate according to the value in
// .Params.DateFormat, if it's set.
func formatBuildDate() {
	t, _ := time.Parse("2006-01-02T15:04:05-0700", buildDate)
	buildDate = t.Format(time.RFC3339)
}
