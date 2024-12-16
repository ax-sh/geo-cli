package tui

import (
	"geo/pkg/color"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(color.GRAY)

type (
	errMsg error
)
type model struct {
	textInput textinput.Model
	callback  func(input string) string
	err       error
}

func initialModel(callback func(input string) string) model {
	ti := textinput.New()
	ti.Placeholder = "Phone"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
		callback:  callback,
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
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}
	m.callback(m.textInput.Value())

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	input := m.textInput.View()
	tableString := m.callback(m.textInput.Value())
	helpText := "Type to filter | q/ctrl+c to quit"
	label := "Enter country code for country?"
	return lipgloss.JoinVertical(lipgloss.Left,
		baseStyle.Render(label),
		input,
		color.YellowColorText.Render(helpText),
		tableString,
	)
}

func FilterPhone(callback func(input string) string) {
	p := tea.NewProgram(initialModel(callback))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
