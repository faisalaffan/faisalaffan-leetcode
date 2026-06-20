# 2923 — Find Champion I

## Deskripsi

**Soal:** [2923. Find Champion I](https://leetcode.com/problems/find-champion-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2923: Find Champion I
// https://leetcode.com/problems/find-champion-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findChampion
	fmt.Println(FindChampionI([][]int{{0, 1}, {0, 0}}))          // 0
	fmt.Println(FindChampionI([][]int{{0, 0, 1}, {1, 0, 1}, {0, 0, 0}})) // 1
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: findChampion
func FindChampionI(grid [][]int) int {
	n := len(grid)
	for i := 0; i < n; i++ {
		isChampion := true
		for j := 0; j < n; j++ {
			if i != j && grid[i][j] != 1 {
				isChampion = false
				break
			}
		}
		if isChampion {
			return i
		}
	}
	return -1
}
```
