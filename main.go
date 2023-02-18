package main

import (
	"fmt"
	"os"

	"github.com/A11Might/shuangpin/pkg/model"
	"github.com/A11Might/shuangpin/pkg/shuangpin"
	"github.com/A11Might/shuangpin/pkg/util"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mozillazg/go-pinyin"
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
	chineses := util.HandlingText(shuangpin.GenerateChinese("你好", 100))
	words := make([]*model.Word, 0, len([]rune(chineses)))
	for _, word := range []rune(chineses) {
		if !util.Symbol(string(word)) || string(word) == "" {
			pinyins := pinyin.LazyConvert(string(word), nil)
			shuangpin := shuangpin.Pinyin2NaturalCode(pinyins[0])
			words = append(words, &model.Word{
				Word:       string(word),
				Pinyin:     pinyins[0],
				Shuangpyin: shuangpin,
			})
		}
	}
	return model.Model{
		Word:     words,
		Index:    0,
		KeyBoard: model.NewKeyBoard(),
		Typed:    "",
	}
}
