package cmd

import (
	"fmt"
	"os"
	"readitlater/data"
	"regexp"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/spf13/cobra"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput textinput.Model
	found     map[string]string
}

type found map[string]string

func (m model) matchNotes() found {
	var pattern regexp.Regexp = *regexp.MustCompile(m.textInput.Value())
	notes, err := data.GetAllNotes()
	if err != nil {
		fmt.Printf("Error fetching notes: %v\n", err)
		return map[string]string{}
	}

	if len(notes) == 0 {
		return map[string]string{}
	}
	var results map[string]string = map[string]string{}
	for name, note := range notes {
		if pattern.MatchString(name) {
			results[name] = note
		}
	}
	return results
}

func initialModel() model {
	textInput := textinput.New()
	textInput.Focus()
	textInput.Width = 50
	textInput.Placeholder = "^.*$"

	return model{
		textInput: textInput,
		found:     map[string]string{},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			return m, tea.Quit

		default:
			return m, tea.Batch(func() tea.Msg {
				return m.matchNotes()
			}, cmd)
		}
	case found:
		m.found = msg
	}
	return m, cmd
}

func (m model) View() string {
	s := "What are you looking for?\n\n"
	s += m.textInput.View() + "\n\n"
	for name, note := range m.found {
		s += name + " --> " + note + "\n"
	}
	s += "Ctrl + Q to exit"
	return s
}

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"sr"},
	Short:   "enters interactive search mode",
	Long:    "allows you to fuzzy-find through your enormous list of notes:)",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Huston, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
