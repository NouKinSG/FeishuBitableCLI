package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View 分发入口：区分配置编辑、多维表格名称输入、普通菜单
func (m *model) View() string {
	// 配置管理的编辑流程
	if m.current == menuConfig && m.editStage != stageNone {
		return m.viewConfigEditFlow()
	}

	if m.current == menuBitable && m.editStage == stageEnterBitableName {
		return m.viewEnterBitableName()
	}

	// 普通菜单视图
	return m.viewMenu()
}

// viewMenu 渲染常规菜单（含主菜单 & 各子菜单）
func (m *model) viewMenu() string {
	var b strings.Builder
	b.WriteString(m.renderSeparator())
	b.WriteString(m.renderHeader())
	b.WriteString(m.renderSubHeaders())
	b.WriteString(m.renderSeparator())
	b.WriteString(m.renderHint())
	b.WriteString("\n\n")

	b.WriteString(m.renderOptions())

	// 子菜单下的额外内容
	switch m.current {
	case menuConfig:
		if m.cursor == 0 {
			b.WriteString(m.renderConfigData())
		}
	case menuBitable:
		// TODO: 在此处可添加多维表格管理的列表展示
	}

	b.WriteString(m.renderStatus())
	return b.String()
}

// viewConfigEditFlow 渲染配置管理的编辑流程（选键 & 输入）
func (m *model) viewConfigEditFlow() string {
	var b strings.Builder
	b.WriteString(m.renderSeparator())

	switch m.editStage {
	case stagePickKey:
		b.WriteString(headerStyle.Render("⚙️ 配置管理 > 选择键名") + "\n\n")
		for i, key := range m.editKeys {
			cursor := normalSymbol
			if i == m.cursor {
				cursor = cursorSymbol
			}
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, key))
		}

	case stageEnterValue:
		b.WriteString(headerStyle.Render(fmt.Sprintf("⚙️ 配置管理 > 修改 %s", m.editKey)) + "\n\n")
		b.WriteString(m.textInput.View() + "\n\n")
		b.WriteString(footerHintStyle.Render("Enter 提交，Esc 取消") + "\n")
	}

	b.WriteString(m.renderStatus())
	return b.String()
}

// 统一渲染函数
func (m *model) renderSeparator() string {
	return separator + "\n"
}

func (m *model) renderHeader() string {
	return headerStyle.Render("🚀 飞书多维表格 CLI 工具 v0.1.0") + "\n"
}

func (m *model) renderSubHeaders() string {
	return strings.Join([]string{
		subHeaderStyle.Render(fmt.Sprintf("📂 当前表格：%s", displayOr(m.selectedBitable, "未选择"))),
		subHeaderStyle.Render(fmt.Sprintf("📑 当前数据表：%s", displayOr(m.selectedTable, "未选择"))),
		subHeaderStyle.Render(fmt.Sprintf("📍 当前路径：%s", m.currentPath())),
	}, "\n") + "\n"
}

func (m *model) renderHint() string {
	return footerHintStyle.Render("↑↓ 选择，Enter 确认，Esc/q 退出")
}

// renderOptions 渲染光标移动样式
func (m *model) renderOptions() string {
	var b strings.Builder
	for i, opt := range m.currentOptions() {
		prefix, style := normalSymbol, lipgloss.NewStyle()
		if i == m.cursor {
			prefix = cursorSymbol
			style = highlightStyle
		}
		b.WriteString(fmt.Sprintf("%s %s\n", prefix, style.Render(opt)))
	}
	return b.String()
}

// renderConfigData 渲染配置键值列表
func (m *model) renderConfigData() string {
	var b strings.Builder
	b.WriteString("\n当前配置：\n")
	for k, v := range m.configData {
		b.WriteString(fmt.Sprintf(" • %s: %v\n", k, v))
	}
	return b.String()
}

// renderStatus 底部渲染提示
func (m *model) renderStatus() string {
	if m.statusMsg != "" {
		return "\n" + m.statusMsg + "\n"
	}
	return ""
}

// currentPath 返回当前菜单路径
func (m *model) currentPath() string {
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
		return ""
	}
}

// currentOptions 返回当前菜单可选项
func (m *model) currentOptions() []string {
	switch m.current {
	case menuMain:
		return []string{"🗂 多维表格管理", "🗄 数据表管理", "🧱 字段管理", "📝 记录管理", "⚙️ 配置管理", "🚪 退出"}
	case menuBitable:
		return []string{"查看已有多维表格（TODO）", "创建新的多维表格", "删除多维表格（TODO）", "⬅️ 返回主菜单"}
	case menuTable:
		return []string{"查看数据表列表（TODO）", "创建数据表（TODO）", "删除数据表（TODO）", "⬅️ 返回主菜单"}
	case menuField:
		return []string{"添加字段（TODO）", "查看字段（TODO）", "⬅️ 返回主菜单"}
	case menuRecord:
		return []string{"插入 Mock 数据（TODO）", "查看所有记录（TODO）", "⬅️ 返回主菜单"}
	case menuConfig:
		return []string{"查看当前配置", "修改配置项", "刷新配置项", "重置为默认配置", "⬅️ 返回主菜单"}
	default:
		return nil
	}
}

// viewEnterBitableName 渲染“输入多维表格名称”流程
func (m *model) viewEnterBitableName() string {
	var b strings.Builder
	b.WriteString(separator + "\n")
	b.WriteString(headerStyle.Render("🗂 新建多维表格 > 输入名称") + "\n\n")
	b.WriteString(m.textInput.View() + "\n\n")
	b.WriteString(footerHintStyle.Render("Enter 提交，Esc 取消") + "\n")
	if m.statusMsg != "" {
		b.WriteString("\n" + m.statusMsg + "\n")
	}
	return b.String()
}

func (m *model) optionCount() int {
	return len(m.currentOptions())
}
