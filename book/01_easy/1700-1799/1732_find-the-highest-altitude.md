# 1732 — Find The Highest Altitude

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func LargestAltitude(gain []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1732: Find the Highest Altitude
// https://leetcode.com/problems/find-the-highest-altitude/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func LargestAltitude(gain []int) int {
	maxAlt := 0
	current := 0
	for _, g := range gain {
		current += g
		if current > maxAlt {
			maxAlt = current
		}
	}
	return maxAlt
}

func main() {
	fmt.Println(LargestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(LargestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
```
