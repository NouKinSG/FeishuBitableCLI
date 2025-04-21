package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

// View 渲染函数
func (m model) View() string {
	var b strings.Builder

	b.WriteString(separator + "\n")
	b.WriteString(headerStyle.Render("🚀 飞书多维表格 CLI 工具 v0.1.0") + "\n")
	b.WriteString(subHeaderStyle.Render(fmt.Sprintf("📂 当前表格：%s", displayOr(m.selectedBitable, "未选择"))) + "\n")
	b.WriteString(subHeaderStyle.Render(fmt.Sprintf("📑 当前数据表：%s", displayOr(m.selectedTable, "未选择"))) + "\n")
	b.WriteString(subHeaderStyle.Render(fmt.Sprintf("📍 当前路径：%s", m.currentPath())) + "\n")
	b.WriteString(separator + "\n")
	b.WriteString(footerHintStyle.Render("↑↓ 选择，Enter 确认，q 退出") + "\n")

	opts := m.currentOptions()
	for i, opt := range opts {
		prefix := normalSymbol
		style := lipgloss.NewStyle()
		if i == m.cursor {
			prefix = cursorSymbol
			style = highlightStyle
		}
		b.WriteString(fmt.Sprintf("%s %s\n", prefix, style.Render(opt)))
	}
	return b.String()
}

// currentPath 根据 m.current 返回路径字符串
func (m model) currentPath() string {
	switch m.current {
	case menuMain:
		return "主菜单"
	case menuBitable:
		return "主菜单 > 多维表格管理"
	case menuTable:
		return "主菜单 > 数据表管理"
	case menuField:
		return "主菜单 > 字段管理"
	case menuRecord:
		return "主菜单 > 记录管理"
	case menuConfig:
		return "主菜单 > 配置管理"
	default:
		return "未知路径"
	}
}

// currentOptions 根据 m.current 返回菜单选项列表
func (m model) currentOptions() []string {
	switch m.current {
	case menuMain:
		return []string{"🗂 多维表格管理", "🗄 数据表管理", "🧱 字段管理", "📝 记录管理", "⚙️ 配置管理", "🚪 退出"}
	case menuBitable:
		return []string{"查看已有多维表格（TODO）", "创建新的多维表格（TODO）", "删除多维表格（TODO）", "⬅️ 返回主菜单"}
	case menuTable:
		return []string{"查看数据表列表（TODO）", "创建数据表（TODO）", "删除数据表（TODO）", "⬅️ 返回主菜单"}
	case menuField:
		return []string{"添加字段（TODO）", "查看字段（TODO）", "⬅️ 返回主菜单"}
	case menuRecord:
		return []string{"插入 Mock 数据（TODO）", "查看所有记录（TODO）", "⬅️ 返回主菜单"}
	case menuConfig:
		return []string{"查看配置（TODO）", "修改配置（TODO）", "⬅️ 返回主菜单"}
	default:
		return []string{}
	}
}

func (m model) optionCount() int {
	return len(m.currentOptions())
}
