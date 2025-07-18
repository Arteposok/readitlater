package cmd

import (
	"fmt"
	"readitlater/data"
	"regexp"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "outputs all the items that match the pattern",
	Long:  "searches for all entries and outputs the ones that match the provided pattern",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var pattern regexp.Regexp = *regexp.MustCompile(args[0])
		notes, err := data.GetAllNotes()
		if err != nil {
			fmt.Printf("Error fetching notes: %v\n", err)
			return
		}

		if len(notes) == 0 {
			fmt.Println("No notes found.")
			return
		}
		fmt.Println("Notes found:")
		for name, note := range notes {
			if pattern.MatchString(name) {
				fmt.Printf("%s %s\n", name, note)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
