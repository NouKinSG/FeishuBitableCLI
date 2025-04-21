package tui

import (
	"Base/internal/config"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return nil
}

// Update 处理按键事件，并新增对 menuConfig 的处理
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// 优先处理自定义消息
	switch msg.(type) {
	case startEditConfigMsg:
		// 进入“选择键”阶段，将 configData 的键列表填充到 editKeys
		keys := make([]string, 0, len(m.configData))
		m.isRefresh = false
		keys = make([]string, 0, len(m.configData))

		for k := range m.configData {
			keys = append(keys, k)
		}
		m.editKeys = keys
		m.editStage = stagePickKey
		m.cursor = 0
		return m, nil

	case startRefreshConfigMsg:
		// 进入“选择键”阶段（刷新模式）
		m.isRefresh = true
		m.editKeys = make([]string, 0, len(m.configData))
		for k := range m.configData {
			m.editKeys = append(m.editKeys, k)
		}
		m.editStage = stagePickKey
		m.cursor = 0
		return m, nil

	default:
		// 非自定义消息，继续后续逻辑

	}

	// 如果处于编辑阶段，交给专门处理
	if m.editStage != stageNone {
		return m.updateEditing(msg)
	}

	// 否则走原有菜单逻辑

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
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
			switch m.current {
			case menuMain:
				// 主菜单跳转
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

			case menuConfig:
				// 配置管理子菜单
				switch m.cursor {
				case 0:
					// 查看当前配置：不修改 state，直接在 View 里渲染
				case 1:
					// 修改配置项：返回一个命令去弹交互框
					return m, editConfigCmd()
				case 2:
					// 刷新 Access Token
					return m, refreshTokenCmd()
				case 3:
					// 重置为默认配置并重新加载
					_ = config.Reset("local")
					if cfgMap, err := config.LoadMap("local"); err == nil {
						m.configData = cfgMap
					}
				case 4:
					// 返回主菜单，清除提示消息
					m.current = menuMain
					m.statusMsg = ""
				}
				m.cursor = 0

			default:
				// 其它所有二级菜单，按 Enter 或 Esc 都回主菜单
				m.current = menuMain
				m.cursor = 0
			}

		case "esc", "left", "h":
			// 通用的返回上一级
			m.current = menuMain
			m.cursor = 0
		}
	}
	return m, nil
}

// 定义两个自定义消息，启动编辑和启动刷新流程
type startEditConfigMsg struct{}
type startRefreshTokenMsg struct{}

type startRefreshConfigMsg struct{}

// editConfigCmd 返回一个 Cmd，用来进入“选择键”阶段
func editConfigCmd() tea.Cmd {
	return func() tea.Msg {
		return startEditConfigMsg{}
	}
}

// refreshTokenCmd 返回一个 Cmd，用来进入“刷新 token”流程
func refreshTokenCmd() tea.Cmd {
	return func() tea.Msg {
		return startRefreshTokenMsg{}
	}
}

func refreshConfigCmd() tea.Cmd {
	return func() tea.Msg {
		return startRefreshConfigMsg{}
	}
}

func (m model) updateEditing(msg tea.Msg) (model, tea.Cmd) {
	switch m.editStage {
	case stagePickKey:
		// 使用上下键在 editKeys 里切换
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
				// 选中了某个键
				m.editKey = m.editKeys[m.cursor]
				m.textInput.SetValue(fmt.Sprintf("%v", m.configData[m.editKey]))
				m.textInput.Focus()
				m.editStage = stageEnterValue

				// 用户确认了要操作的键
				selected := m.editKeys[m.cursor]
				if m.isRefresh {
					// 刷新模式：调用通用刷新函数
					newVal, err := config.RefreshKey(selected)
					if err == nil {
						m.configData[selected] = newVal
						_ = config.SaveMap("local", m.configData)
						m.statusMsg = fmt.Sprintf("✅ %s 刷新为 %s", selected, newVal)
					} else {
						m.statusMsg = fmt.Sprintf("❌ 刷新 %s 失败：%v", selected, err)
					}
					// 回到配置菜单
					m.editStage = stageNone
					m.current = menuConfig
					m.cursor = 2
				} else {
					// 编辑模式
					m.editKey = selected
					m.textInput.SetValue(fmt.Sprintf("%v", m.configData[selected]))
					m.textInput.Focus()
					m.editStage = stageEnterValue
				}
				return m, nil

			case "esc":
				// 取消编辑，回到配置菜单
				m.editStage = stageNone
			}
		}
		return m, nil

	case stageEnterValue:
		// 将键盘事件传给 textInput
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		// 如果用户按下 Enter，就保存
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
			newVal := m.textInput.Value()
			m.configData[m.editKey] = newVal
			_ = config.SaveMap("local", m.configData)
			m.statusMsg = fmt.Sprintf("✅ %s 修改为 %s", m.editKey, newVal)
			// 编辑完成，回到配置菜单
			m.editStage = stageNone
			m.current = menuConfig
			m.cursor = 0
		}
		return m, cmd
	}

	return m, nil
}

type startRefreshTenantMsg struct{}
type startRefreshUserMsg struct{}

func refreshTenantCmd() tea.Cmd {
	return func() tea.Msg { return startRefreshTenantMsg{} }
}
func refreshUserCmd() tea.Cmd {
	return func() tea.Msg { return startRefreshUserMsg{} }
}
