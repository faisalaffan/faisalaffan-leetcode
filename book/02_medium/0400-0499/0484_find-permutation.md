# 0484 — Find Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindPermutation(s string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #484: Find Permutation
// https://leetcode.com/problems/find-permutation/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindPermutation("I"))
	fmt.Println(FindPermutation("DI"))
}

func FindPermutation(s string) []int {
	n := len(s) + 1
  // Alokasi slice
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i + 1
	}

	// Reverse contiguous segments for each 'D'
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] == 'D' {
			j := i
			for j < len(s) && s[j] == 'D' {
				j++
			}
			// Reverse segment from i to j
			left, right := i, j
  // Two-pointer loop
			for left < right {
				result[left], result[right] = result[right], result[left]
				left++
				right--
			}
			i = j
		}
	}

	return result
}
```
