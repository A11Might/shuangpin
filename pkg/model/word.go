package model

import (
	"math/rand"
	"time"

	"github.com/A11Might/shuangpin/pkg/shuangpin"
)

type practiceMode string

const (
	Sequence practiceMode = "sequence"
	Random   practiceMode = "random"
)

type Word struct {
	Word       string // "中"
	Pinyin     string // "zhong"
	Shuangpyin string // "vs"

	// mode
	mode     practiceMode // 练习模式
	position *position    // 当前单词位置

	rand      *rand.Rand
	Transform *shuangpin.Transform
}

func NewRandomWord(scheme shuangpin.ShuangpinScheme, mode practiceMode) *Word {
	w := Word{
		mode:      mode,
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		Transform: shuangpin.NewTransform(scheme),
	}

	switch mode {
	case Sequence:
		w.position = firstPosition
	}

	w.Next()
	return &w
}

func (w *Word) Next() {
	switch w.mode {
	case Sequence:
		w.getNextWord()
	case Random:
		w.getRandomWord()
	}
}

// 随机取汉字
func (w *Word) getRandomWord() {
	sheng := shuangpin.Dict[w.rand.Intn(len(shuangpin.Dict))]
	yun := sheng.Yuns[w.rand.Intn(len(sheng.Yuns))]

	// 重新生成汉字双拼
	w.Word = yun.Word
	w.Pinyin = sheng.Sheng + yun.Yun
	w.Shuangpyin = w.Transform.Pinyin2Shuangpin(w.Pinyin)
}

// 顺序取汉字
func (w *Word) getNextWord() {
	sheng := shuangpin.Dict[w.position.X]
	yun := sheng.Yuns[w.position.Y]

	// 重新生成汉字双拼
	w.Word = yun.Word
	w.Pinyin = sheng.Sheng + yun.Yun
	w.Shuangpyin = w.Transform.Pinyin2Shuangpin(w.Pinyin)

	// 获取下一个汉字位置
	w.position.Y++
	c := len(shuangpin.Dict[w.position.X].Yuns)
	if w.position.Y == c {
		// 到达最后一个字，从下一行开始
		w.position.Y = 0
		w.position.X++
	}
	r := len(shuangpin.Dict)
	if w.position.X == r {
		// 到达最后一行，从第一行开始
		w.position.X = 0
	}
}
