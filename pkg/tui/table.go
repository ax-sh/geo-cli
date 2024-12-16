package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/go-gota/gota/dataframe"
)

var (
	HeaderStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7CF143"))
	EvenRowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FC318A"))
	OddRowStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#B13CC9"))
)

// PrintDataframe printable pretty tui
func PrintDataframe(df dataframe.DataFrame) *table.Table {
	prettyTable := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style
			switch {
			case row == table.HeaderRow:
				return HeaderStyle
			case row%2 == 0:
				style = EvenRowStyle
			default:
				style = OddRowStyle
			}
			return style
		}).
		Headers(df.Names()...).
		Rows(df.Records()[1:]...)
	return prettyTable
}
