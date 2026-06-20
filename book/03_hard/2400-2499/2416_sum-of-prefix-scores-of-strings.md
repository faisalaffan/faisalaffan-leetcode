# 2416 — Sum Of Prefix Scores Of Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func sumPrefixScores(words []string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Trie, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Trie** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2416: Sum of Prefix Scores of Strings
// https://leetcode.com/problems/sum-of-prefix-scores-of-strings/
// Difficulty: Hard
//
// Trie with visit count. Insert each word, counting how many words pass through
// each prefix node. Then for each word, sum the counts along its prefix path.
// Time O(N * L) | Space O(N * L) where N = len(words), L = avg word length.

import "fmt"

func main() {
	// Example 1
	fmt.Println(sumPrefixScores([]string{"abc", "ab", "bc", "b"}))
	// Example 2
	fmt.Println(sumPrefixScores([]string{"abcd"}))
	// Edge: single char
	fmt.Println(sumPrefixScores([]string{"a", "a", "a"}))
}

type trieNode struct {
	children [26]*trieNode
	count    int
}

func sumPrefixScores(words []string) []int {
	root := &trieNode{}

	// Insert all words into trie, incrementing count at each node
	for _, w := range words {
		cur := root
  // Linear scan O(n)
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			if cur.children[idx] == nil {
				cur.children[idx] = &trieNode{}
			}
			cur = cur.children[idx]
			cur.count++
		}
	}

  // Alokasi slice
	ans := make([]int, len(words))
	for wi, w := range words {
		cur := root
		total := 0
  // Linear scan O(n)
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			cur = cur.children[idx]
			total += cur.count
		}
		ans[wi] = total
	}
	return ans
}
```
