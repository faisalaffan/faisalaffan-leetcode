# 2428 — Maximum Sum Of An Hourglass

## Deskripsi

**Soal:** [2428. Maximum Sum Of An Hourglass](https://leetcode.com/problems/maximum-sum-of-an-hourglass/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2428: Maximum Sum of an Hourglass
// https://leetcode.com/problems/maximum-sum-of-an-hourglass/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)
// Sliding hourglass over the grid.

import "fmt"

func main() {
	fmt.Println(maxSum([][]int{{6, 2, 1, 3}, {4, 2, 1, 5}, {9, 2, 8, 7}, {4, 1, 2, 9}})) // 30
	fmt.Println(maxSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))                         // 35
}

func maxSum(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	ans := 0
	for i := 0; i+2 < r; i++ {
		for j := 0; j+2 < c; j++ {
			sum := grid[i][j] + grid[i][j+1] + grid[i][j+2] + // top row
				grid[i+1][j+1] + // middle
				grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2] // bottom row
			if sum > ans {
				ans = sum
			}
		}
	}
	return ans
}
```
