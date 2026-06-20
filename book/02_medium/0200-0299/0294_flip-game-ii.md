# 0294 — Flip Game Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func canWin(currentState string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** O(n!!) worst case with memo, Space: O(n!)  |  **Ruang:** O(n!)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #294: Flip Game II
// https://leetcode.com/problems/flip-game-ii/
// Difficulty: Medium [Paid]
// Time: O(n!!) worst case with memo, Space: O(n!)

import "fmt"

func canWin(currentState string) bool {
  // HashMap: O(1) lookup
	memo := make(map[string]bool)
	return canWinHelper(currentState, memo)
}

func canWinHelper(state string, memo map[string]bool) bool {
	if res, ok := memo[state]; ok {
		return res
	}

	bytes := []byte(state)
  // Linear scan O(n)
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
