package shuangpin

import "strings"

type ShuangpinScheme string

const (
	NaturalCode ShuangpinScheme = "zrm"    // 自然码
	FlyPY       ShuangpinScheme = "flypy"  // 小鹤双拼
	SouGou      ShuangpinScheme = "sougou" // 搜狗双拼
	MS          ShuangpinScheme = "ms"     // 微软双拼
)

var (
	// 声母，按顺序匹配（先 sh 再 s）
	initials = []string{"b", "p", "m", "f", "d", "t", "n", "l", "g", "k", "h", "j", "q", "x", "zh", "ch", "sh", "r", "z", "c", "s", "y", "w"}

	// 双拼方案
	shuangpinMap = map[ShuangpinScheme]*shuangpin{
		NaturalCode: {
			name:              "自然码",
			pinyin2Key:        naturalCodePinyinToKey,
			specialPinyin2Key: naturalCodeSpecialPinyinToKey,
		},
		FlyPY: {
			name:              "小鹤双拼",
			pinyin2Key:        flypyPinyinToKey,
			specialPinyin2Key: flypySpecialPinyinToKey,
		},
		SouGou: {
			name:              "搜狗双拼",
			pinyin2Key:        sougouPinyinToKey,
			specialPinyin2Key: sougouSpecialPinyinToKey,
		},
		MS: {
			name:              "微软双拼",
			pinyin2Key:        msPinyinToKey,
			specialPinyin2Key: msSpecialPinyinToKey,
		},
	}
)

type shuangpin struct {
	name              string
	pinyin2Key        map[string]string
	specialPinyin2Key map[string]string
}

type Transform struct {
	Name              string
	pinyin2Key        map[string]string
	specialPinyin2Key map[string]string
}

func NewTransform(scheme ShuangpinScheme) *Transform {
	shuangpin, ok := shuangpinMap[scheme]
	if !ok {
		panic("unsupported shuangpin type")
	}
	t := Transform{
		Name:              shuangpin.name,
		pinyin2Key:        shuangpin.pinyin2Key,
		specialPinyin2Key: shuangpin.specialPinyin2Key,
	}
	return &t
}

// Pinyin2Shuangpin 单个汉字拼音转化为双拼键位
// "zhong"(中) to "vs"
// "guo"(国) to "go"
// "ren"(人) to "rf"
// "o"(哦) to "oo"
func (b *Transform) Pinyin2Shuangpin(pinyin string) string {
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

// Pinyins2Shuangpins 多个汉字拼音转化为双拼键位
// ["zhong"(中), "guo"(国), "ren"(人), "o"(哦)] to ["vs", "go", "rf", "oo"]
func (n *Transform) Pinyins2Shuangpins(pinyins []string) []string {
	shuangpins := make([]string, 0, len(pinyins))
	for i := range pinyins {
		shuangpins = append(shuangpins, n.Pinyin2Shuangpin(pinyins[i]))
	}
	return shuangpins
}
