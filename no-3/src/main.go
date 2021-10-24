package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findFirstStringInBracket("Hello, (play)ground)"))
}
func findFirstStringInBracket(str string) string {
	indexFirstBracketFound := strings.Index(str, "(")
	if len(str) > 0 && indexFirstBracketFound >= 0 {
		runes := []rune(str)
		wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
		indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
		return checkIndexClosingBracketFound(indexClosingBracketFound, wordsAfterFirstBracket)

	} else {
		return ""
	}

}
func checkIndexClosingBracketFound(indexClosingBracketFound int, wordsAfterFirstBracket string) string {
	if indexClosingBracketFound >= 0 {
		runes := []rune(wordsAfterFirstBracket)
		return string(runes[1 : indexClosingBracketFound-1])
	} else {
		return ""
	}
}
