package shuangpin

type Shuangpin interface {
	PinyinToShuangpin(string) string
	PinyinsToShuangpins(string) string
}
