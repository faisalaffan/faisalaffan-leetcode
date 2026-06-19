package main

// LeetCode #676: Implement Magic Dictionary
// https://leetcode.com/problems/implement-magic-dictionary/
// Difficulty: Medium
// Time: O(n * L) for buildDict, O(26 * L) for search
// Space: O(n * L)

import "fmt"

func main() {
	md := MagicDictionary{}
	md.BuildDict([]string{"hello", "leetcode"})
	fmt.Println(md.Search("hello"))
	fmt.Println(md.Search("hhllo"))
	fmt.Println(md.Search("hell"))
	fmt.Println(md.Search("leetcoded"))
}

type MagicDictionary struct {
	words []string
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
	m.words = dictionary
}

func (m *MagicDictionary) Search(searchWord string) bool {
	for _, word := range m.words {
		if len(word) != len(searchWord) {
			continue
		}
		diff := 0
		for i := 0; i < len(word); i++ {
			if word[i] != searchWord[i] {
				diff++
			}
			if diff > 1 {
				break
			}
		}
		if diff == 1 {
			return true
		}
	}
	return false
}
