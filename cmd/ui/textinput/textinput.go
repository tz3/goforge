package textinput

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	program "github.com/tz3/goforge/internal/project"
)

var (
	titleStyle        = lipgloss.NewStyle().Background(lipgloss.Color("#00BFFF")).Foreground(lipgloss.Color("#FFFFFF")).Bold(true).Padding(0, 1, 0)
	errorMessageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true).Padding(0, 0, 0)
)

type (
	errMsg error
)

type Output struct {
	Output string
}

func (o *Output) update(val string) {
	o.Output = val
}

type model struct {
	textInput textinput.Model
	err       error
	output    *Output
	header    string
	exit      *bool
}

func InitialTextInputModel(output *Output, header string, program *program.ProjectConfig) model {
	if !isValidInput(header) || header == "" {
		return CreateErrorModel(errors.New("input contains non-alphanumeric characters"))
	}

	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
		output:    output,
		header:    titleStyle.Render(header),
		exit:      &program.Exit,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if len(m.textInput.Value()) > 1 {
				m.output.update(m.textInput.Value())
				return m, tea.Quit
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			*m.exit = true
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		*m.exit = true
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf("%s\n\n%s\n\n",
		m.header,
		m.textInput.View(),
	)
}

// CreateErrorModel generates a model for handling error input.
// It initializes a new text input, sets focus on it, and limits the character input to 156.
// The width of the text input is set to 20.
// It returns a model with the initialized text input, the error message, and an exit flag set to true.
func CreateErrorModel(err error) model {
	textInput := textinput.New()
	textInput.Focus()
	textInput.CharLimit = 156
	textInput.Width = 20
	exitFlag := true

	return model{
		textInput: textInput,
		err:       errors.New(errorMessageStyle.Render(err.Error())),
		output:    nil,
		header:    "",
		exit:      &exitFlag,
	}
}

// Error returns the error message of the model.
// It is a method of the model type.
func (m model) Error() string {
	return m.err.Error()
}

// isValidInput checks if a string only contains alphanumeric characters.
func isValidInput(input string) bool {
	// accept only normal english chars + numbers
	match, _ := regexp.MatchString("^[a-zA-Z0-9]*$", input)
	return match
}
