# 3137 — Minimum Number Of Operations To Make Word K Periodic

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumOperationsToMakeKPeriodic(word string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n / k)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3137: Minimum Number of Operations to Make Word K-Periodic
// https://leetcode.com/problems/minimum-number-of-operations-to-make-word-k-periodic/
// Difficulty: Medium
// Time: O(n) | Space: O(n / k)

import "fmt"

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)
  // HashMap: O(1) lookup
	freq := make(map[string]int)
	maxFreq := 0

	for i := 0; i < n; i += k {
		sub := word[i : i+k]
		freq[sub]++
		if freq[sub] > maxFreq {
			maxFreq = freq[sub]
		}
	}

	return n/k - maxFreq
}

func main() {
	fmt.Println(minimumOperationsToMakeKPeriodic("leetcodeleet", 4)) // Expected: 1
	fmt.Println(minimumOperationsToMakeKPeriodic("abcabcabc", 3))    // Expected: 0
	fmt.Println(minimumOperationsToMakeKPeriodic("aabbccddee", 5))   // Expected: 1
}
```
