# 0294 — Flip Game Ii

## Deskripsi

**Soal:** [0294. Flip Game Ii](https://leetcode.com/problems/flip-game-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n!!) worst case with memo, Space: O(n!)  
**Kompleksitas Ruang:** O(n!)

**Algoritma:** —

**Fungsi Solusi:** `func canWin(currentState string) bool`

## Solusi Go

```go
package main

// LeetCode #294: Flip Game II
// https://leetcode.com/problems/flip-game-ii/
// Difficulty: Medium [Paid]
// Time: O(n!!) worst case with memo, Space: O(n!)

import "fmt"

func canWin(currentState string) bool {
  // Membuat map untuk pencarian O(1): key → value
	memo := make(map[string]bool)
	return canWinHelper(currentState, memo)
}

func canWinHelper(state string, memo map[string]bool) bool {
	if res, ok := memo[state]; ok {
		return res
	}

	bytes := []byte(state)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(state)-1; i++ {
		if bytes[i] == '+' && bytes[i+1] == '+' {
			bytes[i], bytes[i+1] = '-', '-'
			opponentWins := canWinHelper(string(bytes), memo)
			bytes[i], bytes[i+1] = '+', '+'

			if !opponentWins {
				memo[state] = true
				return true
			}
		}
	}

	memo[state] = false
	return false
}

func main() {
	fmt.Println(canWin("++++"))
	fmt.Println(canWin("+"))
	fmt.Println(canWin("+++"))
}
```
