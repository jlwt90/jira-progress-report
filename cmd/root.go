package cmd

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const (
	appName     = "reportify"
	exitOK  int = iota
	exitError
)

var (
	rootCmd = &cobra.Command{
		Use:   "reportify",
		Short: "Simplify your project progress reporting",
	}
	cfgPath, cfgFile string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		terminateCmd(err, "failed to execute command")
	}
}

func init() {
	cobra.OnInitialize(initCfg)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "cfg", "", "config file (default is $HOME/.reportify/config.yaml)")
}

func initCfg() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	cfgPath = home + "/." + appName
	viper.AddConfigPath(cfgPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func terminateCmd(err error, msg string) {
	var msgFmt string
	var vars []interface{}
	if err != nil && msg != "" {
		msgFmt = "%s: %v \n"
		vars = append(vars, msg, err)
	} else if err != nil {
		msgFmt = "terminated with error: %v \n"
		vars = append(vars, err)
	} else {
		msgFmt = "internal error: %s \n"
		vars = append(vars, msg)
	}
	color.BgRed.Printf(msgFmt, vars)
	os.Exit(exitError)
}
