# 0463 — Island Perimeter

## Deskripsi

**Soal:** [0463. Island Perimeter](https://leetcode.com/problems/island-perimeter/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(m*n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func IslandPerimeter(grid [][]int) int`

## Solusi Go

```go
package main

// LeetCode #463: Island Perimeter
// https://leetcode.com/problems/island-perimeter/
// Difficulty: Easy

import "fmt"

// Time: O(m*n), Space: O(1)
func IslandPerimeter(grid [][]int) int {
	perimeter := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				perimeter += 4
				if i > 0 && grid[i-1][j] == 1 {
					perimeter -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					perimeter -= 2
				}
			}
		}
	}
	return perimeter
}

func main() {
	fmt.Println(IslandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
	fmt.Println(IslandPerimeter([][]int{{1}}))
}
```
