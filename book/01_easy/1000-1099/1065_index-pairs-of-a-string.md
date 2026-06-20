# 1065 — Index Pairs Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func indexPairs(text string, words []string) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Trie

**Waktu:** O(n^2 + m*k)  |  **Ruang:** O(m*k) for trie

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1065: Index Pairs of a String
// https://leetcode.com/problems/index-pairs-of-a-string/
// Difficulty: Easy [Paid]
// Time: O(n^2 + m*k) | Space: O(m*k) for trie

import "fmt"

func main() {
	fmt.Println(indexPairs("thestoryofleetcodeandme", []string{"story", "fleet", "leetcode"}))
	// [[1,5],[3,7],[10,13],[10,18]]
	fmt.Println(indexPairs("ababa", []string{"aba", "ab"}))
	// [[0,1],[0,2],[2,3],[2,4]]
}

// LeetCode submission: indexPairs
func indexPairs(text string, words []string) [][]int {
  // HashMap: O(1) lookup
	wordSet := make(map[string]bool)
	for _, w := range words {
		wordSet[w] = true
	}
	var ans [][]int
  // Linear scan O(n)
	for i := 0; i < len(text); i++ {
		for j := i; j < len(text); j++ {
			if wordSet[text[i:j+1]] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
```
