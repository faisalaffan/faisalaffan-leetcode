package main

// LeetCode #290: Word Pattern
// https://leetcode.com/problems/word-pattern/
// Difficulty: Easy

import "strings"
import "fmt"

// Time: O(n) | Space: O(n)
func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	p2w := make(map[byte]string)
	w2p := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w := words[i]
		if mappedW, ok := p2w[p]; ok && mappedW != w {
			return false
		}
		if mappedP, ok := w2p[w]; ok && mappedP != p {
			return false
		}
		p2w[p] = w
		w2p[w] = p
	}
	return true
}

func main() {
	fmt.Println(WordPattern("abba", "dog cat cat dog"))
	fmt.Println(WordPattern("abba", "dog cat cat fish"))
	fmt.Println(WordPattern("aaaa", "dog cat cat dog"))
}
