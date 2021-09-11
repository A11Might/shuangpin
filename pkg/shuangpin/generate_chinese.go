package shuangpin

import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"
)

var data Data

type Data struct {
	Title  string   `json:"title"`
	Famous []string `json:"famous"`
	Bosh   []string `json:"bosh"`
	After  []string `json:"after"`
	Before []string `json:"before"`
}

func init() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	json.Unmarshal([]byte(getJsonDate()), &data)
}

func generateChinese(title string, length int) string {
	body := ""
	for len(body) < length {
		num := rand.Intn(100)
		if num < 10 {
			body += "\r\n"
		} else if num < 20 {
			sentence := choice(data.Famous)
			sentence = strings.Replace(sentence, "a", choice(data.Before), 1)
			sentence = strings.Replace(sentence, "b", choice(data.After), 1)
			body += sentence
		} else {
			body += choice(data.Bosh)
		}
		body = strings.Replace(body, "x", title, 1)
	}

	return body
}

func choice(sentences []string) string {
	return sentences[rand.Intn(len(sentences))]
}
