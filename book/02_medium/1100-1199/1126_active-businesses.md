# 1126 — Active Businesses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func activeBusinesses(events [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1126: Active Businesses
// https://leetcode.com/problems/active-businesses/
// Difficulty: Medium
//
// Approach: Count occurrences per (event_type, occurences), find avg per event_type.
//           Filter businesses where > 1 event type has > avg occurences.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// (business_id, event_type, occurences)
	// Biz 1: reviews=10, ads=7 -> reviews(10 > avg 6.67) ✓, ads(7 > avg 5) ✓ -> 2 above, active
	// Biz 2: reviews=3, ads=2  -> reviews(3 < 6.67) ✗, ads(2 < 5) ✗
	// Biz 3: reviews=7, ads=6  -> reviews(7 > 6.67) ✓, ads(6 > 5) ✓ -> 2 above, active
	events := [][]int{
		{1, 1, 10}, {1, 2, 7},
		{2, 1, 3}, {2, 2, 2},
		{3, 1, 7}, {3, 2, 6},
	}
	fmt.Println(activeBusinesses(events)) // [1,3]
}

func activeBusinesses(events [][]int) []int {
	// events[i] = [business_id, event_type, occurences]
  // HashMap: O(1) lookup
	typeTotals := make(map[int]int)
  // HashMap: O(1) lookup
	typeCount := make(map[int]int)
  // HashMap: O(1) lookup
	bizEvents := make(map[int]map[int]int)

	for _, e := range events {
		bizID, eventType, occ := e[0], e[1], e[2]
		typeTotals[eventType] += occ
		typeCount[eventType]++
		if bizEvents[bizID] == nil {
			bizEvents[bizID] = make(map[int]int)
		}
		bizEvents[bizID][eventType] = occ
	}

	// Compute average per event type
  // HashMap: O(1) lookup
	typeAvg := make(map[int]float64)
	for t := range typeTotals {
		typeAvg[t] = float64(typeTotals[t]) / float64(typeCount[t])
	}

  // Alokasi slice
	result := make([]int, 0)
	for bizID, events := range bizEvents {
		aboveAvgCount := 0
		for eventType, occ := range events {
			if float64(occ) > typeAvg[eventType] {
				aboveAvgCount++
			}
		}
		if aboveAvgCount > 1 {
			result = append(result, bizID)
		}
	}

	return result
}
```
