package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var debugMode = false

func SetDebugMode(enabled bool) {
	debugMode = enabled
}

func IsDebugMode() bool {
	return debugMode
}

func ClearScreen() {
	if debugMode {
		return
	}
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J")
	}
}

func PrintHeader(path []string) {
	if debugMode {
		if len(path) > 0 {
			if path[0] != "主菜单" {
				fmt.Printf("当前位置：主菜单 > %s\n", strings.Join(path, " > "))
			} else {
				fmt.Printf("当前位置：%s\n", strings.Join(path, " > "))
			}
		}
		return
	}
	fmt.Println("====================================================")
	fmt.Println("🚀 飞书多维表格 CLI 工具 v0.1.0")
	if len(path) > 0 {
		if path[0] != "主菜单" {
			fmt.Printf("当前位置：主菜单 > %s\n", strings.Join(path, " > "))
		} else {
			fmt.Printf("当前位置：%s\n", strings.Join(path, " > "))
		}
	}
	fmt.Println("====================================================")
}
