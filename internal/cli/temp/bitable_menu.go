package cli

import (
	"Base/internal/bitable"
)

func manageBitable() {
	ClearScreen()
	printHeader([]string{"主菜单", "多维表格管理"})

	choice := askSelect("多维表格管理：请选择操作", []string{
		"1) 查看已有多维表格列表",
		"2) 创建新的多维表格",
		"3) 切换多维表格",
		"4) 删除多维表格",
		"5) 返回上一级",
	})

	switch choice {
	case "1) 查看已有多维表格列表":
		printInfo("👉 待实现：列出多维表格")
	case "2) 创建新的多维表格":
		bitable.CreateBitable()
	case "3) 切换多维表格":
		printInfo("👉 待实现：切换多维表格")
	case "4) 删除多维表格":
		printInfo("👉 待实现：删除多维表格")
	case "5) 返回上一级":
		return
	}
	pause()
}
