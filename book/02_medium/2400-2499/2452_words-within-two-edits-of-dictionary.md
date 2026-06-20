# 2452 — Words Within Two Edits Of Dictionary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func twoEditWords(queries []string, dictionary []string) []string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m * L)  |  **Ruang:** O(1) where L = word length


## 💻 Solusi Go

```go
package main

// LeetCode #2452: Words Within Two Edits of Dictionary
// https://leetcode.com/problems/words-within-two-edits-of-dictionary/
// Difficulty: Medium
// Time: O(n * m * L) | Space: O(1) where L = word length
// For each query word, check Hamming distance to each dictionary word.

import "fmt"

func main() {
	fmt.Println(twoEditWords([]string{"word", "note", "ants", "wood"}, []string{"wood", "joke", "moat"})) // ["word","note","wood"]
	fmt.Println(twoEditWords([]string{"yes"}, []string{"not"}))                                            // []
}

func twoEditWords(queries []string, dictionary []string) []string {
	ans := make([]string, 0)
	for _, q := range queries {
		for _, d := range dictionary {
			if hammingDist(q, d) <= 2 {
				ans = append(ans, q)
				break
			}
		}
	}
	return ans
}

func hammingDist(a, b string) int {
	dist := 0
  // Linear scan O(n)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}
```
