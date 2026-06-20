# 0293 — Flip Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func GeneratePossibleNextMoves(currentState string) []string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n) for output


## 💻 Solusi Go

```go
package main

// LeetCode #293: Flip Game
// https://leetcode.com/problems/flip-game/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(n) for output
func GeneratePossibleNextMoves(currentState string) []string {
	var res []string
  // Linear scan O(n)
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
