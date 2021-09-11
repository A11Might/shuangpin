package shuangpin

import (
	"fmt"
	"github.com/A11Might/shuangpin/pkg/util"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

func Text(title string, length int) (string, [][]string, string) {
	chinese := util.HandlingText(generateChinese(title, length))
	pinyin := chineseToPinyin(chinese)
	split, shuangpin := pinyinToShuangpin(pinyin)
	return chinese, split, shuangpin
}

func TextWithSymbol(title string, length int) (string, [][]string, string) {
	chinese := util.HandlingText(generateChinese(title, length))
	pinyin := util.Conversion(chineseToPinyinWithSymbol(chinese))
	split, shuangpin := pinyinToShuangpin(pinyin)
	return chinese, split, shuangpin
}

// chineseToPinyin "中国人" to ["zhong", "guo", "ren"]
func chineseToPinyin(chinese string) []string {
	return pinyin.LazyConvert(chinese, nil)
}

func chineseToPinyinWithSymbol(chinese string) [][]string {
	a := pinyin.NewArgs()
	a.Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}
	return pinyin.Pinyin(chinese, a)
}

// pinyinToShuangpin "zhong gou ren" to ["zh", "ong", "g", "ou", "r", "en"], "vsgorf"
func pinyinToShuangpin(pinyin []string) ([][]string, string) {
	split := make([][]string, 0)
	var builder strings.Builder
	for i := range pinyin {
		if util.Symbol(pinyin[i]) {
			split = append(split, []string{pinyin[i]})
			builder.WriteString(pinyin[i])
		} else {
			pre, succ, shuangpin := transform(pinyin[i])
			split = append(split, []string{pre, succ})
			builder.WriteString(shuangpin)
		}
	}
	return split, builder.String()
}

func transform(pinyin string) (string, string, string) {
	switch len(pinyin) {
	case 1:
		return pinyin, pinyin, pinyin + pinyin

	case 2:
		return string(pinyin[0]), string(pinyin[1]), pinyin

	default:
		return transformCore(pinyin)
	}
}

// transformCore "zhong" to "zh", "ong", "vs"
func transformCore(pinyin string) (string, string, string) {
	for _, initial := range getInitials() {
		if strings.HasPrefix(pinyin, initial) {
			pinyinToKey := getNaturalCodePinyinToKey()
			pre := pinyinToKey[initial]
			succ := pinyinToKey[pinyin[len(initial):]]
			return initial, pinyin[len(initial):], pre + succ
		}
	}
	return "", "", ""
}

func example() {
	hans := "中国人"

	// 默认
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong] [guo] [ren]]

	// 包含声调
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhōng] [guó] [rén]]

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng] [guo2] [re2n]]

	// 开启多音字模式
	a = pinyin.NewArgs()
	a.Heteronym = true
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong zhong] [guo] [ren]]
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng zho4ng] [guo2] [re2n]]

	fmt.Println(pinyin.LazyPinyin(hans, pinyin.NewArgs()))
	// [zhong guo ren]

	fmt.Println(pinyin.Convert(hans, nil))
	// [[zhong] [guo] [ren]]

	fmt.Println(pinyin.LazyConvert(hans, nil))
	// [zhong guo ren]
}
