# 3847 — Find The Score Difference In A Game

## Deskripsi

**Soal:** [3847. Find The Score Difference In A Game](https://leetcode.com/problems/find-the-score-difference-in-a-game/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func FindTheScoreDifferenceInAGame(nums []int) int`

> **Ide Kunci:** Track active player and swap on odd or every 6th game.

## Solusi Go

```go
package main

// LeetCode #3847: Find the Score Difference in a Game
// https://leetcode.com/problems/find-the-score-difference-in-a-game/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Track active player and swap on odd or every 6th game.

import "fmt"

func FindTheScoreDifferenceInAGame(nums []int) int {
	first, second := 0, 0
	activeIsFirst := true

	for i, v := range nums {
		// Swap if odd
		if v%2 == 1 {
			activeIsFirst = !activeIsFirst
		}
		// Swap every 6th game (0-indexed, so i%6 == 5)
		if i%6 == 5 {
			activeIsFirst = !activeIsFirst
		}
		if activeIsFirst {
			first += v
		} else {
			second += v
		}
	}

	return first - second
}

func main() {
	// Example 1
	fmt.Println(FindTheScoreDifferenceInAGame([]int{1, 2, 3})) // Expected: 0

	// Example 2
	fmt.Println(FindTheScoreDifferenceInAGame([]int{2, 4, 2, 1, 2, 1})) // Expected: 4

	// Example 3
	fmt.Println(FindTheScoreDifferenceInAGame([]int{1})) // Expected: -1
}
```
