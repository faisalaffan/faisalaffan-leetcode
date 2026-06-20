# 0135 — Candy

## Deskripsi

**Soal:** [0135. Candy](https://leetcode.com/problems/candy/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func candy(ratings []int) int`

## Solusi Go

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

  // Membuat slice untuk menyimpan hasil
	candies := make([]int, n)
  // Iterasi seluruh elemen
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
