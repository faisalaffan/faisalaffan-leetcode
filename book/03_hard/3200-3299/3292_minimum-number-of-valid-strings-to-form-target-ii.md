# 3292 — Minimum Number Of Valid Strings To Form Target Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minValidStrings(words []string, target string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Trie, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3292: Minimum Number of Valid Strings to Form Target II
// https://leetcode.com/problems/minimum-number-of-valid-strings-to-form-target-ii/
// Difficulty: Hard
//
// A "valid string" is any substring of any word in the words array.
// Build a trie containing ALL suffixes of all words, so that every
// possible valid substring is a prefix of some path in the trie.
// Then DP: dp[i] = min number of valid strings to form target[i:].
// For each position i, walk the trie to find all valid substrings
// starting at i, updating dp[i] = min(1 + dp[i+len]).

import (
	"fmt"
	"math"
)

type trieNode struct {
	child [26]*trieNode
}

func main() {
	// Example 1: words=["abc","aaaaa","bcdef"], target="aabcdabc" => 3
	fmt.Println(minValidStrings([]string{"abc", "aaaaa", "bcdef"}, "aabcdabc"))
	// Example 2: single word covers entire target
	fmt.Println(minValidStrings([]string{"hello"}, "hello"))
	// Example 3: no valid strings
	fmt.Println(minValidStrings([]string{"ab", "cd"}, "ef"))
	// Example 4: multiple chars needed
	fmt.Println(minValidStrings([]string{"a", "b", "c"}, "abc"))
	// Example 5: target empty
	fmt.Println(minValidStrings([]string{"a"}, ""))
	// Example 6: repetition
	fmt.Println(minValidStrings([]string{"aa", "a"}, "aaa"))
}

func minValidStrings(words []string, target string) int {
	if len(target) == 0 {
		return 0
	}

	// Build a trie containing all suffixes of all words.
	// This way every substring of any word is a prefix of some path.
	root := &trieNode{}
	for _, w := range words {
		// Insert every suffix of w into the trie.
		for start := 0; start < len(w); start++ {
			node := root
			for i := start; i < len(w); i++ {
				idx := w[i] - 'a'
				if node.child[idx] == nil {
					node.child[idx] = &trieNode{}
				}
				node = node.child[idx]
			}
		}
	}

	n := len(target)
  // Alokasi slice integer
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		node := root
		// Walk the trie to find all valid substrings starting at i.
		for j := i; j < n; j++ {
			idx := target[j] - 'a'
			if node.child[idx] == nil {
				break
			}
			node = node.child[idx]
			// Every node reached represents a valid substring target[i:j+1].
			if dp[j+1] != math.MaxInt32 {
				candidate := 1 + dp[j+1]
				if candidate < dp[i] {
					dp[i] = candidate
				}
			}
		}
	}

	if dp[0] == math.MaxInt32 {
		return -1
	}
	return dp[0]
}
```
