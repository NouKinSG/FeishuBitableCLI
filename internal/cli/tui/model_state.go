package tui

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

// 模型状态
type model struct {
	cursor          int
	current         menu
	selectedBitable string
	selectedTable   string
}

// 初始化 Model
func initialModel() model {
	return model{
		cursor:  0,
		current: menuMain,
	}
}
