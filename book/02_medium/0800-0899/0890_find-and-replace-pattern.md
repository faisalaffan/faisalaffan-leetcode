# 0890 — Find And Replace Pattern

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindAndReplacePattern(words []string, pattern string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * m) where n = len(words), m = avg word length  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
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
  // HashMap: O(1) lookup
	w2p := make(map[byte]byte)
  // HashMap: O(1) lookup
	p2w := make(map[byte]byte)

  // Linear scan O(n)
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
```
