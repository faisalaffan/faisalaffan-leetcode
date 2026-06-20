# 2171 — Removing Minimum Number Of Magic Beans

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumRemoval(beans []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2171: Removing Minimum Number of Magic Beans
// https://leetcode.com/problems/removing-minimum-number-of-magic-beans/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimumRemoval(beans []int) int64 {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(beans)
	n := len(beans)
	total := int64(0)
	for _, b := range beans {
		total += int64(b)
	}

	minRemoved := total // removing all beans is worst case
	for i, b := range beans {
		// If we make all remaining bags have 'b' beans:
		// we keep (n-i) * b beans
		removed := total - int64(n-i)*int64(b)
		if removed < minRemoved {
			minRemoved = removed
		}
	}

	return minRemoved
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumRemoval([]int{4, 1, 6, 5}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minimumRemoval([]int{2, 10, 3, 2}))
	// Expected: 7
}
```
