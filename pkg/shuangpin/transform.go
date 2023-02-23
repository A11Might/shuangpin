package shuangpin

type ShuangpinScheme string

const (
	NaturalCode ShuangpinScheme = "zrm" // 自然码
	FlyPY       ShuangpinScheme = "xh"  // 小鹤双拼
	SouGou      ShuangpinScheme = "sg"  // 搜狗双拼
	MS          ShuangpinScheme = "wr"  // 微软双拼
)

var (
	// 双拼方案
	shuangpinMap = map[ShuangpinScheme]*shuangpin{
		NaturalCode: {
			Name:              "自然码",
			pinyin2Key:        naturalCodePinyinToKey,
			specialPinyin2Key: naturalCodeSpecialPinyinToKey,
		},
		FlyPY: {
			Name:              "小鹤双拼",
			pinyin2Key:        flypyPinyinToKey,
			specialPinyin2Key: flypySpecialPinyinToKey,
		},
		SouGou: {
			Name:              "搜狗双拼",
			pinyin2Key:        sougouPinyinToKey,
			specialPinyin2Key: sougouSpecialPinyinToKey,
		},
		MS: {
			Name:              "微软双拼",
			pinyin2Key:        msPinyinToKey,
			specialPinyin2Key: msSpecialPinyinToKey,
		},
	}
)

type shuangpin struct {
	Name              string
	pinyin2Key        map[string]string
	specialPinyin2Key map[string][]string
}

type Transform struct {
	*shuangpin
}

func NewTransform(scheme ShuangpinScheme) *Transform {
	sp, ok := shuangpinMap[scheme]
	if !ok {
		panic("unsupported shuangpin type")
	}
	t := Transform{sp}
	return &t
}

// Shengyun2Shuangpin 声母韵母转化为双拼键位
// ["zh", "ong"] to ["v", "s"]
// ["g", "uo"] to ["g", "o"]
// ["r", "en"] to ["r", "f"]
// ["", "o"] to ["o", "o"]
func (t *Transform) Shengyun2Shuangpin(shengyun []string) []string {
	if shengyun[0] == "" {
		return t.specialPinyin2Key[shengyun[1]]
	}

	return []string{t.pinyin2Key[shengyun[0]], t.pinyin2Key[shengyun[1]]}
}
