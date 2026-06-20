# 3932 — Count K Th Roots In A Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func powWithLimit(base int, k int, limit int) int
```

> **💡 Hint:** Find smallest x s.t. x^k >= l, largest x s.t. x^k <= r.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search

**Kompleksitas Waktu:** O(log r)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3932: Count K-th Roots in a Range
// https://leetcode.com/problems/count-k-th-roots-in-a-range/
// Difficulty: Medium
// Time: O(log r) | Space: O(1)
// Approach: Find smallest x s.t. x^k >= l, largest x s.t. x^k <= r.
// Count = hi - lo + 1. Use binary search to avoid overflow.

import (
	"fmt"
)

func powWithLimit(base int, k int, limit int) int {
	// Returns base^k, capped at limit+1 to avoid overflow
	result := 1
	for i := 0; i < k; i++ {
		if result > limit/base {
			return limit + 1
		}
		result *= base
	}
	return result
}

func CountKThRootsInARange(l int, r int, k int) int {
	// Find first x s.t. x^k >= l
	hiX := 1
	for powWithLimit(hiX, k, r) <= r {
		hiX *= 2
	}

	left := 1
	right := hiX
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		if powWithLimit(mid, k, r) >= l {
			right = mid
		} else {
			left = mid + 1
		}
	}
	first := left

	if powWithLimit(first, k, r) > r {
		return 0
	}

	// Find last x s.t. x^k <= r
	left, right = first, hiX
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := (left + right + 1) / 2
		if powWithLimit(mid, k, r) <= r {
			left = mid
		} else {
			right = mid - 1
		}
	}
	last := left

	return last - first + 1
}

func main() {
	// Example 1
	fmt.Println(CountKThRootsInARange(1, 9, 3)) // Expected: 2

	// Example 2
	fmt.Println(CountKThRootsInARange(8, 30, 2)) // Expected: 3
}
```
