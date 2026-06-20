# 1823 — Find The Winner Of The Circular Game

## Deskripsi

**Soal:** [1823. Find The Winner Of The Circular Game](https://leetcode.com/problems/find-the-winner-of-the-circular-game/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func findTheWinner(n int, k int) int`

## Solusi Go

```go
package main

// LeetCode #1823: Find the Winner of the Circular Game
// https://leetcode.com/problems/find-the-winner-of-the-circular-game/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func findTheWinner(n int, k int) int {
	winner := 0 // 0-indexed position for 1 person
	for i := 2; i <= n; i++ {
		winner = (winner + k) % i
	}
	return winner + 1 // convert to 1-indexed
}

func main() {
	fmt.Println(findTheWinner(5, 2)) // Expected: 3
	fmt.Println(findTheWinner(6, 5)) // Expected: 1
	fmt.Println(findTheWinner(1, 1)) // Expected: 1
}
```
