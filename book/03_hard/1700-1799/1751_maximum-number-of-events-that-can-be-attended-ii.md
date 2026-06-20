# 1751 — Maximum Number Of Events That Can Be Attended Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maxValue(events [][]int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, DP, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1751: Maximum Number of Events That Can Be Attended II
// https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/
// Difficulty: Hard
//
// Approach: Sort by end time + DP + Binary Search.
// Sort events by end time. For each event i, find the last event j
// that ends before event i starts (using binary search).
// dp[i][k] = max value using up to k events from first i events (1-indexed).
// Transition: skip event i or take event i + dp[prev][k-1].

import (
	"fmt"
	"sort"
)

func maxValue(events [][]int, k int) int {
	// Sort by end time
  // Custom sort
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	n := len(events)
	// prev[i] = index of last event that ends before events[i] starts
  // Alokasi slice
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		start := events[i][0]
		// Binary search for rightmost event with end < start
		lo, hi := 0, i-1
		prev[i] = -1
		for lo <= hi {
			mid := (lo + hi) / 2
			if events[mid][1] < start {
				prev[i] = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	// dp[i][j] = max value using first i events (0-indexed), at most j events
  // Matriks 2D
	dp := make([][]int, n+1)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		val := events[i-1][2]
		p := prev[i-1] + 1 // 1-indexed version of prev
		for j := 1; j <= k; j++ {
			// Skip event i-1
			best := dp[i-1][j]
			// Take event i-1
			take := val + dp[p][j-1]
			if take > best {
				best = take
			}
			dp[i][j] = best
		}
	}

	return dp[n][k]
}

func main() {
	// Example test case
	events := [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}
	fmt.Println("events=[[1,2,4],[3,4,3],[2,3,1]],k=2 →", maxValue(events, 2)) // Expected: 7

	// Additional tests
	events2 := [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}
	fmt.Println("events=[[1,2,4],[3,4,3],[2,3,1]],k=1 →", maxValue(events2, 1)) // Expected: 4

	events3 := [][]int{{1, 1, 5}, {2, 2, 3}, {3, 3, 4}}
	fmt.Println("events=[[1,1,5],[2,2,3],[3,3,4]],k=3 →", maxValue(events3, 3)) // Expected: 12

	events4 := [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}
	fmt.Println("events=[[1,3,2],[4,5,2],[2,4,3]],k=2 →", maxValue(events4, 2)) // Expected: 5
}
```
