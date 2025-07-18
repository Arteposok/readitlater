package cmd

import (
	"fmt"
	"readitlater/data"

	"github.com/spf13/cobra"
)

var getAllCmd = &cobra.Command{
	Use:   "get_all",
	Short: "outputs all the items",
	Long:  "prints out all the items",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		notes, err := data.GetAllNotes()
		if err != nil {
			fmt.Printf("Error fetching notes: %v\n", err)
			return
		}

		if len(notes) == 0 {
			fmt.Println("No notes found.")
			return
		}

		fmt.Println("Your notes:")
		for name, note := range notes {
			fmt.Printf("%s %s\n", name, note)
		}
	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)
}
