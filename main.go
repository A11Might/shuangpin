package main

import (
	"fmt"
	"os"

	"github.com/A11Might/shuangpin/pkg/model"
	"github.com/A11Might/shuangpin/pkg/shuangpin"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	blue   = "#4776E6"
	purple = "#8E54E9"
	usage  = `在你的命令行中练习双拼 :P

Usage: 
	shuangpin <command> [arguments]
Command:
	-t, --type: zrm, flypy
	           支持自然码、小鹤双拼
`
)

func main() {
	if len(os.Args) > 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Println(usage)
		os.Exit(0)
	}

	if len(os.Args) == 3 && os.Args[1] != "-t" && os.Args[1] != "--type" {
		fmt.Println(usage)
		os.Exit(1)
	}

	var spType shuangpin.ShuangpinType
	if len(os.Args) == 3 && (os.Args[1] == "-t" || os.Args[1] == "--type") {
		spType = shuangpin.ShuangpinType(os.Args[2])
	}

	p := tea.NewProgram(initialModel(spType))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel(spType shuangpin.ShuangpinType) model.Model {
	return model.Model{
		Word:     model.NewRandomWord(spType),
		KeyBoard: model.NewKeyBoard(),
		Typed:    "",
	}
}
