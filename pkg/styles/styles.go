package styles

import (
	"geo/pkg/color"
	"github.com/charmbracelet/lipgloss"
)

var DocStyle = lipgloss.NewStyle().Margin(1, 2)

var BaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(color.GRAY)
