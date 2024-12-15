package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/go-gota/gota/dataframe"
)

var (
	HeaderStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("99"))
	EvenRowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("242"))
	OddRowStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("242"))
)

// PrintDataframe printable pretty tui
func PrintDataframe(df dataframe.DataFrame) *table.Table {
	prettyTable := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				return EvenRowStyle
			default:
				return OddRowStyle
			}
		}).
		Headers(df.Names()...).
		Rows(df.Records()[1:]...)
	return prettyTable
}
