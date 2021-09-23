package chatbot

import "github.com/yanyiwu/gojieba"

var DefaultJieba = gojieba.NewJieba()

func Cut(s string) []string {
	return DefaultJieba.Cut(s, true)
}

func CutMap(s string) map[string]int {
	w := Cut(s)
	m := map[string]int{}
	for _, v := range w {
		m[v] += 1
	}
	return m
}
