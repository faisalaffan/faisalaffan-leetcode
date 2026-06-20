# 0291 — Word Pattern Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func wordPatternMatch(pattern string, s string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Backtracking

**Waktu:** O(2^n) worst case, Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #291: Word Pattern II
// https://leetcode.com/problems/word-pattern-ii/
// Difficulty: Medium [Paid]
// Time: O(2^n) worst case, Space: O(n)

import "fmt"

func wordPatternMatch(pattern string, s string) bool {
  // HashMap: O(1) lookup
	pMap := make(map[byte]string)
  // HashMap: O(1) lookup
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
