package anagram

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func Anagram(first string, second string) {
	str := findFirstInReverseSecond(first, second)
	fmt.Print(str)
}

func findFirstInReverseSecond(first string, second string) string {
	var buf bytes.Buffer
	rev := reverse(second)
	if strings.Contains(first, rev) {
		buf.WriteString("Да")
	} else {
		buf.WriteString("Нет")
	}
	return buf.String()
}

func reverse(str string) string {
	rune := []rune(str)
	n := utf8.RuneCountInString(str)
	last := n - 1

	for i := 0; i < n/2; i++ {
		rune[i], rune[last-i] = rune[last-i], rune[i]
	}

	return string(rune)
}
