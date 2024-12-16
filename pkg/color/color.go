package color

import "github.com/charmbracelet/lipgloss"

const (
	YELLOW = lipgloss.Color("#F1E43C")
	GRAY   = lipgloss.Color("240")
	PURPLE = lipgloss.Color("#7CF143")
)

var YellowColorText = lipgloss.NewStyle().Inline(true).Bold(true).Foreground(YELLOW)
