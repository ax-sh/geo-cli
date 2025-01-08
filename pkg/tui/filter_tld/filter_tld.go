package filter_tld

import (
	"fmt"
	"geo/pkg/country"
	"geo/pkg/styles"
	"geo/pkg/tui"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	textInput textinput.Model
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Top level domain"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return model{
		textInput: ti,
	}
}
func (m model) Init() tea.Cmd {
	return textinput.Blink
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	input := m.textInput.Value()

	styledInput := styles.DocStyle.Inline(true).Foreground(lipgloss.Color("99")).Render(input)
	fil := country.FilterCountryByTLDDataFrame(input)
	sel := country.NormalizeCountryDataFrame(fil)
	sel = sel.Drop([]string{"ISO", "ISO3", "ISO-Numeric"}).
		Drop("neighbours").
		Drop("Languages")

	result := tui.PrintDataframe(sel)
	output := result.String()

	return lipgloss.JoinVertical(lipgloss.Left, styledInput, output)
}

func FilterTldTui() (string, error) {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("error running program: %w", err)
	}
	if finalModel, ok := m.(model); ok {
		value := finalModel.textInput.Value()
		fmt.Println("Result for tld", value)
		return value, nil
	}
	return "", fmt.Errorf("program quit without selecting value")
}
