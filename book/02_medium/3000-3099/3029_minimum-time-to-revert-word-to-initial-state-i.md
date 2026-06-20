# 3029 — Minimum Time To Revert Word To Initial State I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumTimeToInitialState(word string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3029: Minimum Time to Revert Word to Initial State I
// https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-i/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(minimumTimeToInitialState("abacaba", 3))
	fmt.Println(minimumTimeToInitialState("abacaba", 2))
	fmt.Println(minimumTimeToInitialState("abcbabcd", 2))
}

func minimumTimeToInitialState(word string, k int) int {
	n := len(word)
  // Alokasi slice
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && word[i] != word[j] {
			j = pi[j-1]
		}
		if word[i] == word[j] {
			j++
		}
		pi[i] = j
	}
	j := n
	for j > 0 && j > n-k {
		j = pi[j-1]
	}
	for t := 1; ; t++ {
		if t*k >= n {
			return t
		}
		if n-t*k <= j && (n-t*k)%k == 0 {
			return t
		}
	}
}
```
