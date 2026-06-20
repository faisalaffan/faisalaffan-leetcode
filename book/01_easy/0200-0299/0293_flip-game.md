# 0293 — Flip Game

## Deskripsi

**Soal:** [0293. Flip Game](https://leetcode.com/problems/flip-game/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n) for output

**Algoritma:** —

**Fungsi Solusi:** `func GeneratePossibleNextMoves(currentState string) []string`

## Solusi Go

```go
package main

// LeetCode #293: Flip Game
// https://leetcode.com/problems/flip-game/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(n) for output
func GeneratePossibleNextMoves(currentState string) []string {
	var res []string
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(currentState)-1; i++ {
		if currentState[i] == '+' && currentState[i+1] == '+' {
			flipped := currentState[:i] + "--" + currentState[i+2:]
			res = append(res, flipped)
		}
	}
	return res
}

func main() {
	fmt.Println(GeneratePossibleNextMoves("++++"))
	fmt.Println(GeneratePossibleNextMoves("+"))
}
```
