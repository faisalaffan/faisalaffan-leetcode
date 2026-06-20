# 0747 — Largest Number At Least Twice Of Others

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func dominantIndex(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #747: Largest Number At Least Twice of Others
// https://leetcode.com/problems/largest-number-at-least-twice-of-others/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))   // 1
	fmt.Println(dominantIndex([]int{1, 2, 3, 4}))    // -1
	fmt.Println(dominantIndex([]int{1}))              // 0
}

// dominantIndex returns the index of the largest element if it is at least twice as large as all others.
// Time: O(n). Space: O(1).
func dominantIndex(nums []int) int {
	maxIdx := 0
	for i, v := range nums {
		if v > nums[maxIdx] {
			maxIdx = i
		}
	}
	for i, v := range nums {
		if i != maxIdx && v*2 > nums[maxIdx] {
			return -1
		}
	}
	return maxIdx
}
```
