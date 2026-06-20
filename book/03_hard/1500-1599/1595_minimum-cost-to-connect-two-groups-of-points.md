# 1595 — Minimum Cost To Connect Two Groups Of Points

## Deskripsi

**Soal:** [1595. Minimum Cost To Connect Two Groups Of Points](https://leetcode.com/problems/minimum-cost-to-connect-two-groups-of-points/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Heap (priority queue), Bitmask (representasi himpunan dengan bit)

**Fungsi Solusi:** `func connectTwoGroups(cost [][]int) int`

## Solusi Go

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #1595: Minimum Cost to Connect Two Groups of Points
// https://leetcode.com/problems/minimum-cost-to-connect-two-groups-of-points/
// Difficulty: Hard
//
// We have two groups of points (size1 = m, size2 = n). Each point in group 1
// must be connected to at least one point in group 2, and vice versa.
// Cost[i][j] = cost of connecting point i (group 1) to point j (group 2).
//
// DP with bitmask: dp[i][mask] = min cost to connect first i points of group 1
// to a subset of group 2 represented by mask, with the condition that all first
// i points of group 1 have at least one connection.
//
// Optimization: precompute minCost[j] for each group-2 point across all group-1
// points, to guarantee each group-2 point gets at least one connection.
//
// DP transition: dp[i][mask] = min over j (where mask has bit j set) of:
//   dp[i-1][mask without j] + cost[i-1][j]   (first connection for point i to j)
//   dp[i][mask without j] + cost[i-1][j]     (additional connection for point i)
//
// After processing all group 1 points, for each mask that covers all group 1
// points, add the minimum connection cost for any unconnected group 2 point.

func connectTwoGroups(cost [][]int) int {
	m := len(cost)
	n := len(cost[0])

	size := 1 << n
	INF := math.MaxInt32

	// dp[mask] = min cost after processing current prefix of group 1
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, size)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	// Precompute minCostTo[j]: cheapest connection from ANY group-1 point to group-2 point j
  // Membuat slice untuk menyimpan hasil
	minCostTo := make([]int, n)
	for j := 0; j < n; j++ {
		minVal := math.MaxInt32
		for i := 0; i < m; i++ {
			if cost[i][j] < minVal {
				minVal = cost[i][j]
			}
		}
		minCostTo[j] = minVal
	}

	// Process each point in group 1
	for i := 0; i < m; i++ {
  // Membuat slice untuk menyimpan hasil
		ndp := make([]int, size)
		for mask := range ndp {
			ndp[mask] = INF
		}

		for mask := 0; mask < size; mask++ {
			if dp[mask] == INF {
				continue
			}
			// Try each group-2 point j
			for j := 0; j < n; j++ {
				newMask := mask | (1 << j)
				val := dp[mask] + cost[i][j]
				if val < ndp[newMask] {
					ndp[newMask] = val
				}
			}
			// Also option: connect i to j where j is already connected (stays same mask)
			// This is covered by the above loop (newMask may equal mask if j already set)
		}
		dp = ndp
	}

	// After processing all group-1 points, ensure every group-2 point is connected.
	// For each mask, if group-2 point j is NOT connected, add its min connection cost.
	answer := INF


	// Precompute extra cost to cover missing group-2 points for each mask
  // Membuat slice untuk menyimpan hasil
	extra := make([]int, size)
	for mask := 0; mask < size; mask++ {
		sum := 0
		for j := 0; j < n; j++ {
			if mask&(1<<j) == 0 {
				sum += minCostTo[j]
			}
		}
		extra[mask] = sum
	}

	for mask := 0; mask < size; mask++ {
		if dp[mask] == INF {
			continue
		}
		total := dp[mask] + extra[mask]
		if total < answer {
			answer = total
		}
	}

	return answer
}

func main() {
	// Example 1:
	// Input: cost = [[15, 96], [36, 2]]
	// Output: 17
	// Connect 0-0 (15) and 1-1 (2). Group1 both connected, group2 both connected. Total = 17.
	fmt.Println(connectTwoGroups([][]int{{15, 96}, {36, 2}}))

	// Example 2:
	// Input: cost = [[1, 3, 5], [4, 1, 1], [1, 5, 3]]
	// Output: 4
	// Connect 0-0 (1), 1-1 (1), 2-0 (1), 2-2 (1) -> total 4. All connected.
	fmt.Println(connectTwoGroups([][]int{{1, 3, 5}, {4, 1, 1}, {1, 5, 3}}))

	// Example 3:
	// Input: cost = [[2, 5, 1], [3, 4, 7], [8, 1, 2], [6, 2, 4], [3, 8, 8]]
	// Output: 10
	fmt.Println(connectTwoGroups([][]int{
		{2, 5, 1},
		{3, 4, 7},
		{8, 1, 2},
		{6, 2, 4},
		{3, 8, 8},
	}))
}
```
