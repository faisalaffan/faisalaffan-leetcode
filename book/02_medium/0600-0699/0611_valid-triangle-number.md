# 0611 — Valid Triangle Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func TriangleNumber(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(log n) for sorting

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #611: Valid Triangle Number
// https://leetcode.com/problems/valid-triangle-number/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(log n) for sorting

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(TriangleNumber([]int{2, 2, 3, 4}))
	fmt.Println(TriangleNumber([]int{4, 2, 3, 4}))
}

func TriangleNumber(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	count := 0
	n := len(nums)

	for i := n - 1; i >= 2; i-- {
		left, right := 0, i-1
  // Two-pointer: gerakkan kiri atau kanan
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}

	return count
}
```
