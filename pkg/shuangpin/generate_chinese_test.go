package shuangpin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateChinese(t *testing.T) {
	s := GenerateChinese("你好", 100)
	assert.NotNil(t, s)
}
