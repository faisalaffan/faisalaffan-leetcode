# 0877 — Stone Game

## Deskripsi

**Soal:** [0877. Stone Game](https://leetcode.com/problems/stone-game/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #877: Stone Game
// https://leetcode.com/problems/stone-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(StoneGame([]int{5, 3, 4, 5}))
	fmt.Println(StoneGame([]int{3, 7, 2, 5}))
	fmt.Println(StoneGame([]int{1, 100, 3, 2}))
}

// Time: O(1) | Space: O(1)
// Alex always wins because there are an even number of piles
// and total stones is odd (no ties), with Alex going first.
func StoneGame(piles []int) bool {
	return true
}
```
