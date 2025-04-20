package main

import (
	"Base/internal/config"
	create "Base/internal/field"
	"Base/internal/utils"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"

	"Base/internal/bitable"
	"Base/internal/record"
)

func main() {
	// 读取配置
	_, err := config.Load("local")
	if err != nil {
		fmt.Println("配置加载失败:", err)
		os.Exit(1)
	}

	// 初始界面
	utils.ClearScreen()
	utils.PrintHeader([]string{"主菜单"})

	for {
		mainChoice := ""
		prompt := &survey.Select{
			Message: "请选择模块：",
			Options: []string{
				"📁 表格结构构建",
				"🔁 数据写入与同步",
				"🔍 数据查询与验证",
				"⚙️ 配置与环境设置",
				"🚪 退出程序",
			},
		}
		if err := survey.AskOne(prompt, &mainChoice); err != nil {
			fmt.Println("退出:", err)
			return
		}

		switch mainChoice {
		case "📁 表格结构构建":
			handleStructure()
		case "🔁 数据写入与同步":
			handleInsert()
		case "🔍 数据查询与验证":
			handleQuery()
		case "⚙️ 配置与环境设置":
			utils.ClearScreen()
			utils.PrintHeader([]string{"配置与环境设置"})
			fmt.Println("（TODO）未来将支持配置切换与环境加载")
			pause()
		case "🚪 退出程序":
			fmt.Println("👋 感谢使用，再见！")
			os.Exit(0)
		}
	}
}

func handleStructure() {

	utils.ClearScreen()
	utils.PrintHeader([]string{"表格结构构建"})
	choice := ""
	_ = survey.AskOne(&survey.Select{
		Message: "表格结构构建：请选择操作",
		Options: []string{
			"创建多维表格",
			"添加字段",
			"查看字段列表",
			"返回主菜单",
		},
	}, &choice)

	switch choice {
	case "创建多维表格":
		utils.ClearScreen()
		utils.PrintHeader([]string{"表格结构构建", "创建多维表格"})
		bitable.CreateBitable()
		pause()
	case "添加字段":
		create.CreateFields()
	case "查看字段列表":
		fmt.Println("👉 查看字段结构（TODO）")
	case "返回主菜单":
		return
	}
}

func handleInsert() {
	utils.ClearScreen()
	utils.PrintHeader([]string{"数据写入与同步"})
	choice := ""
	_ = survey.AskOne(&survey.Select{
		Message: "数据写入与同步：请选择操作",
		Options: []string{
			"插入 mock 数据",
			"同步真实产品数据（TODO）",
			"返回主菜单",
		},
	}, &choice)

	switch choice {
	case "插入 mock 数据":
		record.InsertMock()
	case "同步真实产品数据（TODO）":
		fmt.Println("👉 正在建设中...")
	case "返回主菜单":
		return
	}
}

func handleQuery() {
	utils.ClearScreen()
	utils.PrintHeader([]string{"数据查询与验证"})
	choice := ""
	_ = survey.AskOne(&survey.Select{
		Message: "数据查询与验证：请选择操作",
		Options: []string{
			"查看所有记录",
			"条件筛选记录（TODO）",
			"返回主菜单",
		},
	}, &choice)

	switch choice {
	case "查看所有记录":
		record.QueryRecords()
	case "条件筛选记录（TODO）":
		fmt.Println("👉 正在建设中...")
	case "返回主菜单":
		return
	}
}

func pause() {
	fmt.Println("\n按 Enter 返回上一级菜单...")
	fmt.Scanln()
}
