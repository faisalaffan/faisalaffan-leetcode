# 0461 — Hamming Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func HammingDistance(x, y int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #461: Hamming Distance
// https://leetcode.com/problems/hamming-distance/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func HammingDistance(x, y int) int {
	xor := x ^ y
	count := 0
	for xor > 0 {
		xor &= xor - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(HammingDistance(1, 4))
	fmt.Println(HammingDistance(3, 1))
}
```
