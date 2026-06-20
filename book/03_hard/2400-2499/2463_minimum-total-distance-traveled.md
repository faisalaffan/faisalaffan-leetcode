# 2463 — Minimum Total Distance Traveled

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumTotalDistance(robot []int, factory [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2463: Minimum Total Distance Traveled
// https://leetcode.com/problems/minimum-total-distance-traveled/
// Difficulty: Hard
//
// Dynamic Programming. Sort robots and factories by position, then expand
// factories into individual slots. DP[i][j] = min distance to repair first i
// robots using first j factory slots. For each factory slot, either skip it
// or assign the i-th robot to it (if previous robots are handled).
// Time O(N * M) | Space O(N * M) where M = total factory slots

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(minimumTotalDistance(
		[]int{0, 4, 6},
		[][]int{{2, 2}, {6, 2}}))
	// Example 2
	fmt.Println(minimumTotalDistance(
		[]int{1, -1},
		[][]int{{-2, 1}, {2, 1}}))
	// Single robot, single factory
	fmt.Println(minimumTotalDistance(
		[]int{0},
		[][]int{{10, 1}}))
}

func minimumTotalDistance(robot []int, factory [][]int) int64 {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(robot)

  // Custom sort dengan comparator
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	// Expand factories into individual positions
	var fpos []int
	for _, f := range factory {
		pos, limit := f[0], f[1]
		for k := 0; k < limit; k++ {
			fpos = append(fpos, pos)
		}
	}

	m := len(robot)
	n := len(fpos)

  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int64, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// Skip this factory slot
			dp[i][j] = dp[i][j-1]

			// Assign robot i-1 to this factory slot
			dist := abs64(int64(robot[i-1]) - int64(fpos[j-1]))
			if dp[i-1][j-1] != math.MaxInt64 {
				dp[i][j] = min64(dp[i][j], dp[i-1][j-1]+dist)
			}
		}
	}

	return dp[m][n]
}

func abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
```
