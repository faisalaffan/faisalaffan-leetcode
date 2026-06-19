package main

// LeetCode #966: Vowel Spellchecker
// https://leetcode.com/problems/vowel-spellchecker/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

// Time: O(n * L) | Space: O(n * L)
func spellchecker(wordlist []string, queries []string) []string {
	exact := make(map[string]bool)
	lower := make(map[string]string)
	vowel := make(map[string]string)

	for _, w := range wordlist {
		exact[w] = true
		lo := strings.ToLower(w)
		if _, ok := lower[lo]; !ok {
			lower[lo] = w
		}
		vw := devowel(lo)
		if _, ok := vowel[vw]; !ok {
			vowel[vw] = w
		}
	}

	ans := make([]string, len(queries))
	for i, q := range queries {
		if exact[q] {
			ans[i] = q
		} else if w, ok := lower[strings.ToLower(q)]; ok {
			ans[i] = w
		} else if w, ok := vowel[devowel(strings.ToLower(q))]; ok {
			ans[i] = w
		} else {
			ans[i] = ""
		}
	}
	return ans
}

func devowel(s string) string {
	var sb strings.Builder
	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			sb.WriteByte('*')
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func main() {
	fmt.Println(spellchecker([]string{"KiTe", "kite", "hare", "Hare"}, []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}))
	fmt.Println(spellchecker([]string{"yellow", "wood"}, []string{"Yello", "wood", "yellow"}))
}
