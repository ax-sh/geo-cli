package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/go-gota/gota/dataframe"
)

// PrintDataframe printable pretty tui
func PrintDataframe(df dataframe.DataFrame) *table.Table {
	prettyTable := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers(df.Names()...).
		Rows(df.Records()[1:]...)
	return prettyTable
}
