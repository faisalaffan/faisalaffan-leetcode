# 3124 — Find Longest Calls

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func findLongestCalls(calls [][]int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3124: Find Longest Calls
// https://leetcode.com/problems/find-longest-calls/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findLongestCalls(calls [][]int, k int) []int {
	type call struct {
		id     int
		dur    int
		caller int
	}

	var list []call
	for _, c := range calls {
		list = append(list, call{c[0], c[2] - c[1], c[3]})
	}

  // Custom sort
	sort.Slice(list, func(i, j int) bool {
		if list[i].dur != list[j].dur {
			return list[i].dur > list[j].dur
		}
		return list[i].id < list[j].id
	})

  // Alokasi slice
	ans := make([]int, 0, k)
	for i := 0; i < k && i < len(list); i++ {
		ans = append(ans, list[i].id)
	}
	return ans
}

func main() {
	fmt.Println(findLongestCalls([][]int{{1, 0, 30, 1}, {2, 5, 25, 2}, {3, 10, 20, 1}}, 2)) // Expected: [1 2]
	fmt.Println(findLongestCalls([][]int{{1, 0, 10, 1}, {2, 0, 5, 2}}, 3))                    // Expected: [1 2]
}
```
