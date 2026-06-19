package main

// LeetCode #1451: Rearrange Words in a Sentence
// https://leetcode.com/problems/rearrange-words-in-a-sentence/
// Difficulty: Medium

import "fmt"
import "sort"
import "strings"
import "unicode"

func main() {
	// Test case 1
	fmt.Println(arrangeWords("Leetcode is cool")) // "Is cool leetcode"

	// Test case 2
	fmt.Println(arrangeWords("Keep calm and code on")) // "On and keep calm code"

	// Test case 3
	fmt.Println(arrangeWords("To be or not to be")) // "To be or to be not"
}

type wordInfo struct {
	word   string
	index  int
	length int
}

// Time: O(n log n) where n = number of words
// Space: O(n) for storing words
func arrangeWords(text string) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}

	// Lowercase first word
	runes := []rune(words[0])
	runes[0] = unicode.ToLower(runes[0])
	words[0] = string(runes)

	infos := make([]wordInfo, len(words))
	for i, w := range words {
		infos[i] = wordInfo{w, i, len(w)}
	}

	sort.Slice(infos, func(i, j int) bool {
		if infos[i].length != infos[j].length {
			return infos[i].length < infos[j].length
		}
		return infos[i].index < infos[j].index
	})

	result := make([]string, len(infos))
	for i, info := range infos {
		result[i] = info.word
	}

	// Uppercase first letter
	firstWord := result[0]
	runes = []rune(firstWord)
	runes[0] = unicode.ToUpper(runes[0])
	result[0] = string(runes)

	return strings.Join(result, " ")
}
