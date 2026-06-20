# 1066 — Campus Bikes Ii

## Deskripsi

**Soal:** [1066. Campus Bikes Ii](https://leetcode.com/problems/campus-bikes-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(W * 2^B) where W = workers, B = bikes  
**Kompleksitas Ruang:** O(2^B)

**Algoritma:** Dynamic Programming (DP), Bitmask (representasi himpunan dengan bit)

> **Ide Kunci:** DP with bitmask (minimum assignment cost)

## Solusi Go

```go
package main

// LeetCode #1066: Campus Bikes II
// https://leetcode.com/problems/campus-bikes-ii/
// Difficulty: Medium
//
// Approach: DP with bitmask (minimum assignment cost)
// Time: O(W * 2^B) where W = workers, B = bikes
// Space: O(2^B)

import "fmt"

func main() {
	fmt.Println(assignBikesII([][]int{{0, 0}, {2, 1}}, [][]int{{1, 2}, {3, 3}})) // 6
	fmt.Println(assignBikesII([][]int{{0, 0}, {1, 1}, {2, 0}}, [][]int{{1, 0}, {2, 2}, {2, 1}})) // 4
}

func assignBikesII(workers [][]int, bikes [][]int) int {
	w, b := len(workers), len(bikes)
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, 1<<b)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = -1
	}

	var dfs func(workerIdx int, mask int) int
	dfs = func(workerIdx int, mask int) int {
		if workerIdx == w {
			return 0
		}
		if dp[mask] != -1 {
			return dp[mask]
		}

		minDist := 1<<31 - 1
		for j := 0; j < b; j++ {
			if mask&(1<<j) == 0 {
				dist := abs(workers[workerIdx][0]-bikes[j][0]) + abs(workers[workerIdx][1]-bikes[j][1])
				total := dist + dfs(workerIdx+1, mask|(1<<j))
				if total < minDist {
					minDist = total
				}
			}
		}
		dp[mask] = minDist
		return minDist
	}

	return dfs(0, 0)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
