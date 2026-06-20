# 0873 — Length Of Longest Fibonacci Subsequence

## Deskripsi

**Soal:** [0873. Length Of Longest Fibonacci Subsequence](https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #873: Length of Longest Fibonacci Subsequence
// https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LengthOfLongestFibonacciSubsequence([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(LengthOfLongestFibonacciSubsequence([]int{1, 3, 7, 11, 12, 14, 18}))
	fmt.Println(LengthOfLongestFibonacciSubsequence([]int{1, 3, 5}))
}

// Time: O(n^2) | Space: O(n^2)
func LengthOfLongestFibonacciSubsequence(arr []int) int {
	n := len(arr)
  // Membuat map untuk pencarian O(1): key → value
	index := make(map[int]int, n)
	for i, v := range arr {
		index[v] = i
	}

  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, n)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n)
	}

	ans := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			prev := arr[i] - arr[j]
			if k, ok := index[prev]; ok && k < j {
				dp[i][j] = dp[j][k] + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			} else {
				dp[i][j] = 2
			}
		}
	}

	if ans >= 3 {
		return ans
	}
	return 0
}
```
