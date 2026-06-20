# 1109 — Corporate Flight Bookings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func corpFlightBookings(bookings [][]int, n int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n + len(bookings))  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1109: Corporate Flight Bookings
// https://leetcode.com/problems/corporate-flight-bookings/
// Difficulty: Medium
//
// Approach: Difference array (prefix sum)
// Time: O(n + len(bookings))
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5)) // [10,55,45,25,25]
	fmt.Println(corpFlightBookings([][]int{{1, 1, 5}, {2, 2, 10}}, 2))             // [5,10]
}

func corpFlightBookings(bookings [][]int, n int) []int {
  // Alokasi slice
	result := make([]int, n+2)
	for _, b := range bookings {
		first, last, seats := b[0], b[1], b[2]
		result[first] += seats
		result[last+1] -= seats
	}

	for i := 1; i <= n; i++ {
		result[i] += result[i-1]
	}

	return result[1 : n+1]
}
```
