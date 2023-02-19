package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/A11Might/shuangpin/pkg/shuangpin"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	width = 104.
)

var (
	// Dialog

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)
	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.Copy().
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				MarginRight(2).
				Underline(true)
	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	// Keyboard

	keyStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63"))
	keyPressStyle = keyStyle.Copy().
			Reverse(true)

	// Help

	keys = keyMap{
		Prompt: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("Tab", "显示答案"),
		),
		Confirm: key.NewBinding(
			key.WithKeys(" ", "enter"),
			key.WithHelp("Space/Enter", "切换或清空"),
		),
		Quit: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("ESC", "退出程序"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
	}
)

type Model struct {
	word     *Word
	keyBoard *KeyBoard
	typed    string // 用户输入内容
	quitting bool

	// help
	keys keyMap
	help help.Model

	// config
	disablePyPrompt bool // 禁用拼音提示
	disableKbPrompt bool // 禁用按键提示
}

func NewModel(scheme, mode string, disablePyPrompt, disableKbPrompt bool) Model {
	return Model{
		word:            NewWord(shuangpin.ShuangpinScheme(scheme), practiceMode(mode)),
		keyBoard:        NewKeyBoard(),
		disablePyPrompt: disablePyPrompt,
		disableKbPrompt: disableKbPrompt,

		keys: keys,
		help: help.New(),
	}
}

type TickMsg time.Time

// 100 毫秒后发送 TickMsg 命令，用于消除按键回显效果
func doTick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// 展示帮助
		if key.Matches(msg, m.keys.Help) {
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		}

		// 用户按键回显
		m.keyBoard.Hit(msg.String())

		// 判断用户输入是否正确：错误时清空；正确自动切换汉字
		// 字符不满数量时不清空，返回 false
		check := func() bool {
			if len(m.typed) == len(m.word.Shuangpyin) {
				if strings.ToLower(m.typed) == m.word.Answer {
					m.word.Next()
				}
				m.typed = ""
				return true
			}
			return false
		}

		switch msg.Type {
		case tea.KeyEsc:
			// 使用 Esc 退出程序
			m.quitting = true
			return m, tea.Quit

		case tea.KeyBackspace:
			// 删除用户输入字符
			if m.typed != "" {
				m.typed = m.typed[:len(m.typed)-1]
			}

		case tea.KeyTab:
			// Tab 显示答案
			m.typed = m.word.Answer

		case tea.KeySpace, tea.KeyEnter:
			// 空格/回车，切换或清空
			check()
			m.typed = ""

		default:
			// 当使用 Tab 显示答案后，再按键会超过 2 个字符，需要先 check 校验
			if !check() {
				// check 不通过，可能是不到两个字符，加上新键入的字符再次 check 校验
				if msg.String() != " " &&
					len(msg.String()) == 1 {
					// 空格及其他按键
					m.typed += msg.String()
				}
				check()
			}
		}
		return m, doTick()

	case TickMsg:
		// 消除按键回显效果
		m.keyBoard.hit = defaultPosition
		return m, nil

	default:
		return m, nil
	}
}

func (m Model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	doc := strings.Builder{}

	// Dialog
	{
		okButton := activeButtonStyle.Render(m.word.Word)
		cancelButton := buttonStyle.Render(m.typed)
		question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(m.word.Pinyin)
		if m.disablePyPrompt {
			question = lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("")
		}
		buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
		ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)
		dialog := lipgloss.Place(width, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars(m.word.Transform.Name),
			lipgloss.WithWhitespaceForeground(subtle),
		)
		doc.WriteString(dialog + "\n\n")
	}

	// KeyBoard
	{
		k := m.keyBoard
		lines := make([]string, 0, len(k.keyChar))
		for i, rows := range k.keyChar {
			line := make([]string, 0, len(rows))
			for j, char := range rows {
				// 展示汉字双拼提示
				if !m.disableKbPrompt &&
					(char == m.word.Shuangpyin[0] ||
						char == m.word.Shuangpyin[1]) {
					// 零声母汉字特殊处理
					if m.word.Shengyun[0] == "" {
						m.word.Shengyun[0] = string(m.word.Shengyun[1][0])
					}
					if i == k.hit.X && j == k.hit.Y {
						// 回显按键操作
						line = append(line, keyStyle.Copy().Background(lipgloss.Color("64")).Render(k.keyDisplay[i][j]))
					} else {
						// 双拼提示
						if char == string(m.word.Shuangpyin[0]) {
							line = append(line, keyStyle.Copy().Background(lipgloss.Color("63")).Render(fmt.Sprintf("%s%5s", strings.ToUpper(char), m.word.Shengyun[0])))
						} else {
							line = append(line, keyStyle.Copy().Background(lipgloss.Color("63")).Render(fmt.Sprintf("%s%5s", strings.ToUpper(char), m.word.Shengyun[1])))
						}
					}
				} else if i == k.hit.X && j == k.hit.Y {
					// 回显用户按键操作
					line = append(line, keyPressStyle.Render(k.keyDisplay[i][j]))
				} else {
					// 正常展示键盘按键
					line = append(line, keyStyle.Render(k.keyDisplay[i][j]))
				}
			}
			lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Center, line...))
		}
		doc.WriteString(lipgloss.JoinVertical(lipgloss.Center, lines...) + "\n")
	}

	// help
	{
		helpView := m.help.View(m.keys)
		doc.WriteString(helpView + "\n")
	}

	return doc.String()
}
