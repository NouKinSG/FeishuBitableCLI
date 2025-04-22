package tui

import (
	"Base/internal/config"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

// Init 初始化入口
func (m *model) Init() tea.Cmd {
	return nil
}

// Update 分发函数：自定义消息 > 编辑阶段 > 普通菜单
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if cmd := m.handleCustomMsg(msg); cmd != nil {
		return m, cmd
	}
	if m.editStage != stageNone {
		return m.updateEditing(msg)
	}
	return m.handleMenuKey(msg)
}

// handleCustomMsg 处理特殊消息（启动编辑或刷新）
func (m *model) handleCustomMsg(msg tea.Msg) tea.Cmd {

	switch msg := msg.(type) {
	// —— 新增：接收“创建多维表格”完成消息 ——
	case startCreateBitableMsg:
		m.handleCreateBitableMsg(msg)
		return nil
		// —— 原有配置管理消息 ——
	case startEditConfigMsg:
		m.prepareEdit(false)
		return nil
	case startRefreshConfigMsg:
		m.prepareEdit(true)
		return nil
	default:
		return nil
	}
}

// prepareEdit 初始化编辑流程
func (m *model) prepareEdit(refresh bool) {
	m.isRefresh = refresh
	m.cursor = 0
	m.editStage = stagePickKey
	m.editKeys = make([]string, 0, len(m.configData))
	for k := range m.configData {
		m.editKeys = append(m.editKeys, k)
	}
}

// handleMenuKey 响应菜单界面按键
func (m *model) handleMenuKey(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch key := msg.(type) {
	case tea.KeyMsg:
		switch key.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < m.optionCount()-1 {
				m.cursor++
			}
		case "enter":
			return m.handleEnter()
		case "esc", "left", "h":
			m.current = menuMain
			m.cursor = 0
		}
	}
	return m, nil
}

// handleEnter 响应菜单选择确认
func (m *model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.current {
	case menuMain:
		return m.enterMain()
	case menuBitable:
		return m.enterBitable()
	case menuConfig:
		return m.enterConfig()
	default:
		m.current = menuMain
		m.cursor = 0
		return m, nil
	}
}

// enterMain 主菜单行为
func (m *model) enterMain() (tea.Model, tea.Cmd) {
	switch m.cursor {
	case 0:
		m.current = menuBitable
	case 1:
		m.current = menuTable
	case 2:
		m.current = menuField
	case 3:
		m.current = menuRecord
	case 4:
		m.current = menuConfig
	case 5:
		return m, tea.Quit
	}
	m.cursor = 0
	return m, nil
}

// enterConfig 配置管理行为
func (m *model) enterConfig() (tea.Model, tea.Cmd) {
	switch m.cursor {
	case 0:
		// 查看配置，由 View 渲染
	case 1:
		return m, editConfigCmd()
	case 2:
		return m, refreshConfigCmd()
	case 3:
		_ = config.Reset("local")
		if cfg, err := config.LoadMap("local"); err == nil {
			m.configData = cfg
		}
		m.statusMsg = "🔄 已重置为默认配置"
	case 4:
		m.current = menuMain
		m.statusMsg = ""
	}
	m.cursor = 0
	return m, nil
}

// enterBitable 多维表格子菜单行为
func (m *model) enterBitable() (tea.Model, tea.Cmd) {
	switch m.cursor {
	case 0:
		m.statusMsg = "🔍 功能待实现：查看已有多维表格"
	case 1:
		// 进入“输入名称”状态
		m.editStage = stageEnterBitableName
		m.textInput.Placeholder = "请输入表格名称"
		m.textInput.SetValue("")
		m.textInput.Focus()
	case 2:
		m.statusMsg = "🗑️ 功能待实现：删除多维表格"
	case 3:
		m.current = menuMain
		m.statusMsg = ""
	}
	m.cursor = 0
	return m, nil
}

// updateEditing 编辑流程：选键或输入新值
func (m *model) updateEditing(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.editStage {
	case stagePickKey:
		return m.handlePickKey(msg)
	case stageEnterValue:
		return m.handleEnterValue(msg)
	case stageEnterBitableName:
		return m.handleEnterBitableName(msg)
	default:
		return m, nil
	}
}

// handlePickKey 配置键选择
func (m *model) handlePickKey(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.editKeys)-1 {
				m.cursor++
			}
		case "enter":
			selected := m.editKeys[m.cursor]
			if m.isRefresh {
				val, err := config.RefreshKey(selected)
				if err == nil {
					m.configData[selected] = val
					_ = config.SaveMap("local", m.configData)
					m.statusMsg = fmt.Sprintf("✅ %s 刷新为 %s", selected, val)
				} else {
					m.statusMsg = fmt.Sprintf("❌ 刷新 %s 失败：%v", selected, err)
				}
				m.editStage = stageNone
				m.current = menuConfig
				m.cursor = 2
			} else {
				m.editKey = selected
				m.textInput.SetValue(fmt.Sprintf("%v", m.configData[selected]))
				m.textInput.Focus()
				m.editStage = stageEnterValue
			}
		case "esc":
			m.editStage = stageNone
		}
	}
	return m, nil
}

// handleEnterValue 输入新值并保存
func (m *model) handleEnterValue(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
		val := m.textInput.Value()
		m.configData[m.editKey] = val
		_ = config.SaveMap("local", m.configData)
		m.statusMsg = fmt.Sprintf("✅ %s 修改为 %s", m.editKey, val)
		m.editStage = stageNone
		m.current = menuConfig
		m.cursor = 0
	}
	return m, cmd
}

// handleEnterBitableName 负责“多维表格名称输入”流程
func (m *model) handleEnterBitableName(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// 文本框更新
	m.textInput, cmd = m.textInput.Update(msg)

	// Esc——取消
	if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
		m.editStage = stageNone
		m.current = menuBitable
		m.cursor = 1 // “创建” 那一项
		return m, nil
	}

	// Enter——发起创建
	if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
		name := m.textInput.Value()
		return m, createBitableCmd(name, "")
	}

	return m, cmd
}

// 启动编辑与刷新流程消息
func editConfigCmd() tea.Cmd    { return func() tea.Msg { return startEditConfigMsg{} } }
func refreshConfigCmd() tea.Cmd { return func() tea.Msg { return startRefreshConfigMsg{} } }

// 消息结构体定义
type startEditConfigMsg struct{}
type startRefreshConfigMsg struct{}
