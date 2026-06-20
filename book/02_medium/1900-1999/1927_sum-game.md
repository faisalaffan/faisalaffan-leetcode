# 1927 — Sum Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func SumGame(num string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1927: Sum Game
// https://leetcode.com/problems/sum-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SumGame("5023"))
	fmt.Println(SumGame("25??"))
	fmt.Println(SumGame("?3295???"))
}

// Time: O(n), Space: O(1)
func SumGame(num string) bool {
	n := len(num)
	leftSum, rightSum := 0, 0
	leftQ, rightQ := 0, 0

	for i := 0; i < n/2; i++ {
		if num[i] == '?' {
			leftQ++
		} else {
			leftSum += int(num[i] - '0')
		}
	}
	for i := n / 2; i < n; i++ {
		if num[i] == '?' {
			rightQ++
		} else {
			rightSum += int(num[i] - '0')
		}
	}

	// Alice wants to avoid tie, Bob wants tie
	// "?" on left side favors Alice when she puts 9, etc.
	// The optimal strategy:
	// Bob will try to minimize the difference,
	// Alice will try to maximize it.

	// If total number of ? is odd, Alice can always win
	if (leftQ+rightQ)%2 == 1 {
		return true
	}

	// Each pair of '?' on opposite sides can cancel out (one puts 9, other puts 0)
	diff := leftSum - rightSum
	diff += (leftQ - rightQ) * 9 / 2

	return diff != 0
}
```
