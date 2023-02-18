package shuangpin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaturalCode_PinyinToShuangpin(t *testing.T) {
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
	for i, pinyin := range pinyins {
		shuangpin := Pinyin2NaturalCode(pinyin)
		assert.Equal(t, want[i], shuangpin)
	}
}
