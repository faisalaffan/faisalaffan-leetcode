# 2139 — Minimum Moves To Reach Target Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minMoves(target int, maxDoubles int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log target)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2139: Minimum Moves to Reach Target Score
// https://leetcode.com/problems/minimum-moves-to-reach-target-score/
// Difficulty: Medium
// Time: O(log target) | Space: O(1)

import "fmt"

func minMoves(target int, maxDoubles int) int {
	moves := 0
	for target > 1 {
		if maxDoubles == 0 {
			moves += target - 1
			break
		}
		if target%2 == 1 {
			target--
			moves++
		} else {
			target /= 2
			maxDoubles--
			moves++
		}
	}
	return moves
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minMoves(5, 0))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minMoves(19, 2))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", minMoves(10, 4))
	// Expected: 4
}
```
