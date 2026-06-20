# 1561 — Maximum Number Of Coins You Can Get

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxCoins(piles []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N log N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1561: Maximum Number of Coins You Can Get
// https://leetcode.com/problems/maximum-number-of-coins-you-can-get/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxCoins([]int{2, 4, 1, 2, 7, 8}))
	fmt.Println(MaxCoins([]int{2, 4, 5}))
	fmt.Println(MaxCoins([]int{9, 8, 7, 6, 5, 1, 2, 3, 4}))
}

func MaxCoins(piles []int) int {
	// Time: O(N log N), Space: O(1)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(piles)

	coins := 0
	n := len(piles)
	// Alice gets the largest n/3, you get the next n/3, Bob gets the smallest n/3
	// You take the second largest in each group of 3
	start := n / 3

	for i := start; i < n; i += 2 {
		coins += piles[i]
	}

	return coins
}
```
