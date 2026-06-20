# 1921 — Eliminate Maximum Number Of Monsters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func EliminateMaximum(dist []int, speed []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1921: Eliminate Maximum Number of Monsters
// https://leetcode.com/problems/eliminate-maximum-number-of-monsters/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(EliminateMaximum([]int{1, 3, 4}, []int{1, 1, 1}))
	fmt.Println(EliminateMaximum([]int{1, 1, 2, 3}, []int{1, 1, 1, 1}))
	fmt.Println(EliminateMaximum([]int{3, 2, 4}, []int{5, 3, 2}))
}

// Time: O(n log n), Space: O(n)
func EliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
  // Alokasi slice integer
	time := make([]int, n)
	for i := 0; i < n; i++ {
		time[i] = (dist[i] + speed[i] - 1) / speed[i] // ceil division
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(time)

	for i := 0; i < n; i++ {
		if time[i] <= i {
			return i
		}
	}
	return n
}
```
