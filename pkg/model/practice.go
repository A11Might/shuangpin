package model

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	width = 66.

	// charsPerWord 2 个字母等于一个汉字
	charsPerWord = 2.
)

var (
	keyStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63"))
	keyPressStyle = keyStyle.Copy().
			Reverse(true)

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
)

type Model struct {
	Word     []*Word
	Index    int // 当前使用 Word 下标
	KeyBoard *KeyBoard
	Typed    string // 用户输入内容
}

type TickMsg time.Time

// 消除按键回显效果
func clearKeyPress() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (k KeyBoard) Init() tea.Cmd {
	return nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.KeyBoard.Hit(msg.String())

		switch msg.Type {
		case tea.KeyCtrlC:
			// 使用 Ctrl + c 退出程序
			return m, tea.Quit

		case tea.KeyBackspace:
			// 删除用户输入字符
			if m.Typed != "" {
				m.Typed = m.Typed[:len(m.Typed)-1]
			}
			return m, nil

		case tea.KeyTab, tea.KeyEnter:
			// 不输入 Tab, Enter 按键名
			return m, nil

		default:
			curWord := m.Word[m.Index]
			m.Typed += msg.String()
			// 判断用户输入是否正确：错误时清空；正确自动切换汉字
			if len(m.Typed) == 2 {
				if m.Typed == curWord.Shuangpyin {
					m.Index++
					if m.Index == len(m.Word) {
						m.Index--
						return m, tea.Quit
					}
				}
				m.Typed = ""
			}
			return m, clearKeyPress()
		}

	case TickMsg:
		m.KeyBoard.hit = defaultPosition
		return m, nil

	default:
		return m, nil
	}
}

func (m Model) View() string {
	doc := strings.Builder{}

	curWord := m.Word[m.Index]
	okButton := activeButtonStyle.Render(curWord.Word)
	cancelButton := buttonStyle.Render(m.Typed)
	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(curWord.Pinyin)
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)
	dialog := lipgloss.Place(width, 9,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("双拼"),
		lipgloss.WithWhitespaceForeground(subtle),
	)
	doc.WriteString(dialog + "\n\n")

	k := m.KeyBoard
	lines := make([]string, 0, len(k.keyboard))
	for i, rows := range k.keyboard {
		line := make([]string, 0, len(rows))
		for j, col := range rows {
			// 展示汉字双拼提示
			if strings.Contains(col, strings.ToUpper(string(curWord.Shuangpyin[0]))) ||
				strings.Contains(col, strings.ToUpper(string(curWord.Shuangpyin[1]))) {
				if i == k.hit.X && j == k.hit.Y {
					line = append(line, keyStyle.Copy().Background(lipgloss.Color("64")).Render(col))
				} else {
					line = append(line, keyStyle.Copy().Background(lipgloss.Color("63")).Render(col))
				}
			} else if i == k.hit.X && j == k.hit.Y {
				// 回显用户按键操作
				line = append(line, keyPressStyle.Render(col))
			} else {
				// 正常展示键盘按键
				line = append(line, keyStyle.Render(col))
			}
		}
		lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Center, line...))
	}
	doc.WriteString(lipgloss.JoinVertical(lipgloss.Center, lines...))

	return doc.String()
}
