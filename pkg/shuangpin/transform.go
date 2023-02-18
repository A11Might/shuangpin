package shuangpin

import "strings"

// 声母，按顺序匹配（先 sh 再 s）
var initials = []string{"b", "p", "m", "f", "d", "t", "n", "l", "g", "k", "h", "j", "q", "x", "zh", "ch", "sh", "r", "z", "c", "s", "y", "w"}

type ITransform interface {
	Type() string
	// 单个汉字拼音转化为双拼键位
	// "zhong"(中) to "vs"
	// "guo"(国) to "go"
	// "ren"(人) to "rf"
	// "o"(哦) to "oo"
	Pinyin2Shuangpin(string) string
	// 多个汉字拼音转化为双拼键位
	// ["zhong"(中), "guo"(国), "ren"(人), "o"(哦)] to ["vs", "go", "rf", "oo"]
	Pinyins2Shuangpins([]string) []string
}

var _ ITransform = (*transform)(nil)

type transform struct {
	name              string
	pinyin2Key        map[string]string
	specialPinyin2Key map[string]string
}

type ShuangpinType string

const (
	NaturalCodeT ShuangpinType = "zrm"
	FlyPYT       ShuangpinType = "flypy"
)

func NewTransform(spType ...ShuangpinType) ITransform {
	if len(spType) == 0 || spType[0] == "" {
		spType = []ShuangpinType{NaturalCodeT}
	}

	t := new(transform)
	switch spType[0] {
	case NaturalCodeT:
		t.name = "自然码"
		t.pinyin2Key = naturalCodePinyinToKey
		t.specialPinyin2Key = naturalCodeSpecialPinyinToKey

	case FlyPYT:
		t.name = "小鹤双拼"
		t.pinyin2Key = flypyPinyinToKey
		t.specialPinyin2Key = flypySpecialPinyinToKey

	default:
		panic("unsupported shuangpin type")
	}
	return t
}

func (b *transform) Type() string {
	return string(b.name)
}

func (b *transform) Pinyin2Shuangpin(pinyin string) string {
	var builder strings.Builder
	switch pinyin {
	case "a", "e", "o", "ai", "ei", "ou", "an", "en", "ao", "er", "ang", "eng":
		// 零声母
		builder.WriteString(b.specialPinyin2Key[pinyin])

	default:
		// 声母 + 韵母
		for _, initial := range initials {
			if strings.HasPrefix(pinyin, initial) {
				pre := b.pinyin2Key[initial]                // 声母
				succ := b.pinyin2Key[pinyin[len(initial):]] // 韵母
				builder.WriteString(pre + succ)
				break
			}
		}
	}

	return builder.String()
}

func (n *transform) Pinyins2Shuangpins(pinyins []string) []string {
	shuangpins := make([]string, 0, len(pinyins))
	for i := range pinyins {
		shuangpins = append(shuangpins, n.Pinyin2Shuangpin(pinyins[i]))
	}
	return shuangpins
}
