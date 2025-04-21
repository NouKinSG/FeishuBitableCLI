package main

import (
	"Base/internal/cli/tui"
	"fmt"
	"os"

	"Base/internal/config"
)

func main() {
	// 读取配置
	if _, err := config.Load("local"); err != nil {
		fmt.Println("配置加载失败:", err)
		os.Exit(1)
	}
	tui.StartTUI()

}
