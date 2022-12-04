package main

import (
	"fmt"
	"os"

	"github.com/clh021/text-parser/tui"
	tuigithub "github.com/clh021/text-parser/tui-github"
	tuitodo "github.com/clh021/text-parser/tui-todo"
)

// 支持启动时显示构建日期和构建版本
// 需要通过命令 ` go build -ldflags "-X main.build=`git rev-parse HEAD`" ` 打包
var build = "not set"

func main() {
	fmt.Printf("Build: %s\n", build)

	if _, err := tuigithub.Run(); err != nil {
		// if _, err := tea.NewProgram(tui.NewModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}

	if _, err := tuitodo.Run(); err != nil {
		// if _, err := tea.NewProgram(tui.NewModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}

	if _, err := tui.Run(); err != nil {
		// if _, err := tea.NewProgram(tui.NewModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}
}
