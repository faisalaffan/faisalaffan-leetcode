# 3222 — Find The Winning Player In Coin Game

## Deskripsi

**Soal:** [3222. Find The Winning Player In Coin Game](https://leetcode.com/problems/find-the-winning-player-in-coin-game/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(min(x, y/4)). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3222: Find the Winning Player in Coin Game
// https://leetcode.com/problems/find-the-winning-player-in-coin-game/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheWinningPlayerInCoinGame(2, 7))
	fmt.Println(FindTheWinningPlayerInCoinGame(4, 11))
}

// FindTheWinningPlayerInCoinGame returns the player who wins the coin game.
// Players alternately take 1 "75" coin and 4 "10" coins (value 115).
// Alice goes first. If a player cannot take exactly 115, they lose.
// Time: O(min(x, y/4)). Space: O(1).
func FindTheWinningPlayerInCoinGame(x int, y int) string {
	turns := 0
	for x >= 1 && y >= 4 {
		x -= 1
		y -= 4
		turns++
	}
	if turns%2 == 0 {
		return "Bob"
	}
	return "Alice"
}
```
