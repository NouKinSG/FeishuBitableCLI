package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func StartTUI() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("TUI 启动失败:", err)
		os.Exit(1)
	}
}
