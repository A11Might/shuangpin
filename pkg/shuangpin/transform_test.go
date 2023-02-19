package shuangpin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaturalCode(t *testing.T) {
	pinyins := [][]string{
		{"sh", "i"},
		{"zh", "ong"},
		{"g", "uo"},
		{"r", "en"},
		{"", "o"},
	}
	want := [][]string{
		{"u", "i"},
		{"v", "s"},
		{"g", "o"},
		{"r", "f"},
		{"o", "o"},
	}
	n := NewTransform(NaturalCode)
	for i, pinyin := range pinyins {
		shuangpin := n.Shengyun2Shuangpin(pinyin)
		assert.Equal(t, want[i], shuangpin)
	}
}

func TestFlyPy(t *testing.T) {
	pinyins := [][]string{
		{"s", "ao"},
		{"m", "ie"},
	}
	want := [][]string{
		{"s", "c"},
		{"m", "p"},
	}
	n := NewTransform(FlyPY)
	for i, pinyin := range pinyins {
		shuangpin := n.Shengyun2Shuangpin(pinyin)
		assert.Equal(t, want[i], shuangpin)
	}
}
