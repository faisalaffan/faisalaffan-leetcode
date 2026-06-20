# 1375 — Number Of Times Binary String Is Prefix Aligned

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numTimesAllBlue(flips []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n) where n = length of flips  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1375: Number of Times Binary String Is Prefix-Aligned
// https://leetcode.com/problems/number-of-times-binary-string-is-prefix-aligned/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numTimesAllBlue([]int{3, 2, 4, 1, 5})) // 2

	// Test case 2
	fmt.Println(numTimesAllBlue([]int{4, 1, 2, 3})) // 1

	// Test case 3
	fmt.Println(numTimesAllBlue([]int{2, 1, 3})) // 1
}

// Time: O(n) where n = length of flips
// Space: O(1)
func numTimesAllBlue(flips []int) int {
	count := 0
	maxFlip := 0

	for i, f := range flips {
		if f > maxFlip {
			maxFlip = f
		}
		if maxFlip == i+1 {
			count++
		}
	}

	return count
}
```
