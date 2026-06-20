# 1406 — Stone Game Iii

## Deskripsi

**Soal:** [1406. Stone Game Iii](https://leetcode.com/problems/stone-game-iii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func stoneGameIII(stoneValue []int) string`

## Solusi Go

```go
package main

// LeetCode #1406: Stone Game III
// https://leetcode.com/problems/stone-game-iii/
// Difficulty: Hard

import "fmt"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = stoneValue[i] - dp[i+1]
		if i+2 <= n {
			if sum := stoneValue[i] + stoneValue[i+1] - dp[i+2]; sum > dp[i] {
				dp[i] = sum
			}
		}
		if i+3 <= n {
			if sum := stoneValue[i] + stoneValue[i+1] + stoneValue[i+2] - dp[i+3]; sum > dp[i] {
				dp[i] = sum
			}
		}
	}
	if dp[0] > 0 {
		return "Alice"
	} else if dp[0] < 0 {
		return "Bob"
	}
	return "Tie"
}

func main() {
	// Example: [1,2,3,7] -> "Bob"
	fmt.Println(stoneGameIII([]int{1, 2, 3, 7}))
}
```
