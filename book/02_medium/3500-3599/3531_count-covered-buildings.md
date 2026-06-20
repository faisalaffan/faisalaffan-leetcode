# 3531 — Count Covered Buildings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountCoveredBuildings(buildings [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3531: Count Covered Buildings
// https://leetcode.com/problems/count-covered-buildings/
// Difficulty: Medium
// Complexity: O(n log n) time, O(1) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	b := [][]int{{1, 5}, {2, 6}, {8, 10}}
	fmt.Println("Test 1:", CountCoveredBuildings(b))
	// Test case 2
	b2 := [][]int{{1, 3}, {4, 6}}
	fmt.Println("Test 2:", CountCoveredBuildings(b2))
	// Test case 3
	b3 := [][]int{{1, 10}, {2, 5}, {3, 8}}
	fmt.Println("Test 3:", CountCoveredBuildings(b3))
}

func CountCoveredBuildings(buildings [][]int) int {
	if len(buildings) == 0 {
		return 0
	}
	// Sort by start, then by end descending
  // Custom sort dengan comparator
	sort.Slice(buildings, func(i, j int) bool {
		if buildings[i][0] != buildings[j][0] {
			return buildings[i][0] < buildings[j][0]
		}
		return buildings[i][1] > buildings[j][1]
	})
	count := 0
	maxEnd := 0
	for _, b := range buildings {
		if b[1] <= maxEnd {
			count++
		} else {
			maxEnd = b[1]
		}
	}
	return count
}
```
