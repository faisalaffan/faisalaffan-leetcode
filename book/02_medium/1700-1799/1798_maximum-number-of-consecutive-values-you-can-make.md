# 1798 — Maximum Number Of Consecutive Values You Can Make

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func getMaximumConsecutive(coins []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1798: Maximum Number of Consecutive Values You Can Make
// https://leetcode.com/problems/maximum-number-of-consecutive-values-you-can-make/
// Difficulty: Medium
// Time: O(n log n), Space: O(1)

import (
	"fmt"
	"sort"
)

func getMaximumConsecutive(coins []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(coins)
	maxReach := 0
	for _, c := range coins {
		if c > maxReach+1 {
			break
		}
		maxReach += c
	}
	return maxReach + 1
}

func main() {
	fmt.Println(getMaximumConsecutive([]int{1, 3}))          // Expected: 2
	fmt.Println(getMaximumConsecutive([]int{1, 1, 1, 4}))   // Expected: 8
	fmt.Println(getMaximumConsecutive([]int{1, 4, 10, 3, 1})) // Expected: 20
}
```
