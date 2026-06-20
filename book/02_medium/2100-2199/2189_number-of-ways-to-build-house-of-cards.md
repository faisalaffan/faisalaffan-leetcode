# 2189 — Number Of Ways To Build House Of Cards

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func houseOfCards(remainingCards int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2189: Number of Ways to Build House of Cards
// https://leetcode.com/problems/number-of-ways-to-build-house-of-cards/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n)

import "fmt"

func houseOfCards(remainingCards int) int {
	// DP: dp[c] = number of ways to use c cards
  // Alokasi slice
	dp := make([]int, remainingCards+1)
	dp[0] = 1

	// Each row (level) uses 3k - 1 cards where k >= 2
	for k := 2; 3*k-1 <= remainingCards; k++ {
		cards := 3*k - 1
		for c := remainingCards; c >= cards; c-- {
			dp[c] += dp[c-cards]
		}
	}
	return dp[remainingCards]
}

func main() {
	// Test case 1
	fmt.Println(houseOfCards(16))
	// Expected: 2

	// Test case 2
	fmt.Println(houseOfCards(2))
	// Expected: 0

	// Test case 3
	fmt.Println(houseOfCards(6))
	// Expected: 1
}
```
