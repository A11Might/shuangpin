package shuangpin

var (
	// 自然码键位映射
	naturalCodePinyinToKey = map[string]string{
		"q":   "q",
		"iu":  "q",
		"w":   "w",
		"ia":  "w",
		"ua":  "w",
		"e":   "e",
		"r":   "r",
		"uan": "r",
		"t":   "t",
		"ue":  "t",
		"ve":  "t",
		"y":   "y",
		"ing": "y",
		"uai": "y",
		"sh":  "u",
		"u":   "u",
		"ch":  "i",
		"i":   "i",
		"o":   "o",
		"uo":  "o",
		"p":   "p",
		"un":  "p",

		"a":    "a",
		"s":    "s",
		"iong": "s",
		"ong":  "s",
		"d":    "d",
		"iang": "d",
		"uang": "d",
		"f":    "f",
		"en":   "f",
		"eng":  "g",
		"g":    "g",
		"h":    "h",
		"ang":  "h",
		"j":    "j",
		"an":   "j",
		"k":    "k",
		"ao":   "k",
		"l":    "l",
		"ai":   "l",

		"z":   "z",
		"ei":  "z",
		"x":   "x",
		"ie":  "x",
		"c":   "c",
		"iao": "c",
		"zh":  "v",
		"ui":  "v",
		"v":   "v",
		"b":   "b",
		"ou":  "b",
		"n":   "n",
		"in":  "n",
		"m":   "m",
		"ian": "m",
	}

	// 零声母自然码键位映射
	naturalCodeSpecialPinyinToKey = map[string][]string{
		"a":   {"a", "a"},
		"e":   {"e", "e"},
		"o":   {"o", "o"},
		"ai":  {"a", "i"},
		"ei":  {"e", "i"},
		"ou":  {"o", "u"},
		"an":  {"a", "n"},
		"en":  {"e", "n"},
		"ao":  {"a", "o"},
		"er":  {"e", "r"},
		"ang": {"a", "h"},
		"eng": {"e", "g"},
	}
)
