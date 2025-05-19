package tui

import (
	"fmt"
	"github.com/NouKinSG/FeishuBitableCLI/internal/bitable"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

// startCreateBitableMsg 在后台调用完成后发回的信息
type startCreateBitableMsg struct {
	Info *bitable.AppInfo
	Err  error
}

// createBitableCmd 发起异步创建多维表格的命令
func createBitableCmd(name string, folderToken string) tea.Cmd {
	return func() tea.Msg {
		info, err := bitable.CreateApp(name, folderToken)
		return startCreateBitableMsg{Info: info, Err: err}
	}
}

// —— 视图片段 ——

// viewCreateFlow 渲染“输入表格名称”流程
func (m *model) viewCreateFlow() string {
	var b strings.Builder
	b.WriteString(separator + "\n")
	b.WriteString(headerStyle.Render("🗂 多维表格管理 > 新建表格") + "\n\n")
	b.WriteString(m.textInput.View() + "\n\n")
	b.WriteString(footerHintStyle.Render("Enter 创建，Esc 取消") + "\n")
	b.WriteString(m.renderStatus())
	return b.String()
}

// —— 结果处理 ——

// handleCreateBitableMsg 将 startCreateBitableMsg 写回 model
func (m *model) handleCreateBitableMsg(msg startCreateBitableMsg) {
	if msg.Err != nil {
		m.statusMsg = "❌ 创建失败: " + msg.Err.Error()
	} else {
		// Info.AppToken 和 Name 都是 string 类型，直接使用
		m.selectedBitable = msg.Info.AppToken
		m.statusMsg = fmt.Sprintf("✅ 创建成功：%s (%s)", msg.Info.Name, msg.Info.AppToken)
	}
	// 回到“多维表格管理”菜单
	m.editStage = stageNone
	m.current = menuBitable
	m.cursor = 0
}
