package main

import (
	"fmt"
	"strings"
)

func dealWrongWord(word string) string {
	word = strings.Trim(word, " ")
	if len(word) == 0 {
		return word
	}
	wordList := strings.Split(word, "")
	current := wordList[0]
	pNum := 0
	cNum := 1
	for i := 1; i < len(wordList); i++ {
		if wordList[i] == current {
			cNum++
		} else {
			pNum = cNum
			current = wordList[i]
			cNum = 1
		}
		if cNum > 2 || (cNum == 2 && pNum == 2) {
			wordList[i] = ""
			cNum--
		}
	}
	return strings.Join(wordList, "")
}

func checkWord() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}
	words := make([]string, n)
	for i := 0; i < n; i++ {
		var word string
		_, err := fmt.Scanln(&word)
		words[i] = dealWrongWord(word)
		if err != nil {
			return
		}
	}
	for _, word := range words {
		fmt.Println(word)
	}
}

func main() {
	checkWord()
}
