package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
	"github.com/jlwt90/reportify/tracker"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// initCmd initialises the project profile of any supported tracker engine.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise project profile",
	Run:   runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) {
	// Load config
	if err := viper.ReadInConfig(); err != nil {
		viper.Set("ConfigFilePath", cfgPath)

		if err = os.MkdirAll(cfgPath, os.ModePerm); err != nil {
			fmt.Println(err)
			os.Exit(exitError)
		}

		if err = viper.WriteConfigAs(cfgPath + "/config.yaml"); err != nil {
			fmt.Println(err)
			os.Exit(exitError)
		}
	}

	// Select project tracker
	sys := ""
	prompt := &survey.Select{
		Message: "Choose a Project Tracking System:",
		Options: tracker.SupportedTrackers,
	}
	if err := survey.AskOne(prompt, &sys); err != nil {
		fmt.Println(err)
		os.Exit(exitError)
	}
	color.Bold.Printf("Generating project profile for %s. \n", sys)

	// Set up tracker if necessary
	t, ok := tracker.NewTracker(sys)
	if !ok {
		fmt.Println("Tracker type is unsupported")
		os.Exit(exitError)
	}
	if err := t.SetUpTracker(); err != nil {
		fmt.Println(err)
		os.Exit(exitError)
	}

	// persist configuration to config directory
	if err := viper.WriteConfig(); err != nil {
		fmt.Println(err)
		os.Exit(exitError)
	}
}
