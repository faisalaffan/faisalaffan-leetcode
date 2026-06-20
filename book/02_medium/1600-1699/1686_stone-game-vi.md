# 1686 — Stone Game Vi

## Deskripsi

**Soal:** [1686. Stone Game Vi](https://leetcode.com/problems/stone-game-vi/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func stoneGameVI(aliceValues []int, bobValues []int) int`

## Solusi Go

```go
package main

// LeetCode #1686: Stone Game VI
// https://leetcode.com/problems/stone-game-vi/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
  // Membuat slice untuk menyimpan hasil
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{aliceValues[i] + bobValues[i], i}
	}

	// Sort by sum descending
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	aliceScore := 0
	bobScore := 0
	for i := 0; i < n; i++ {
		idx := pairs[i][1]
		if i%2 == 0 {
			aliceScore += aliceValues[idx]
		} else {
			bobScore += bobValues[idx]
		}
	}

	if aliceScore > bobScore {
		return 1
	} else if bobScore > aliceScore {
		return -1
	}
	return 0
}

func main() {
	fmt.Println(stoneGameVI([]int{1, 3}, []int{2, 1}))      // Expected: 1 (Alice wins)
	fmt.Println(stoneGameVI([]int{1, 2}, []int{3, 1}))      // Expected: 0 (tie)
	fmt.Println(stoneGameVI([]int{2, 4, 3}, []int{1, 6, 7})) // Expected: -1 (Bob wins)
}
```
