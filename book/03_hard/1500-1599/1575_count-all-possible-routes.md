# 1575 — Count All Possible Routes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countRoutes(locations []int, start int, finish int, fuel int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1575: Count All Possible Routes
// https://leetcode.com/problems/count-all-possible-routes/
// Difficulty: Hard
//
// DP[fuel][city] approach:
// - dp[f][i] = number of ways to reach city i with exactly f fuel remaining.
// - Base: dp[fuel][start] = 1 (we start there with full fuel).
// - Transition: for each city i, for each city j != i, if fuel >= |loc[i]-loc[j]|,
//   dp[f-fuelUsed][j] += dp[f][i].
// - Result: sum of dp[any fuel][finish].
//
// Since fuel up to 200 and cities up to 100, O(fuel * n^2) works.

import "fmt"

func main() {
	// Example: locations=[2,3,6,8,4], start=1, finish=3, fuel=5 -> 4
	fmt.Println(countRoutes([]int{2, 3, 6, 8, 4}, 1, 3, 5))

	// Additional tests
	fmt.Println(countRoutes([]int{1, 2, 3}, 0, 2, 3))
	fmt.Println(countRoutes([]int{1, 2, 3}, 0, 2, 1))
	fmt.Println(countRoutes([]int{5, 2, 1}, 0, 2, 3))
	fmt.Println(countRoutes([]int{2, 3, 6, 8, 4}, 1, 3, 3))
}

const MODR = 1_000_000_007

func countRoutes(locations []int, start int, finish int, fuel int) int {
	n := len(locations)
	// dp[f][i] = number of ways to reach city i with exactly f fuel
  // Matriks 2D
	dp := make([][]int, fuel+1)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[fuel][start] = 1

	result := 0
	if start == finish {
		result = 1
	}

	for f := fuel; f >= 0; f-- {
		for i := 0; i < n; i++ {
			if dp[f][i] == 0 {
				continue
			}
			if i == finish && f != fuel {
				result = (result + dp[f][i]) % MODR
			}
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				cost := abs(locations[i] - locations[j])
				if f >= cost {
					dp[f-cost][j] = (dp[f-cost][j] + dp[f][i]) % MODR
				}
			}
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
