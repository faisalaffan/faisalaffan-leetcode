package main

// LeetCode #890: Find and Replace Pattern
// https://leetcode.com/problems/find-and-replace-pattern/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
	fmt.Println(FindAndReplacePattern([]string{"a", "b", "c"}, "a"))
	fmt.Println(FindAndReplacePattern([]string{"aa", "ab"}, "aa"))
}

// Time: O(n * m) where n = len(words), m = avg word length | Space: O(n)
func FindAndReplacePattern(words []string, pattern string) []string {
	var ans []string
	for _, word := range words {
		if isMatch(word, pattern) {
			ans = append(ans, word)
		}
	}
	return ans
}

func isMatch(word, pattern string) bool {
	if len(word) != len(pattern) {
		return false
	}
	w2p := make(map[byte]byte)
	p2w := make(map[byte]byte)

	for i := 0; i < len(word); i++ {
		wc, pc := word[i], pattern[i]
		if v, ok := w2p[wc]; ok && v != pc {
			return false
		}
		if v, ok := p2w[pc]; ok && v != wc {
			return false
		}
		w2p[wc] = pc
		p2w[pc] = wc
	}

	return true
}
