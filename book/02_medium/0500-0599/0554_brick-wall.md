# 0554 — Brick Wall

## Deskripsi

**Soal:** [0554. Brick Wall](https://leetcode.com/problems/brick-wall/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * m) where n = rows, m = avg bricks per row  
**Kompleksitas Ruang:** O(n * m)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #554: Brick Wall
// https://leetcode.com/problems/brick-wall/
// Difficulty: Medium
// Time: O(n * m) where n = rows, m = avg bricks per row
// Space: O(n * m)

import "fmt"

func main() {
	wall := [][]int{
		{1, 2, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{2, 4},
		{3, 1, 2},
		{1, 3, 1, 1},
	}
	fmt.Println(LeastBricks(wall))
}

func LeastBricks(wall [][]int) int {
  // Membuat map untuk pencarian O(1): key → value
	gapCount := make(map[int]int)

	for _, row := range wall {
		pos := 0
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(row)-1; i++ {
			pos += row[i]
			gapCount[pos]++
		}
	}

	maxGaps := 0
	for _, count := range gapCount {
		if count > maxGaps {
			maxGaps = count
		}
	}

	return len(wall) - maxGaps
}
```
