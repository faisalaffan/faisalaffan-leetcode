# 1701 — Average Waiting Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func averageWaitingTime(customers [][]int) float64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1701: Average Waiting Time
// https://leetcode.com/problems/average-waiting-time/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func averageWaitingTime(customers [][]int) float64 {
	currentTime := 0
	totalWait := 0

	for _, c := range customers {
		arrival, prepTime := c[0], c[1]
		if currentTime < arrival {
			currentTime = arrival
		}
		currentTime += prepTime
		totalWait += currentTime - arrival
	}

	return float64(totalWait) / float64(len(customers))
}

func main() {
	fmt.Println(averageWaitingTime([][]int{{1, 2}, {2, 5}, {4, 3}})) // Expected: 5.0
	fmt.Println(averageWaitingTime([][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}})) // Expected: 3.25
	fmt.Println(averageWaitingTime([][]int{{2, 3}, {6, 3}, {7, 5}, {11, 3}, {15, 2}, {18, 1}})) // Expected: 4.16667
}
```
