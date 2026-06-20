# 0651 — 4 Keys Keyboard

## Deskripsi

**Soal:** [0651. 4 Keys Keyboard](https://leetcode.com/problems/4-keys-keyboard/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #651: 4 Keys Keyboard
// https://leetcode.com/problems/4-keys-keyboard/
// Difficulty: Medium [Paid]
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxA(3))
	fmt.Println(maxA(7))
	fmt.Println(maxA(10))
}

func maxA(n int) int {
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], dp[j]*(i-j-1))
		}
	}
	return dp[n]
}
```
