# 0957 — Prison Cells After N Days

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func prisonAfterNDays(cells []int, n int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #957: Prison Cells After N Days
// https://leetcode.com/problems/prison-cells-after-n-days/
// Difficulty: Medium

import "fmt"

// Time: O(1) | Space: O(1)
func prisonAfterNDays(cells []int, n int) []int {
  // HashMap: O(1) lookup
	seen := make(map[[8]int]int)
	cycle := false

	for n > 0 {
		key := toArray(cells)
		if day, ok := seen[key]; ok && !cycle {
			n %= day - n
			cycle = true
		}
		seen[key] = n

		if n > 0 {
			n--
			cells = nextDay(cells)
		}
	}

	return cells
}

func toArray(cells []int) [8]int {
	return [8]int{cells[0], cells[1], cells[2], cells[3], cells[4], cells[5], cells[6], cells[7]}
}

func nextDay(cells []int) []int {
  // Alokasi slice
	next := make([]int, 8)
	for i := 1; i < 7; i++ {
		if cells[i-1] == cells[i+1] {
			next[i] = 1
		}
	}
	return next
}

func main() {
	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 7))
	fmt.Println(prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000))
}
```
