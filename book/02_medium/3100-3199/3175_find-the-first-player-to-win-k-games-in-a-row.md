# 3175 — Find The First Player To Win K Games In A Row

## Deskripsi

**Soal:** [3175. Find The First Player To Win K Games In A Row](https://leetcode.com/problems/find-the-first-player-to-win-k-games-in-a-row/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func findWinningPlayer(skills []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #3175: Find The First Player to Win K Games in a Row
// https://leetcode.com/problems/find-the-first-player-to-win-k-games-in-a-row/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func findWinningPlayer(skills []int, k int) int {
	n := len(skills)
	maxIdx := 0
	curWins := 0

	for i := 1; i < n; i++ {
		if skills[i] > skills[maxIdx] {
			maxIdx = i
			curWins = 1
		} else {
			curWins++
		}
		if curWins >= k {
			return maxIdx
		}
	}
	return maxIdx
}

func main() {
	fmt.Println(findWinningPlayer([]int{4, 2, 6, 3, 9}, 2)) // Expected: 2
	fmt.Println(findWinningPlayer([]int{2, 5, 4}, 3))        // Expected: 1
}
```
