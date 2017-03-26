package main

import (
	// "strings"
	"fmt"
)

func main() {

	type wordlistItem struct {
		word      string
		frequency int
	}

	var wordList []wordlistItem

	wordList = append(wordList, wordlistItem{word: "richtig", frequency: 8})
	wordList = append(wordList, wordlistItem{word: "genau", frequency: 18})

	for _, v := range wordList {
		fmt.Println(v.frequency)
		return
	}

}
