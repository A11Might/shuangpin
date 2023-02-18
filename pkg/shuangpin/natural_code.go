package shuangpin

import (
	"strings"
)

type NaturalCode struct {
}

// PinyinToShuangpin
// "zhong"(中) to "vs"
// "guo"(国) to "go"
// "ren"(人) to "rf"
// "o"(哦) to "oo"
func (n *NaturalCode) PinyinToShuangpin(pinyin string) string {
	var builder strings.Builder
	switch pinyin {
	// 零声母
	case "a", "e", "o":
		builder.WriteString(pinyin + pinyin)

	case "ai", "ei", "ou", "an", "en", "ao", "er":
		builder.WriteString(pinyin)

	case "ang":
		builder.WriteString("ah")

	case "eng":
		builder.WriteString("eg")

	default:
		// 声母+韵母
		for _, initial := range initials {
			if strings.HasPrefix(pinyin, initial) {
				pre := naturalCodePinyinToKey[initial]                // 声母
				succ := naturalCodePinyinToKey[pinyin[len(initial):]] // 韵母
				builder.WriteString(pre + succ)
				break
			}
		}
	}

	return builder.String()
}

func (n *NaturalCode) PinyinsToShuangpins(pinyins []string) []string {
	shuangpins := make([]string, 0, len(pinyins))
	for i := range pinyins {
		shuangpins = append(shuangpins, n.PinyinToShuangpin(pinyins[i]))
	}
	return shuangpins
}

var defaultNaturalCode = &NaturalCode{}

func Pinyin2NaturalCode(pinyin string) string {
	return defaultNaturalCode.PinyinToShuangpin(pinyin)
}

func Pinyins2NaturalCodes(pinyin []string) []string {
	return defaultNaturalCode.PinyinsToShuangpins(pinyin)
}
