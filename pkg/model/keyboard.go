package model

var (
	// QWERTYKeyBoard keyboard layout
	QWERTYKeyBoard = [][]string{
		{" Q ", " W ", " E ", " R ", " T ", " Y ", " U ", " I ", " O ", " P ", " [{ ", " ]} ", " |\\ "},
		{" A ", " S ", " D ", " F ", " G ", " H ", " J ", " K ", " L ", " ;: ", " '\" "},
		{" Z ", " X ", " C ", " V ", " B ", " N ", " M ", " ,< ", " .> ", " /? "},
		{"                    "},
	}
	QWERTYKeyIndex = map[string]*position{
		"q": {0, 0}, "w": {0, 1}, "e": {0, 2}, "r": {0, 3}, "t": {0, 4}, "y": {0, 5}, "u": {0, 6}, "i": {0, 7}, "o": {0, 8}, "p": {0, 9}, "[": {0, 10}, "]": {0, 11}, "|": {0, 12},
		"Q": {0, 0}, "W": {0, 1}, "E": {0, 2}, "R": {0, 3}, "T": {0, 4}, "Y": {0, 5}, "U": {0, 6}, "I": {0, 7}, "O": {0, 8}, "P": {0, 9}, "{": {0, 10}, "}": {0, 11}, "\\": {0, 12},
		"a": {1, 0}, "s": {1, 1}, "d": {1, 2}, "f": {1, 3}, "g": {1, 4}, "h": {1, 5}, "j": {1, 6}, "k": {1, 7}, "l": {1, 8}, ";": {1, 9}, "'": {1, 10},
		"A": {1, 0}, "S": {1, 1}, "D": {1, 2}, "F": {1, 3}, "G": {1, 4}, "H": {1, 5}, "J": {1, 6}, "K": {1, 7}, "L": {1, 8}, ":": {1, 9}, "\"": {1, 10},
		"z": {2, 0}, "x": {2, 1}, "c": {2, 2}, "v": {2, 3}, "b": {2, 4}, "n": {2, 5}, "m": {2, 6}, ",": {2, 7}, ".": {2, 8}, "/": {2, 9},
		"Z": {2, 0}, "X": {2, 1}, "C": {2, 2}, "V": {2, 3}, "B": {2, 4}, "N": {2, 5}, "M": {2, 6}, "<": {2, 7}, ">": {2, 8}, "?": {2, 9},
		" ": {3, 0},
	}

	defaultPosition = &position{X: -1, Y: -1}
)

type layout string

const (
	QWERTY  layout = "QWERTY"
	DVORAK  layout = "DVORAK"  // unsupported
	COLEMAK layout = "COLEMAK" // unsupported
)

type KeyBoard struct {
	keyboard [][]string
	hit      *position
}

type position struct {
	X int
	Y int
}

func NewKeyBoard(keyBoardLayout ...layout) (keyboard *KeyBoard) {
	if len(keyBoardLayout) == 0 || keyBoardLayout[0] == "" {
		// default keyboard layout
		keyBoardLayout = []layout{QWERTY}
	}

	keyboard = &KeyBoard{hit: defaultPosition}
	switch keyBoardLayout[0] {
	case QWERTY:
		keyboard.keyboard = QWERTYKeyBoard
	default:
		panic("unsupported keyboard layout")
	}
	return
}

// Hit 显示按键回显
func (k *KeyBoard) Hit(keyMsg string) {
	k.hit = QWERTYKeyIndex[keyMsg]
	if k.hit == nil {
		// 没有展示的按键不回显
		k.hit = defaultPosition
	}
}
