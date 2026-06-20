# 2037 — Minimum Number Of Moves To Seat Everyone

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumNumberOfMovesToSeatEveryone(seats []int, students []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2037: Minimum Number of Moves to Seat Everyone
// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumNumberOfMovesToSeatEveryone([]int{3, 1, 5}, []int{2, 7, 4}))   // 4
	fmt.Println(MinimumNumberOfMovesToSeatEveryone([]int{4, 1, 5, 9}, []int{1, 3, 2, 6})) // 7
}

// Time: O(n log n), Space: O(1)
func MinimumNumberOfMovesToSeatEveryone(seats []int, students []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(seats)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(students)
	moves := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(seats); i++ {
		diff := seats[i] - students[i]
		if diff < 0 {
			diff = -diff
		}
		moves += diff
	}
	return moves
}
```
