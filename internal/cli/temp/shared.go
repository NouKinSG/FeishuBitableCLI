package cli

import (
	"Base/internal/utils"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func askSelect(message string, options []string) string {
	var choice string
	_ = survey.AskOne(&survey.Select{
		Message: message,
		Options: options,
	}, &choice)
	return choice
}

func pause() {
	fmt.Println("\n按 Enter 返回...")
	fmt.Scanln()
}

func printInfo(msg string) {
	fmt.Println(msg)
}

func printError(msg string) {
	fmt.Printf("❗ %s\n", msg)
}

func displayOr(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

func printHeader(path []string) {
	fmt.Println("====================================================")
	fmt.Println("🚀 飞书多维表格 CLI 工具 v0.1.0")
	fmt.Printf("📂 当前表格：%s\n", displayOr(currentBitableName, "未选择"))
	fmt.Printf("📑 当前数据表：%s\n", displayOr(currentTableName, "未选择"))
	if len(path) > 0 {
		fmt.Printf("📍 当前路径：%s\n", strings.Join(path, " > "))
	}
	fmt.Println("====================================================")
}

func ClearScreen() {
	utils.ClearScreen()
}
