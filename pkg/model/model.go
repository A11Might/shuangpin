package model

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

const (
	width = 60.

	// charsPerWord is the average characters per word used by most typing tests
	// to calculate your WPM score.
	charsPerWord = 5.
)

type Model struct {
	// Percent is a value from 0 to 1 that represents the current completion of the typing test
	Percent  float64
	Progress *progress.Model
	// Text is the randomly generated text for the user to type
	Text []rune
	// Display is Text's pinyin
	Display []string
	// TextShuangpin is Text's shuangpin
	TextShuangpin string
	// Typed is the text that the user has typed so far
	Typed string
	// Start and end are the start and end time of the typing test
	Start time.Time
	// Mistakes is the number of characters that were mistyped by the user
	Mistakes int
	// Score is the user's score calculated by correct characters typed
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
	//remaining := m.TextShuangpin[len(m.Typed):]

	var chinese, curChinese, pinyin, typed string
	var idx int
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
		if i%2 != 0 {
			chinese += curChinese
			idx++
		}
	}
	remaining := string(m.Text[idx:])

	s := fmt.Sprintf("\n  %s\n\n%s%s", m.Progress.ViewAs(m.Percent), chinese, termenv.String(remaining).Faint())
	s += fmt.Sprintf("\n%s", pinyin)
	//s += fmt.Sprintf("\n%s", typed)
	s += fmt.Sprintf("\n%s", m.TextShuangpin)
	s += fmt.Sprintf("\n\nWPM: %.2f\n", (m.Score/charsPerWord)/(time.Since(m.Start).Minutes()))

	return s
}
