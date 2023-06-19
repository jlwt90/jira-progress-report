package cmd

import (
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
		if err = os.MkdirAll(cfgPath, os.ModePerm); err != nil {
			terminateCmd(err, "")
		}
		if err = viper.WriteConfigAs(cfgPath + "/config.yaml"); err != nil {
			terminateCmd(err, "")
		}
	}

	// Select project tracker
	sys := ""
	prompt := &survey.Select{
		Message: "Choose a Project Tracking System:",
		Options: tracker.SupportedTrackers,
	}
	if err := survey.AskOne(prompt, &sys); err != nil {
		terminateCmd(err, "")
	}
	color.Bold.Printf("Generating project profile for %s. \n", sys)

	// Set up tracker if necessary
	t, ok := tracker.NewTracker(sys)
	if !ok {
		terminateCmd(nil, "Tracker type is unsupported")
	}
	if err := t.SetUpTracker(); err != nil {
		terminateCmd(err, "")
	}

	// persist configuration to config directory
	if err := viper.WriteConfig(); err != nil {
		terminateCmd(err, "")
	}
}
