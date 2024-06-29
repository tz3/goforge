package multiselect

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	program "github.com/tz3/goforge/internal/project"
	"github.com/tz3/goforge/internal/steps"
)

var (
	focusedStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Bold(true)
	titleStyle            = lipgloss.NewStyle().Background(lipgloss.Color("#00BFFF")).Foreground(lipgloss.Color("#FFFFFF")).Bold(true).Padding(0, 1, 0)
	selectedItemStyle     = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#FFD700")).Bold(true)
	selectedItemDescStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#FFD700"))
	descriptionStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#87CEEB"))
)

// A Selection represents a choice made in a multiSelect step
type Selection struct {
	Choices map[string]bool
}

// Update changes the value of a Selection's Choice
func (s *Selection) Update(optionName string, value bool) {
	s.Choices[optionName] = value
}

// A multiSelect.model contains the data for the multiSelect step.
//
// It has the required methods that make it a bubbletea.Model
type model struct {
	cursor   int
	options  []steps.Option
	selected map[int]struct{}
	choices  *Selection
	header   string
	exit     *bool
}

func (m model) Init() tea.Cmd {
	return nil
}

// InitialModelMulti initializes a multiSelect step with
// the given data
func InitialModelMultiSelect(options []steps.Option, selection *Selection, header string, program *program.ProjectConfig) model {
	return model{
		options:  options,
		selected: make(map[int]struct{}),
		choices:  selection,
		header:   titleStyle.Render(header),
		exit:     &program.Exit,
	}
}

// Update is called when "things happen", it checks for
// important keystrokes to signal when to quit, change selection,
// and confirm the selection.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			*m.exit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "y":
			for selectedKey := range m.selected {
				m.choices.Update(m.options[selectedKey].Flag, true)
				m.cursor = selectedKey
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

// View is called to draw the multiSelect step
func (m model) View() string {
	s := m.header + "\n\n"

	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = focusedStyle.Render(">")
			option.Title = selectedItemStyle.Render(option.Title)
			option.Desc = selectedItemDescStyle.Render(option.Desc)
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = focusedStyle.Render("*")
		}

		title := focusedStyle.Render(option.Title)
		description := descriptionStyle.Render(option.Desc)

		s += fmt.Sprintf("%s [%s] %s\n%s\n\n", cursor, checked, title, description)
	}

	s += fmt.Sprintf("Press %s to confirm choice.\n", focusedStyle.Render("y"))
	return s
}
