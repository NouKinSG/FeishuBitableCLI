package cli

import "Base/internal/bitable"

func manageField() {
	ClearScreen()
	printHeader([]string{"主菜单", "字段操作"})

	choice := askSelect("字段操作：请选择操作", []string{
		"➕ 添加字段",
		"📋 查看字段列表",
		"⬅️ 返回上一级",
	})

	switch choice {
	case "➕ 添加字段":
		bitable.CreateFields()
	case "📋 查看字段列表":
		printInfo("👉 待实现：查看字段")
	case "⬅️ 返回上一级":
		return
	}
	pause()
}
