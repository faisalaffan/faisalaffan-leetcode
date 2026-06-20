# 2271 — Maximum White Tiles Covered By A Carpet

## Deskripsi

**Soal:** [2271. Maximum White Tiles Covered By A Carpet](https://leetcode.com/problems/maximum-white-tiles-covered-by-a-carpet/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maximumWhiteTiles(tiles [][]int, carpetLen int) int`

## Solusi Go

```go
package main

// LeetCode #2271: Maximum White Tiles Covered by a Carpet
// https://leetcode.com/problems/maximum-white-tiles-covered-by-a-carpet/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})

	maxCovered := 0
	right := 0
	prefix := 0

	for left := 0; left < len(tiles); left++ {
		// Expand right pointer to cover tiles within carpet
		for right < len(tiles) && tiles[right][1] < tiles[left][0]+carpetLen {
			prefix += tiles[right][1] - tiles[right][0] + 1
			right++
		}

		covered := prefix
		if right < len(tiles) {
			partial := tiles[left][0] + carpetLen - 1 - tiles[right][0] + 1
			if partial > 0 {
				covered += partial
			}
		}

		if covered > maxCovered {
			maxCovered = covered
		}

		// Remove left tile prefix before moving
		prefix -= tiles[left][1] - tiles[left][0] + 1
	}
	return maxCovered
}

func main() {
	// Test case 1
	fmt.Println(maximumWhiteTiles([][]int{{1, 5}, {10, 11}, {12, 18}, {20, 25}, {30, 32}}, 10))
	// Expected: 9

	// Test case 2
	fmt.Println(maximumWhiteTiles([][]int{{10, 11}, {1, 1}}, 2))
	// Expected: 2
}
```
