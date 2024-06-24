package spinner

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type spinnerModel struct {
	done bool
	err  error
}

func InitialModel() spinnerModel {
	return spinnerModel{}
}

func (m spinnerModel) Init() tea.Cmd {
	return tea.Batch(tickCmd(), tea.EnterAltScreen)
}

func (m spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tickMsg:
		if m.done {
			return m, tea.Quit
		}
		return m, tickCmd()
	case CompleteMsg:
		m.done = true
		m.err = msg.Err
		return m, tea.Quit
	}
	return m, nil
}

func (m spinnerModel) View() string {
	if m.done {
		if m.err != nil {
			return fmt.Sprintf("Project setup failed: %v\n", m.err)
		}
		return "Project setup complete!\n"
	}
	return fmt.Sprintf("Setting up project ... %s\n", spinner())
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

var spinnerState int

func spinner() string {
	spinners := []string{"|", "/", "-", "\\"}
	spinnerState = (spinnerState + 1) % len(spinners)
	return spinners[spinnerState]
}

// CompleteMsg signals that the project setup is complete.
type CompleteMsg struct {
	Err error
}
