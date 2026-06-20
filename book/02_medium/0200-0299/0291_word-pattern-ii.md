# 0291 — Word Pattern Ii

## Deskripsi

**Soal:** [0291. Word Pattern Ii](https://leetcode.com/problems/word-pattern-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(2^n) worst case, Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func wordPatternMatch(pattern string, s string) bool`

## Solusi Go

```go
package main

// LeetCode #291: Word Pattern II
// https://leetcode.com/problems/word-pattern-ii/
// Difficulty: Medium [Paid]
// Time: O(2^n) worst case, Space: O(n)

import "fmt"

func wordPatternMatch(pattern string, s string) bool {
  // Membuat map untuk pencarian O(1): key → value
	pMap := make(map[byte]string)
  // Membuat map untuk pencarian O(1): key → value
	sMap := make(map[string]byte)

	var backtrack func(patIdx, strIdx int) bool
	backtrack = func(patIdx, strIdx int) bool {
		if patIdx == len(pattern) && strIdx == len(s) {
			return true
		}
		if patIdx >= len(pattern) || strIdx >= len(s) {
			return false
		}

		ch := pattern[patIdx]
		if mapped, ok := pMap[ch]; ok {
			if strIdx+len(mapped) > len(s) || s[strIdx:strIdx+len(mapped)] != mapped {
				return false
			}
			return backtrack(patIdx+1, strIdx+len(mapped))
		}

		for end := strIdx + 1; end <= len(s); end++ {
			candidate := s[strIdx:end]
			if existing, ok := sMap[candidate]; ok && existing != ch {
				continue
			}

			pMap[ch] = candidate
			sMap[candidate] = ch
			if backtrack(patIdx+1, end) {
				return true
			}
			delete(pMap, ch)
			delete(sMap, candidate)
		}

		return false
	}

	return backtrack(0, 0)
}

func main() {
	fmt.Println(wordPatternMatch("abab", "redblueredblue"))
	fmt.Println(wordPatternMatch("aaaa", "asdasdasdasd"))
	fmt.Println(wordPatternMatch("ab", "aa"))
}
```
