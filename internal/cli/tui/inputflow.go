package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type InputFlow struct {
	title     string
	prompt    string
	textInput textinput.Model
	onSubmit  func(val string) tea.Msg
}

func NewInputFlow(title, prompt string, onSubmit func(val string) tea.Msg) *InputFlow {
	ti := textinput.New()
	ti.Placeholder = prompt
	ti.CharLimit = 256
	ti.Width = 40
	ti.Focus()

	return &InputFlow{
		title:     title,
		prompt:    prompt,
		textInput: ti,
		onSubmit:  onSubmit,
	}
}

func (f *InputFlow) Init() tea.Cmd {
	return textinput.Blink
}

func (f *InputFlow) Update(msg tea.Msg) (tea.Msg, tea.Cmd) {
	var cmd tea.Cmd
	f.textInput, cmd = f.textInput.Update(msg)
	if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
		return f.onSubmit(f.textInput.Value()), nil
	}
	return nil, cmd
}

func (f *InputFlow) View() string {
	s := fmt.Sprintf("%s\n\n%s\n\n", f.title, f.textInput.View())
	s += "Enter 提交，Esc 取消"
	return s
}
