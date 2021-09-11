package util

import (
	"strings"
)

// Conversion [["zhong"] ["guo"] ["ren"]] to ["zhong", "guo", "ren"]
func Conversion(twoD [][]string) []string {
	result := make([]string, 0)
	for _, item := range twoD {
		result = append(result, item[0])
	}
	return result
}

// Symbol 判断是否是符号
func Symbol(str string) bool {
	if str == "," || str == "." {
		return true
	}
	return false
}

// HandlingText ，/。 to ,/. 带符号的句子还在开发
func HandlingText(text string) string {
	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "，", "", -1)
	text = strings.Replace(text, "。", "", -1)
	text = strings.Replace(text, ",", "", -1)
	text = strings.Replace(text, ".", "", -1)
	text = strings.Replace(text, "·", "", -1)
	text = strings.Replace(text, ":", "", -1)
	text = strings.Replace(text, "：", "", -1)
	text = strings.Replace(text, "?", "", -1)
	text = strings.Replace(text, "？", "", -1)
	text = strings.Replace(text, ";", "", -1)
	text = strings.Replace(text, "；", "", -1)
	text = strings.Replace(text, "x", "", -1)
	return text
}
