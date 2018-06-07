package eazy

import "strings"

//反转字符串中的单词 III
func reverseWords(s string) string {
	ls := strings.Split(s, " ")

	for i, item := range ls {
		ls[i] = reverseString(item)
	}

	result := strings.Join(ls, " ")
	return result
}

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func reverseWordsBetter(s string) string {
	ret := make([]byte, 0)
	word := strings.Split(s, " ")
	for idx, v := range word {
		for i := len(v) - 1; i >= 0; i-- {
			ret = append(ret, byte(v[i]))
		}
		if idx < len(word)-1 {
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}
