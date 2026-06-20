# 1639 — Number Of Ways To Form A Target String Given A Dictionary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numWays(words []string, target string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1639: Number of Ways to Form a Target String Given a Dictionary
// https://leetcode.com/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/
// Difficulty: Hard

import "fmt"

const MOD1639 = 1000000007

func numWays(words []string, target string) int {
	m := len(words[0]) // number of columns
	n := len(target)   // target length

	// count[c][ch] = number of words with character ch at column c
  // Alokasi slice
	count := make([][26]int, m)
	for _, w := range words {
		for c, ch := range w {
			count[c][ch-'a']++
		}
	}

	// dp[j] = number of ways to form first j characters of target
  // Alokasi slice
	dp := make([]int, n+1)
	dp[0] = 1

	for c := 0; c < m; c++ {
		// Process right-to-left to avoid using the same column twice
		for j := n; j > 0; j-- {
			ch := target[j-1] - 'a'
			if count[c][ch] > 0 {
				dp[j] = (dp[j] + dp[j-1]*count[c][ch]) % MOD1639
			}
		}
	}

	return dp[n]
}

func main() {
	// Test case 1: words=["acca","bbbb","caca"], target="aba" -> 6
	words := []string{"acca", "bbbb", "caca"}
	target := "aba"
	result := numWays(words, target)
	fmt.Printf("words=%v target=%s -> %d (expected 6)\n", words, target, result)

	// Test case 2: words=["abba","baab"], target="bab" -> 4
	words2 := []string{"abba", "baab"}
	target2 := "bab"
	result2 := numWays(words2, target2)
	fmt.Printf("words=%v target=%s -> %d (expected 4)\n", words2, target2, result2)

	// Test case 3: words=["abcd"], target="abcd" -> 1
	words3 := []string{"abcd"}
	target3 := "abcd"
	result3 := numWays(words3, target3)
	fmt.Printf("words=%v target=%s -> %d (expected 1)\n", words3, target3, result3)

	// Test case 4: words=["abcd"], target="ac" -> 1 (col 0='a', col 2='c')
	words4 := []string{"abcd"}
	target4 := "ac"
	result4 := numWays(words4, target4)
	fmt.Printf("words=%v target=%s -> %d (expected 1)\n", words4, target4, result4)
}
```
