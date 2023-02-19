package model

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Prompt  key.Binding
	Confirm key.Binding
	Quit    key.Binding
	Help    key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Help}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Prompt},  // first column
		{k.Confirm}, // second column
		{k.Quit},    // fourth column
		{k.Help},    // third column
	}
}
