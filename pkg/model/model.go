package model

import (
	"fmt"
	"github.com/A11Might/shuangpin/pkg/shuangpin"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

const (
	width = 60.

	// charsPerWord 2 个字母等于一个汉字
	charsPerWord = 2.
)

type Model struct {
	// Percent 代表当前打字练习的进度, 值的范围为 [0, 1]
	Percent  float64
	Progress *progress.Model
	// Text 随机生成的用于用户练习的中文文本
	Text []rune
	// Display Text 的拼音
	Display []string
	// TextShuangpin Text 的双拼
	TextShuangpin string
	// Typed 用户到目前为止的输入
	Typed string
	// Notice 用户当前按键代表的声母或韵母
	Notice []string
	// Start 和 end 是打字练习开始和结束的时间
	Start time.Time
	// Mistakes 用户错误输入的字符数量
	Mistakes int
	// Score 用户正确输入的字符数量
	Score float64
}

// Init inits the bubbletea model for use
func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) updateProgress() (tea.Model, tea.Cmd) {
	m.Percent = float64(len(m.Typed)) / float64(len(m.TextShuangpin))
	if m.Percent >= 1.0 {
		return m, tea.Quit
	}
	return m, nil
}

// Update updates the bubbletea model by handling the progress bar update
// and adding typed characters to the state if they are valid typing characters
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// User wants to cancel the typing test
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

		// Deleting characters
		if msg.Type == tea.KeyBackspace && len(m.Typed) > 0 {
			m.Typed = m.Typed[:len(m.Typed)-1]
			return m.updateProgress()
		}

		// Ensure we are adding characters only that we want the user to be able to type
		if msg.Type != tea.KeyRunes {
			return m, nil
		}

		char := msg.Runes[0]
		next := rune(m.TextShuangpin[len(m.Typed)])

		// To properly account for line wrapping we need to always insert a new line
		// Where the next line starts to not break the user interface, even if the user types a random character
		if next == '\n' {
			m.Typed += "\n"

			// Since we need to perform a line break
			// if the user types a space we should simply ignore it.
			if char == ' ' {
				return m, nil
			}
		}

		m.Typed += msg.String()
		m.Notice = shuangpin.GetShuangpin(string(char))

		if char == next {
			m.Score += 1.
		}

		return m.updateProgress()
	case tea.WindowSizeMsg:
		m.Progress.Width = msg.Width - 4
		if m.Progress.Width > width {
			m.Progress.Width = width
		}
		return m, nil

	default:
		return m, nil
	}
}

// View shows the current state of the typing test.
// It displays a progress bar for the progression of the typing test,
// the typed characters (with errors displayed in red) and remaining
// characters to be typed in a faint display
func (m Model) View() string {
	var chinese, curChinese, pinyin, typed string
	var count int
	for i, c := range m.Typed {
		if c == rune(m.TextShuangpin[i]) {
			curChinese = string(m.Text[i/2])
			pinyin += m.Display[i]
			typed += string(c)
		} else {
			curChinese = termenv.String(string(m.Text[i/2])).Background(termenv.ANSIBrightRed).String()
			pinyin += termenv.String(m.Display[i]).Background(termenv.ANSIBrightRed).String()
			typed += termenv.String(string(m.Typed[i])).Background(termenv.ANSIBrightRed).String()
		}
		// 两个字符一个汉字
		if i%2 != 0 {
			chinese += curChinese
			count++
		}
	}
	remainChinese := string(m.Text[count:])
	remainAnswer := m.TextShuangpin[len(m.Typed):]

	s := fmt.Sprintf("\n  %s\n\n%s%s", m.Progress.ViewAs(m.Percent), chinese, termenv.String(remainChinese).Faint())
	s += fmt.Sprintf("\n%s", pinyin)
	s += fmt.Sprintf("\n%v", m.Notice)
	s += fmt.Sprintf("\n%s%s", typed, termenv.String(remainAnswer).Faint())
	s += fmt.Sprintf("\n\nWPM: %.2f\n", (m.Score/charsPerWord)/(time.Since(m.Start).Minutes()))

	return s
}
