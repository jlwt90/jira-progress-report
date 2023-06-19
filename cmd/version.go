package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// versionCmd represents the version of this tool.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of this command-line tool",
	Run:   runVersion,
}

// version is version of this tool.
// it is overwritten by -ldflags while building.
var version string

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) {
	if version == "" {
		color.BgRed.Println("dev build")
		return
	}
	color.Normal.Println(version)
}
