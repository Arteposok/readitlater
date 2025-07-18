package cmd

import (
	"readitlater/data"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds an item",
	Long:  "add a readitlater item, so you don't have to think 'bout it now",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var name string = args[0]
		var content string = args[1]
		data.AddNote(content, name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
