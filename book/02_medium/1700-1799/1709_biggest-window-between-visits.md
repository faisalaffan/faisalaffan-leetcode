# 1709 — Biggest Window Between Visits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func biggestWindow(visits []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1709: Biggest Window Between Visits
// https://leetcode.com/problems/biggest-window-between-visits/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func biggestWindow(visits []int) int {
	if len(visits) == 0 {
		return 0
	}
  // Sort O(n log n)
	sort.Ints(visits)
	maxGap := 0
	for i := 1; i < len(visits); i++ {
		gap := visits[i] - visits[i-1]
		if gap > maxGap {
			maxGap = gap
		}
	}
	return maxGap
}

func main() {
	fmt.Println(biggestWindow([]int{1, 3, 7, 10})) // Expected: 4 (between 3 and 7)
	fmt.Println(biggestWindow([]int{1, 2, 3, 4}))  // Expected: 1
	fmt.Println(biggestWindow([]int{5}))            // Expected: 0
}
```
