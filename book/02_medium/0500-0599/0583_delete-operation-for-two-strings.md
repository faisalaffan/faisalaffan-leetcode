# 0583 — Delete Operation For Two Strings

## Deskripsi

**Soal:** [0583. Delete Operation For Two Strings](https://leetcode.com/problems/delete-operation-for-two-strings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #583: Delete Operation for Two Strings
// https://leetcode.com/problems/delete-operation-for-two-strings/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(MinDistance("sea", "eat"))
	fmt.Println(MinDistance("leetcode", "etco"))
}

func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n+1)

	for i := 1; i <= m; i++ {
		prev := 0
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = prev + 1
			} else {
				if dp[j] > dp[j-1] {
					dp[j] = dp[j]
				} else {
					dp[j] = dp[j-1]
				}
			}
			prev = temp
		}
	}

	lcs := dp[n]
	return (m - lcs) + (n - lcs)
}
```
