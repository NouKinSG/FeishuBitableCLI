package cli

func manageTable() {
	ClearScreen()
	printHeader([]string{"主菜单", "数据表管理"})

	choice := askSelect("数据表管理：请选择操作", []string{
		"1) 查看数据表列表",
		"2) 创建新的数据表",
		"3) 切换数据表",
		"4) 删除数据表",
		"5) 返回上一级",
	})

	switch choice {
	case "1) 查看数据表列表":
		printInfo("👉 待实现：列出数据表")
	case "2) 创建新的数据表":
		printInfo("👉 待实现：创建数据表")
	case "3) 切换数据表":
		printInfo("👉 待实现：切换数据表")
	case "4) 删除数据表":
		printInfo("👉 待实现：删除数据表")
	case "5) 返回上一级":
		return
	}
	pause()
}
