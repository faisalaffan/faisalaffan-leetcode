# 1908 — Game Of Nim

## Deskripsi

**Soal:** [1908. Game Of Nim](https://leetcode.com/problems/game-of-nim/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1908: Game of Nim
// https://leetcode.com/problems/game-of-nim/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(NimGame([]int{1, 2, 3}))
	fmt.Println(NimGame([]int{1, 1, 1}))
	fmt.Println(NimGame([]int{1, 2}))
}

// Time: O(n), Space: O(1)
func NimGame(piles []int) bool {
	xor := 0
	for _, p := range piles {
		xor ^= p
	}
	return xor != 0
}
```
