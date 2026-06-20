# 1927 — Sum Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SumGame(num string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
