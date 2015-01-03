package commands

import (
	"github.com/dnotez/dnotez-cli/config"
	"github.com/dnotez/dnotez-cli/utils"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"strings"
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

var Quiet bool
var VerboseLog, ServerAddress string

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
	//keep number of flags short, extra parameters can be configured in the config file
	DNotezCmd.PersistentFlags().StringVar(&ServerAddress, "server", "http://localhost:5050", "Backend server url")
	DNotezCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "Do not print detail information")
	DNotezCmd.PersistentFlags().StringVarP(&VerboseLog, "verbose", "v", "", "Stdout verbose log, can be ERROR, WARN, INFO, DEBUG")
	//DNotezCmd.PersistentFlags().StringVar(&CfgFile, "cfg", "config", "config file (default is config.yml in $HOME/.dnotez/ folder)")

	dzCmdV = DNotezCmd
}

// InitializeConfig initializes a config file with sensible default configuration flags.
func InitializeConfig() {
	//jww.SetStdoutThreshold(jww.LevelDebug)
	//default values
	viper.SetDefault("LogFile", "")
	viper.SetDefault("LogLevel", "")
	viper.SetDefault("server", "http://localhost:5050")

	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.dnotez")
	err := viper.ReadInConfig()
	if err != nil {
		jww.ERROR.Println("Unable to locate config file:", err)
	}

	if len(VerboseLog) > 0 || (viper.IsSet("LogFile") && viper.GetString("LogFile") != "") {
		if viper.IsSet("LogFile") && viper.GetString("LogFile") != "" {
			jww.SetLogFile(viper.GetString("LogFile"))
		} else {
			jww.UseTempLogFile("dnotez")
		}

		if viper.IsSet("LogLevel") && viper.GetString("LogLevel") != "" {
			logLevel := getLogLevelFrom(viper.GetString("LogLevel"), jww.LevelWarn)
			jww.SetLogThreshold(logLevel)
		}
	} else {
		jww.DiscardLogging()
	}

	if len(VerboseLog) > 0 {
		jww.SetStdoutThreshold(getLogLevelFrom(VerboseLog, jww.LevelWarn))
	}

	jww.DEBUG.Println("viper config 'server':", viper.GetString("server"))
	if dzCmdV.PersistentFlags().Lookup("server").Changed {
		jww.DEBUG.Println("Using flag for viper config 'server':", viper.GetString("server"))
		viper.SetDefault("server", ServerAddress)
	}

	config.Server.URL = viper.GetString("server")
	jww.INFO.Println("Using backend server:", config.Server.URL)

	jww.INFO.Println("Using config file:", viper.ConfigFileUsed())
}

func run() {
	dzCmdV.Help()
}

func getLogLevelFrom(levelStr string, defaultLevel jww.Level) jww.Level {
	levelStr = strings.ToUpper(levelStr)
	switch levelStr {
	case "DEBUG":
		return jww.LevelDebug
	case "INFO":
		return jww.LevelInfo
	case "ERROR":
		return jww.LevelError
	case "WARN":
		return jww.LevelWarn
	}

	return defaultLevel
}
