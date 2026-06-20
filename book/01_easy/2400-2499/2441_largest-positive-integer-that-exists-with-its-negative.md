# 2441 — Largest Positive Integer That Exists With Its Negative

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func LargestPositiveIntegerThatExistsWithItsNegative(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2441: Largest Positive Integer That Exists With Its Negative
// https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(LargestPositiveIntegerThatExistsWithItsNegative([]int{-1, 2, -3, 3}))  // 3
	fmt.Println(LargestPositiveIntegerThatExistsWithItsNegative([]int{-1, 10, 6, 7, -7, 1})) // 7
	fmt.Println(LargestPositiveIntegerThatExistsWithItsNegative([]int{-10, 8, 6, 7, -2, -3})) // -1
}

func LargestPositiveIntegerThatExistsWithItsNegative(nums []int) int {
	seen := map[int]bool{}
	for _, n := range nums {
		seen[n] = true
	}
	best := -1
	for _, n := range nums {
		if n > 0 && seen[-n] && n > best {
			best = n
		}
	}
	return best
}
```
