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
				Name:    "scheme",
				Aliases: []string{"s"},
				Value:   "zrm",
				Usage:   "choose shuangpin scheme",
			},
			&cli.StringFlag{
				Name:    "mode",
				Aliases: []string{"m"},
				Value:   "random",
				Usage:   "choose practice mode",
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
			p := tea.NewProgram(model.NewModel(cCtx.String("scheme"), cCtx.String("mode"), cCtx.Bool("pinyin"), cCtx.Bool("keyboard")))
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
				Usage:   "View the supported shuangpin schemes and practice mode",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(
						"support shuangpin scheme: ziranma(zrm), xiaohe(xh), sougou(sg), weiruan(wr)\n" +
							"support practice mode: sequence, random",
					)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
