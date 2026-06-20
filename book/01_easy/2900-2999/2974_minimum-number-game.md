# 2974 — Minimum Number Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumNumberGame(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2974: Minimum Number Game
// https://leetcode.com/problems/minimum-number-game/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: numberGame
	fmt.Println(MinimumNumberGame([]int{5, 4, 2, 3})) // [3, 2, 5, 4]
	fmt.Println(MinimumNumberGame([]int{2, 5}))       // [5, 2]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: numberGame
func MinimumNumberGame(nums []int) []int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
  // Alokasi slice integer
	result := make([]int, len(nums))
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i += 2 {
		result[i] = nums[i+1]
		result[i+1] = nums[i]
	}
	return result
}
```
