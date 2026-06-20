# 0135 — Candy

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func candy(ratings []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #135: Candy
// https://leetcode.com/problems/candy/
// Difficulty: Hard

import (
	"fmt"
)

func candy(ratings []int) int {
	n := len(ratings)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

  // Alokasi slice
	candies := make([]int, n)
  // Range loop
	for i := range candies {
		candies[i] = 1
	}

	// Left to right
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Right to left
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	total := 0
	for _, c := range candies {
		total += c
	}
	return total
}

func main() {
	ratings := []int{1, 0, 2}
	result := candy(ratings)
	expected := 5

	fmt.Printf("candy(%v) = %d\n", ratings, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```
