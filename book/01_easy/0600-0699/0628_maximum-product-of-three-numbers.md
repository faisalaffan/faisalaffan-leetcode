# 0628 — Maximum Product Of Three Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumProductOfThreeNumbers(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #628: Maximum Product of Three Numbers
// https://leetcode.com/problems/maximum-product-of-three-numbers/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(1)
func MaximumProductOfThreeNumbers(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)
	// The max product is either the three largest numbers
	// or two smallest (negative) numbers times the largest.
	p1 := nums[n-1] * nums[n-2] * nums[n-3]
	p2 := nums[0] * nums[1] * nums[n-1]
	if p1 > p2 {
		return p1
	}
	return p2
}

func main() {
	fmt.Println(MaximumProductOfThreeNumbers([]int{1, 2, 3}))
	fmt.Println(MaximumProductOfThreeNumbers([]int{1, 2, 3, 4}))
	fmt.Println(MaximumProductOfThreeNumbers([]int{-1, -2, -3}))
}
```
