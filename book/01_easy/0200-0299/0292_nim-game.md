# 0292 — Nim Game

## Deskripsi

**Soal:** [0292. Nim Game](https://leetcode.com/problems/nim-game/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func CanWinNim(n int) bool`

## Solusi Go

```go
package main

// LeetCode #292: Nim Game
// https://leetcode.com/problems/nim-game/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func CanWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	fmt.Println(CanWinNim(4))
	fmt.Println(CanWinNim(1))
	fmt.Println(CanWinNim(2))
}
```
