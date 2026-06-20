# 1525 — Number Of Good Ways To Split A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func NumSplits(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(N), Space: O(1) (26 chars)  |  **Ruang:** O(1) (26 chars)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1525: Number of Good Ways to Split a String
// https://leetcode.com/problems/number-of-good-ways-to-split-a-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumSplits("aacaba"))
	fmt.Println(NumSplits("abcd"))
	fmt.Println(NumSplits("aaaaa"))
}

func NumSplits(s string) int {
	// Time: O(N), Space: O(1) (26 chars)
	n := len(s)
  // Alokasi slice
	leftCount := make([]int, 26)
  // Alokasi slice
	rightCount := make([]int, 26)
	leftUnique := 0
	rightUnique := 0

	// Initialize right side
	for i := 0; i < n; i++ {
		idx := s[i] - 'a'
		if rightCount[idx] == 0 {
			rightUnique++
		}
		rightCount[idx]++
	}

	result := 0
	for i := 0; i < n-1; i++ {
		idx := s[i] - 'a'
		// Move char from right to left
		if leftCount[idx] == 0 {
			leftUnique++
		}
		leftCount[idx]++

		rightCount[idx]--
		if rightCount[idx] == 0 {
			rightUnique--
		}

		if leftUnique == rightUnique {
			result++
		}
	}

	return result
}
```
