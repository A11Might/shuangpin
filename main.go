package main

import (
	"fmt"
	"github.com/A11Might/shuangpin/pkg/model"
	"github.com/A11Might/shuangpin/pkg/util"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"time"
)

const (
	blue   = "#4776E6"
	purple = "#8E54E9"
	usage  = `shuangpin`
)

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model.Model {
	bar := progress.NewModel(progress.WithScaledGradient(blue, purple))

	chinese, display, shuangpin := util.Text("东东你好", 100)

	return model.Model{
		Progress:      &bar,
		Text:          []rune(chinese),
		Display:       display,
		TextShuangpin: shuangpin,
		Start:         time.Now(),
	}
}
