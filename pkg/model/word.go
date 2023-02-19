package model

import (
	"math/rand"
	"time"

	"github.com/A11Might/shuangpin/pkg/shuangpin"
)

type Word struct {
	Word       string // "中"
	Pinyin     string // "zhong"
	Shuangpyin string // "vs"

	rand      *rand.Rand
	Transform *shuangpin.Transform
}

func NewRandomWord(scheme shuangpin.ShuangpinScheme) *Word {
	w := Word{
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		Transform: shuangpin.NewTransform(scheme),
	}
	w.getRandomWord()
	return &w
}

func (w *Word) Next() {
	w.getRandomWord()
}

func (w *Word) getRandomWord() {
	// 随机取汉字
	sheng := shuangpin.Dict[w.rand.Intn(len(shuangpin.Dict))]
	yun := sheng.Yuns[w.rand.Intn(len(sheng.Yuns))]
	// 重新生成汉字双拼
	w.Word = yun.Word
	w.Pinyin = sheng.Sheng + yun.Yun
	w.Shuangpyin = w.Transform.Pinyin2Shuangpin(w.Pinyin)
}
