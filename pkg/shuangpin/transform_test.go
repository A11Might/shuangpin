package shuangpin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaturalCode(t *testing.T) {
	pinyins := []string{
		"shi",
		"zhong",
		"guo",
		"ren",
		"o",
	}
	want := []string{
		"ui",
		"vs",
		"go",
		"rf",
		"oo",
	}
	n := NewTransform(NaturalCode)
	for i, pinyin := range pinyins {
		shuangpin := n.Pinyin2Shuangpin(pinyin)
		assert.Equal(t, want[i], shuangpin)
	}
}

func TestFlyPy(t *testing.T) {
	pinyins := []string{
		"sao",
		"mie",
	}
	want := []string{
		"sc",
		"mp",
	}
	n := NewTransform(FlyPY)
	for i, pinyin := range pinyins {
		shuangpin := n.Pinyin2Shuangpin(pinyin)
		assert.Equal(t, want[i], shuangpin)
	}
}
