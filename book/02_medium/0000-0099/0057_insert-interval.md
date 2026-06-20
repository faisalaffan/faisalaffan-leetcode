# 0057 — Insert Interval

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func insert(intervals [][]int, newInterval []int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #57: Insert Interval
// https://leetcode.com/problems/insert-interval/
// Difficulty: Medium

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i, n := 0, len(intervals)

	// Add all intervals ending before new interval starts
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Merge overlapping intervals
	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	result = append(result, newInterval)

	// Add remaining intervals
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	// [[1 5] [6 9]]

	// Test case 2
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	// [[1 2] [3 10] [12 16]]

	// Test case 3
	fmt.Println(insert([][]int{}, []int{5, 7})) // [[5 7]]
}

// Time: O(n) | Space: O(n)
```
