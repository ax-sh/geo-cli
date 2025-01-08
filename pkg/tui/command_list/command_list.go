package command_list

import (
	"fmt"
	"geo/pkg/country"
	"geo/pkg/tui"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.list.StartSpinner())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			m.quitting = true
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			// Store the selected item's title when Enter is pressed
			if i, ok := m.list.SelectedItem().(item); ok {
				m.choice = i.title
			}
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func FooMain() {
	items := []list.Item{
		item{title: "phone", desc: "Find country by country code"},
		item{title: "tld", desc: "Search by Top layer domain"},
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.FilterInput.Focus()
	l.SetFilteringEnabled(true)
	l.SetShowFilter(true)

	m := model{
		list: l,
	}
	//m.list.Title = "Tools"
	//m.list.SetShowTitle(true)

	p := tea.NewProgram(m, tea.WithAltScreen())

	finalModel, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	// Get the selected item from the final model

	if finalModel, ok := finalModel.(model); ok {

		switch finalModel.choice {
		case "tld":
			println("todo make this similar to phone")
		case "phone":
			callback := func(countryCode string) string {
				if len(countryCode) == 0 {
					return "Type to filter"
				}
				fil := country.FilterCountryByCountryCodeDataFrame(countryCode)
				sel := country.DropUselessCountryColumn(fil).
					Drop("Area(in sq km)")
				sel = country.MoveImportantColumnsToStart(sel)
				sel = country.MoveColumnsToStart(sel, "Phone")
				result := tui.PrintDataframe(sel)
				return result.String()
			}
			tui.FilterPhone(callback)
		default:
			println("[You picked choice]", docStyle.Render(finalModel.choice))
		}

	}
}
