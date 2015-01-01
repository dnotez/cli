package commands

import (
	"fmt"    "utils"

)

//DNotzCmd is root command. Every other command attached to PlCmd is a child command to it.
var DNotzCmd = &cobra.Command{
	Use:   "dz",
	Short: "dz is command line interface for dNotz.com",
	Long: `dz is command line interface for dnotz.com
	dnotz.com is an online service for collecting, re-using and sharing code snippets, command lines and quick notes among software devlopers.
Complete documentation is available at http://gohugo.io`,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig()
		run()
	},
}

var dzCmdV *cobra.Command

var Host, CfgFile

//Execute adds all child commands to the root command HugoCmd and sets flags appropriately.
func Execute() {
	AddCommands()
	utils.StopOnErr(DNotzCmd.Execute())
}

//AddCommands adds child commands to the root command HugoCmd.
func AddCommands() {
	//DNotzCmd.AddCommand(serverCmd)
	//DNotzCmd.AddCommand(version)
	//DNotzCmd.AddCommand(check)
	//DNotzCmd.AddCommand(benchmark)
	//DNotzCmd.AddCommand(convertCmd)
	//DNotzCmd.AddCommand(newCmd)
	//DNotzCmd.AddCommand(listCmd)
}

func init() {
	DNotzCmd.PersistentFlags().StringVarP(&Host, "host", "h", "", "Backend server host full url or IP address, e.g. http://www.dnotz.com")
	DNotzCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is path/config.yaml|json|toml)")
}

// InitializeConfig initializes a config file with sensible default configuration flags.
func InitializeConfig() {
	viper.SetConfigFile(CfgFile)
	viper.AddConfigPath(Source)
	err := viper.ReadInConfig()
	if err != nil {
		jww.ERROR.Println("Unable to locate Config file. Perhaps you need to create a new site. Run `hugo help new` for details")
	}
}

func run() {
	fmt.Println("dz root command")
}
