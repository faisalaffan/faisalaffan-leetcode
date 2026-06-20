# 0664 — Strange Printer

## Deskripsi

**Soal:** [0664. Strange Printer](https://leetcode.com/problems/strange-printer/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #664: Strange Printer
// https://leetcode.com/problems/strange-printer/
// Difficulty: Hard

import "fmt"

func main() {
	// Test cases
	testCases := []struct {
		s    string
		want int
	}{
		{"aaabbb", 2},
		{"aba", 2},
		{"", 0},
		{"a", 1},
		{"aa", 1},
		{"ab", 2},
		{"abcabc", 5},
		{"aaabbaaa", 2},
		{"ababab", 4},
		{"leetcode", 6},
		{"tbgtgb", 4},
		{"aaaaaaaaaaaaaaaaaaaa", 1},
	}

	for _, tc := range testCases {
		got := strangePrinter(tc.s)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%s: strangePrinter(%q) = %d (want %d)\n", status, tc.s, got, tc.want)
	}
}

func strangePrinter(s string) int {
	n := len(s)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// dp[i][j] = minimum turns to print s[i..j]
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, n)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1 // single character needs 1 turn
	}

	// Process intervals by length
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1

			// Worst case: print last char separately
			dp[i][j] = dp[i][j-1] + 1

			// Try to merge with a matching character
			for k := i; k < j; k++ {
				if s[k] == s[j] {
					cost := dp[i][k] + dp[k+1][j-1]
					if cost < dp[i][j] {
						dp[i][j] = cost
					}
				}
			}
		}
	}

	return dp[0][n-1]
}
```
