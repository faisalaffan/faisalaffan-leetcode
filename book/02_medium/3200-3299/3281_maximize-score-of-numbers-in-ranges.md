# 3281 — Maximize Score Of Numbers In Ranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxPossibleScore(start []int, d int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n + n log D) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3281: Maximize Score of Numbers in Ranges
// https://leetcode.com/problems/maximize-score-of-numbers-in-ranges/
// Difficulty: Medium
// Time: O(n log n + n log D) Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxPossibleScore([]int{2, 6, 13, 13}, 5)) // 5
	fmt.Println(maxPossibleScore([]int{1, 2, 3, 4, 5}, 3)) // 1
	fmt.Println(maxPossibleScore([]int{6, 0, 3}, 2))       // 4
}

func maxPossibleScore(start []int, d int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(start)
	n := len(start)

	lo, hi := 0, start[n-1]+d-start[0]
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if possible(start, d, mid) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

func possible(start []int, d int, score int) bool {
	prev := start[0]
	for i := 1; i < len(start); i++ {
		if prev+score > start[i]+d {
			return false
		}
		if prev+score > start[i] {
			prev = prev + score
		} else {
			prev = start[i]
		}
	}
	return true
}
```
