package util

import "strings"

// conversion [["zhong"] ["guo"] ["ren"]] to ["zhong", "guo", "ren"]
func conversion(twoD [][]string) []string {
	result := make([]string, 0)
	for _, item := range twoD {
		result = append(result, item[0])
	}
	return result
}

// symbol 判断是否是符号
func symbol(str string) bool {
	if str == "," || str == "." {
		return true
	}
	return false
}

// handlingText ，/。 to ,/.
func handlingText(text string) string {
	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "，", "", -1)
	text = strings.Replace(text, "。", "", -1)
	text = strings.Replace(text, ",", "", -1)
	text = strings.Replace(text, ".", "", -1)
	text = strings.Replace(text, "·", "", -1)
	return text
}
