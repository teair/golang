package go_3

import (
	"fmt"
	"unicode/utf8"
)

func runes() {
	s := "Yes我爱中国!" // UTF-8

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}

	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)

	for len(bytes) > 0 {

		ch, size := utf8.DecodeRune(bytes)

		// 获取当前字符之后的元素并返回给bytes直到bytes为空
		bytes = bytes[size:]

		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	for i, ch := range []rune(s) { // rune => int(32)
		fmt.Printf("(%d %c) ", i, ch)
	}

}
