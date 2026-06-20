# 1823 — Find The Winner Of The Circular Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func findTheWinner(n int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1823: Find the Winner of the Circular Game
// https://leetcode.com/problems/find-the-winner-of-the-circular-game/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func findTheWinner(n int, k int) int {
	winner := 0 // 0-indexed position for 1 person
	for i := 2; i <= n; i++ {
		winner = (winner + k) % i
	}
	return winner + 1 // convert to 1-indexed
}

func main() {
	fmt.Println(findTheWinner(5, 2)) // Expected: 3
	fmt.Println(findTheWinner(6, 5)) // Expected: 1
	fmt.Println(findTheWinner(1, 1)) // Expected: 1
}
```
