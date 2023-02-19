package main

import (
	"fmt"
	"log"
	"os"

	"github.com/A11Might/shuangpin/pkg/model"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "shuangpin",
		Usage: "Practice shuangpin in your terminal",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "type",
				Aliases: []string{"t"},
				Value:   "zrm",
				Usage:   "choose shuangpin scheme",
			},
			&cli.BoolFlag{
				Name:    "pinyin",
				Aliases: []string{"p"},
				Value:   false,
				Usage:   "disable pinyin prompt",
			},
			&cli.BoolFlag{
				Name:    "keyboard",
				Aliases: []string{"k"},
				Value:   false,
				Usage:   "disable key prompt",
			},
		},
		Action: func(cCtx *cli.Context) error {
			p := tea.NewProgram(model.NewModel(cCtx.String("type"), cCtx.Bool("pinyin"), cCtx.Bool("keyboard")))
			if _, err := p.Run(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "support",
				Aliases: []string{"s"},
				Usage:   "View the supported shuangpin schemes",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("支持自然码（zrm）、小鹤双拼（flypy）、搜狗双拼（sougou）、微软双拼（ms)")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
