package tui

import (
	"Base/internal/config"
	"github.com/charmbracelet/bubbles/textinput"
)

// 菜单状态定义
type menu int

const (
	menuMain menu = iota
	menuBitable
	menuTable
	menuField
	menuRecord
	menuConfig
)

// 编辑状态流
type editStage int

const (
	stageNone       editStage = iota
	stagePickKey              // 先让用户用↑↓选择要编辑的配置键
	stageEnterValue           // 再让用户在文本框里输入新值
)

// 模型状态
type model struct {
	cursor          int
	current         menu
	selectedBitable string
	selectedTable   string

	// 新增，用于配置管理
	configData map[string]interface{}

	// 编辑状态流相关
	// 以下是编辑流程相关
	editStage editStage       // 当前处于哪个编辑阶段
	editKeys  []string        // 配置键集合，用于“选择键”
	editKey   string          // 正在编辑的键名
	textInput textinput.Model // 文本输入组件
	statusMsg string          // 操作完成后的提示消息
	isRefresh bool            // 当前是在刷新模式
}

// 初始化 Model
func initialModel() model {
	m := model{
		cursor:  0,
		current: menuMain,
	}
	// 预加载配置到 m.configData
	if cfg, err := config.LoadMap("local"); err == nil {
		m.configData = cfg
	} else {
		m.configData = make(map[string]interface{})
	}
	// 初始化 textInput
	ti := textinput.New()
	ti.Placeholder = ""
	ti.CharLimit = 256
	ti.Width = 40
	m.textInput = ti
	return m

}
