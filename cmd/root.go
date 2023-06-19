package cmd

import (
	"fmt"
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
		fmt.Println(err)
		os.Exit(exitError)
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
