# 3021 — Alice And Bob Playing Flower Game

## Deskripsi

**Soal:** [3021. Alice And Bob Playing Flower Game](https://leetcode.com/problems/alice-and-bob-playing-flower-game/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3021: Alice and Bob Playing Flower Game
// https://leetcode.com/problems/alice-and-bob-playing-flower-game/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(flowerGame(3, 2))
	fmt.Println(flowerGame(1, 1))
	fmt.Println(flowerGame(4, 4))
}

func flowerGame(n int, m int) int64 {
	oddN := int64((n + 1) / 2)
	evenN := int64(n / 2)
	oddM := int64((m + 1) / 2)
	evenM := int64(m / 2)
	return oddN*evenM + evenN*oddM
}
```
