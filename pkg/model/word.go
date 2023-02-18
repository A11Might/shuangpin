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
	Transform shuangpin.ITransform
}

func NewRandomWord(spType ...shuangpin.ShuangpinType) *Word {
	w := Word{
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		Transform: shuangpin.NewTransform(spType...),
	}
	w.getRandomWord()
	return &w
}

func (w *Word) Next() {
	w.getRandomWord()
}

func (w *Word) getRandomWord() {
	sheng := shuangpin.Dict[w.rand.Intn(len(shuangpin.Dict))]
	yun := sheng.Yuns[w.rand.Intn(len(sheng.Yuns))]
	pinyin := sheng.Sheng + yun.Yun
	// 万恶的指针
	*w = Word{
		Word:       yun.Word,
		Pinyin:     pinyin,
		Shuangpyin: w.Transform.Pinyin2Shuangpin(pinyin),
		rand:       w.rand,
		Transform:  w.Transform,
	}
}
