# 1427 — Perform String Shifts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func stringShift(s string, shift [][]int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + m), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1427: Perform String Shifts
// https://leetcode.com/problems/perform-string-shifts/
// Difficulty: Easy [Paid]
//
// LeetCode submission: func stringShift(s string, shift [][]int) string

import "fmt"

func main() {
	fmt.Println(PerformStringShifts("abc", [][]int{{0, 1}, {1, 2}}))             // "cab"
	fmt.Println(PerformStringShifts("abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}})) // "efgabcd"
}

// Time: O(n + m), Space: O(n)
func PerformStringShifts(s string, shift [][]int) string {
	total := 0
	for _, sh := range shift {
		if sh[0] == 0 {
			total -= sh[1]
		} else {
			total += sh[1]
		}
	}
	n := len(s)
	total %= n
	if total < 0 {
		total += n
	}
	return s[n-total:] + s[:n-total]
}
```
