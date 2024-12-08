/**
 * Package cmd - The "cmd" package contains the CLI commands for the tool.
 *
 * The "clean" file in particular which is part of the "cmd" package contains
 * the logic to clean up downloaded template(s) on disk.
 */
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Weburz/terox/internal/template"
)

// Command help
var cleanShortUsage = "Clean/delete all downloaded templates"
var cleanLongUsage = `
Cleanup the downloaded templates from disk.
`

// Command to handle the "clean" command of the CLI tool
var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"gc", "cleanup"},
	Short:   cleanShortUsage,
	Long:    cleanLongUsage,
	Example: "terox clean\nterox gc\nterox cleanup",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		template.Clean()
	},
}

// Register the command to the CLI tool
func init() {
	rootCmd.AddCommand(cleanCmd)
}
