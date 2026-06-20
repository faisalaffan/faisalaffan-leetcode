# 2238 — Number Of Times A Driver Was A Passenger

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func countPassengers(rides [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2238: Number of Times a Driver Was a Passenger
// https://leetcode.com/problems/number-of-times-a-driver-was-a-passenger/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func countPassengers(rides [][]int) []int {
	// Each ride: [driver_id, passenger_id]
  // HashMap: O(1) lookup
	passengerCount := make(map[int]int)
  // HashMap: O(1) lookup
	drivers := make(map[int]bool)

	for _, ride := range rides {
		driver, passenger := ride[0], ride[1]
		drivers[driver] = true
		passengerCount[passenger]++
	}

	maxID := 0
	for id := range drivers {
		if id > maxID {
			maxID = id
		}
	}

  // Alokasi slice
	result := make([]int, maxID+1)
	for id := 1; id <= maxID; id++ {
		if drivers[id] {
			result[id] = passengerCount[id]
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(countPassengers([][]int{{1, 2}, {2, 3}, {3, 1}}))
	// Expected: [0, 1, 1, 1]

	// Test case 2
	fmt.Println(countPassengers([][]int{{1, 2}, {1, 3}, {1, 4}}))
	// Expected: [0, 0]
}
```
