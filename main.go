package main

import (
	"fmt"
	"github.com/A11Might/shuangpin/pkg/model"
	"github.com/A11Might/shuangpin/pkg/shuangpin"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"time"
)

const (
	blue   = "#4776E6"
	purple = "#8E54E9"
	usage  = `shuangpin [daan]
daan: 显示双拼的答案
Examples:
 - shuangpin
 - shuangpin daan`
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Println(usage)
		os.Exit(0)
	}

	if len(os.Args) == 2 && os.Args[1] != "daan" {
		fmt.Println(usage)
		os.Exit(1)
	}

	var answer bool
	if len(os.Args) == 2 && os.Args[1] == "daan" {
		answer = true
	}

	p := tea.NewProgram(initialModel(answer))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel(answer bool) model.Model {
	bar := progress.NewModel(progress.WithScaledGradient(blue, purple))

	chinese, display, shuangpin := shuangpin.TextWithSymbol("东东你好", 100)

	return model.Model{
		Progress:      &bar,
		Text:          []rune(chinese),
		Split:         display,
		TextShuangpin: shuangpin,
		Answer:        answer,
		Start:         time.Now(),
	}
}
