package commands

import (
	"github.com/dnotez/dnotez-cli/utils"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

//DNotezCmd is root command. Every other command attached to PlCmd is a child command to it.
var DNotezCmd = &cobra.Command{
	Use:   "dz",
	Short: "dz is command line interface for dNotez",
	Long: `dz is command line interface for dNotez

dNotez is an online service for collecting, re-using and sharing
code snippets, command lines and notes between software devlopers.

Complete documentation is available at dNotez website`,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig()
		run()
	},
}

var dzCmdV *cobra.Command

var VerboseLog, Quiet bool
var Host, CfgFile, LogFile string

//Execute adds all child commands to the root command DNotezCmd and sets flags appropriately.
func Execute() {
	AddCommands()
	utils.StopOnErr(DNotezCmd.Execute())
}

//AddCommands adds child commands to the root command DNotezCmd.
func AddCommands() {
	DNotezCmd.AddCommand(searchCmd)
	DNotezCmd.AddCommand(saveCmd)
	DNotezCmd.AddCommand(getCmd)
	DNotezCmd.AddCommand(removeCmd)
	DNotezCmd.AddCommand(version)
}

func init() {
	DNotezCmd.PersistentFlags().StringVarP(&Host, "url", "u", "", "Backend server host name or IP.")
	DNotezCmd.PersistentFlags().StringVar(&CfgFile, "config", ".dnotez.yaml", "config file (default is dnotez.yaml|json|toml)")
	DNotezCmd.PersistentFlags().StringVar(&LogFile, "log-file", "", "Log File path (if set, logging enabled automatically)")
	DNotezCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "Do not print detail information")
	DNotezCmd.PersistentFlags().BoolVar(&VerboseLog, "verbose-log", false, "Set log level to INFO")
	dzCmdV = DNotezCmd
}

// InitializeConfig initializes a config file with sensible default configuration flags.
func InitializeConfig() {
	viper.SetConfigFile(CfgFile)
	viper.AddConfigPath("$HOME/.dnotez")
	err := viper.ReadInConfig()
	if err != nil {
		jww.ERROR.Println("Unable to locate Config file.")
	}

	if dzCmdV.PersistentFlags().Lookup("log-file").Changed {
		viper.Set("LogFile", LogFile)
	}

	if VerboseLog || (viper.IsSet("LogFile") && viper.GetString("LogFile") != "") {
		if viper.IsSet("LogFile") && viper.GetString("LogFile") != "" {
			jww.SetLogFile(viper.GetString("LogFile"))
		} else {
			jww.UseTempLogFile("dnotez")
		}
	} else {
		jww.DiscardLogging()
	}

	if viper.GetBool("verbose") {
		jww.SetStdoutThreshold(jww.LevelInfo)
	}

	if VerboseLog {
		jww.SetLogThreshold(jww.LevelInfo)
	}

	jww.INFO.Println("Using config file:", viper.ConfigFileUsed())
}

func run() {
	dzCmdV.Help()
}
