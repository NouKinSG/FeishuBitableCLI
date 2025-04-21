package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	subHeaderStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("15")) // 白色，更清晰
	highlightStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212"))
	cursorSymbol    = "❯"
	normalSymbol    = "  "
	separator       = lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Render(strings.Repeat("=", 52))
	footerHintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39")) // 蓝绿色，提高可见性
)
