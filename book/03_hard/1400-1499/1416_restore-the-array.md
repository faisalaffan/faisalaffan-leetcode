# 1416 — Restore The Array

## Deskripsi

**Soal:** [1416. Restore The Array](https://leetcode.com/problems/restore-the-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func numberOfArrays(s string, k int) int`

## Solusi Go

```go
package main

// LeetCode #1416: Restore The Array
// https://leetcode.com/problems/restore-the-array/
// Difficulty: Hard

import "fmt"

const mod1416 = 1_000_000_007

func numberOfArrays(s string, k int) int {
	n := len(s)
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}
		var num int
		for j := i; j < n; j++ {
			num = num*10 + int(s[j]-'0')
			if num > k {
				break
			}
			dp[i] = (dp[i] + dp[j+1]) % mod1416
		}
	}
	return dp[0]
}

func main() {
	// Example: "1317", 2000 -> 8
	fmt.Println(numberOfArrays("1317", 2000))
}
```
