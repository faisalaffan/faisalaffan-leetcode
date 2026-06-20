# 3611 — Find Overbooked Employees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func FindOverbookedEmployees(shifts [][]int, limit int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3611: Find Overbooked Employees
// https://leetcode.com/problems/find-overbooked-employees/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	shifts := [][]int{{1, 3}, {2, 5}, {1, 4}}
	limit := 2
	fmt.Println("Test 1:", FindOverbookedEmployees(shifts, limit))
	// Test case 2
	shifts2 := [][]int{{1, 2}, {3, 4}}
	limit2 := 1
	fmt.Println("Test 2:", FindOverbookedEmployees(shifts2, limit2))
	// Test case 3
	shifts3 := [][]int{{1, 5}, {2, 3}, {4, 6}}
	limit3 := 2
	fmt.Println("Test 3:", FindOverbookedEmployees(shifts3, limit3))
}

func FindOverbookedEmployees(shifts [][]int, limit int) int {
	type event struct {
		time int
		typ  int // 1=start, -1=end
	}
	var events []event
	for _, s := range shifts {
		events = append(events, event{s[0], 1}, event{s[1], -1})
	}
  // Custom sort
	sort.Slice(events, func(i, j int) bool {
		if events[i].time != events[j].time {
			return events[i].time < events[j].time
		}
		return events[i].typ < events[j].typ
	})
	overbooked := 0
	count := 0
	for _, e := range events {
		count += e.typ
		if count > limit {
			overbooked++
		}
	}
	if overbooked > 0 {
		return overbooked
	}
	return 0
}
```
