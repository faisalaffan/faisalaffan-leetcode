# 1025 — Divisor Game

## Deskripsi

**Soal:** [1025. Divisor Game](https://leetcode.com/problems/divisor-game/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1025: Divisor Game
// https://leetcode.com/problems/divisor-game/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(divisorGame(2)) // true
	fmt.Println(divisorGame(3)) // false
	fmt.Println(divisorGame(4)) // true
}

// LeetCode submission: divisorGame
func divisorGame(n int) bool {
	return n%2 == 0
}
```
