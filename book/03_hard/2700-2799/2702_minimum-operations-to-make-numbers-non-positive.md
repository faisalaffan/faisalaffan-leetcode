# 2702 — Minimum Operations To Make Numbers Non Positive

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumOperationsToMakeNumbersNonPositive(nums []int, x int, y int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2702: Minimum Operations to Make Numbers Non-positive
// https://leetcode.com/problems/minimum-operations-to-make-numbers-non-positive/
// Difficulty: Hard [Paid]
//
// Given nums, x, y. In one operation: choose i, decrease nums[i] by x,
// decrease all others by y. Find min ops to make all <= 0.
// Binary search on answer.

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToMakeNumbersNonPositive([]int{3, 4, 1}, 3, 1))
	fmt.Println(MinimumOperationsToMakeNumbersNonPositive([]int{5, 3}, 3, 1))
}

func MinimumOperationsToMakeNumbersNonPositive(nums []int, x int, y int) int {
	diff := x - y
	lo, hi := 0, 0
	for _, v := range nums {
		if v > hi {
			hi = v
		}
	}

	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(nums, mid, y, diff) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func check(nums []int, t, y, diff int) bool {
	var extra int64
	tt := int64(t)
	yy := int64(y)
	dd := int64(diff)
	for _, v := range nums {
		v64 := int64(v)
		if v64 > tt*yy {
			need := v64 - tt*yy
			extra += (need + dd - 1) / dd
			if extra > tt {
				return false
			}
		}
	}
	return extra <= tt
}
```
