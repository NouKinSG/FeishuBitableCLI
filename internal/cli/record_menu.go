package cli

import "Base/internal/bitable"

func manageRecord() {
	ClearScreen()
	printHeader([]string{"主菜单", "记录操作"})

	choice := askSelect("记录操作：请选择操作", []string{
		"🧪 插入 Mock 数据",
		"🔍 查询全部记录",
		"⬅️ 返回上一级",
	})

	switch choice {
	case "🧪 插入 Mock 数据":
		bitable.InsertMock()
	case "🔍 查询全部记录":
		bitable.QueryRecords()
	case "⬅️ 返回上一级":
		return
	}
	pause()
}
