package shuangpin

import (
	"fmt"
	"github.com/A11Might/shuangpin/pkg/util"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

func Text(title string, length int) (string, []string, string) {
	chinese := util.HandlingText(generateChinese(title, length))
	pinyin := chineseToPinyin(chinese)
	display, shuangpin := pinyinToShuangpin(pinyin)
	return chinese, display, shuangpin
}

// chineseToPinyin "中国人" to ["zhong", "guo", "ren"]
func chineseToPinyin(chinese string) []string {
	return pinyin.LazyConvert(chinese, nil)
}

// pinyinToShuangpin "zhong gou ren" to ["zh", "ong", "g", "ou", "r", "en"], "vsgorf"
func pinyinToShuangpin(pinyin []string) ([]string, string) {
	display := make([]string, 0)
	var builder strings.Builder
	for i := range pinyin {
		pre, succ, shuangpin := transform(pinyin[i])
		display = append(display, pre, succ)
		builder.WriteString(shuangpin)
	}
	return display, builder.String()
}

func transform(pinyin string) (string, string, string) {
	switch len(pinyin) {
	case 1:
		if util.Symbol(pinyin) {
			return pinyin, pinyin, pinyin
		}
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
