package cli

var (
	currentBitableToken string
	currentBitableName  string
	currentTableID      string
	currentTableName    string
)

func Run() {
	ClearScreen()
	printHeader([]string{"主菜单"})

	for {
		choice := askSelect("请选择操作：", []string{
			"🗂 多维表格选择 & 管理",
			"🗄 数据表选择 & 管理",
			"🧱 字段操作",
			"📝 记录操作",
			"⚙️ 配置与环境",
			"🚪 退出",
		})

		switch choice {
		case "🗂 多维表格选择 & 管理":
			manageBitable()
		case "🗄 数据表选择 & 管理":
			if currentBitableToken == "" {
				printError("请先在“多维表格选择 & 管理”中选择或创建表格")
			} else {
				manageTable()
			}
		case "🧱 字段操作":
			if currentBitableToken == "" || currentTableID == "" {
				printError("请先选择表格和数据表")
			} else {
				manageField()
			}
		case "📝 记录操作":
			if currentBitableToken == "" || currentTableID == "" {
				printError("请先选择表格和数据表")
			} else {
				manageRecord()
			}
		case "⚙️ 配置与环境":
			printInfo("（TODO）配置管理待实现")
		case "🚪 退出":
			printInfo("👋 再见！")
			return
		}
	}
}
