package main

import (
	"fmt"
	"sort"
	"strings"
)

func Anagram(dummyDict []string) {
	list := make(map[string][]string)

	for _, word := range dummyDict {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		fmt.Println(words)

	}
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func main() {
	var dummyDict = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	Anagram(dummyDict)
}
