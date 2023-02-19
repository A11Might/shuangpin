package model

import (
	"math/rand"
	"strings"
	"time"

	"github.com/A11Might/shuangpin/pkg/shuangpin"
)

type practiceMode string

const (
	Sequence practiceMode = "sequence"
	Random   practiceMode = "random"
)

type Word struct {
	Word       string   // "中"
	Shengyun   []string // ["zh", "ong"]
	Pinyin     string   // "zhong"
	Shuangpyin []string // ["v", "s"]
	Answer     string   // "vs"

	// mode
	mode     practiceMode // 练习模式
	position *position    // 当前单词位置（全部顺序模式使用）

	rand      *rand.Rand
	Transform *shuangpin.Transform
}

func NewWord(scheme shuangpin.ShuangpinScheme, mode practiceMode) *Word {
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

// Next 根据练习模式获取下一个汉字
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
	x := w.rand.Intn(len(shuangpin.Dict))
	y := w.rand.Intn(len(shuangpin.Dict[x].Yuns))
	w.genWord(x, y)
}

// 顺序取汉字
func (w *Word) getNextWord() {
	w.genWord(w.position.X, w.position.Y)

	// 计算下一个汉字位置
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

// 根据 Dict 坐标更新 Word 中的汉字
func (w *Word) genWord(x, y int) {
	sheng := shuangpin.Dict[x]
	yun := sheng.Yuns[y]

	w.Word = yun.Word
	w.Shengyun = []string{sheng.Sheng, yun.Yun}
	w.Pinyin = strings.Join(w.Shengyun, "")
	w.Shuangpyin = w.Transform.Shengyun2Shuangpin(w.Shengyun)
	w.Answer = strings.Join(w.Shuangpyin, "")
}
