# 0594 — Longest Harmonious Subsequence

## Deskripsi

**Soal:** [0594. Longest Harmonious Subsequence](https://leetcode.com/problems/longest-harmonious-subsequence/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func LongestHarmoniousSubsequence(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #594: Longest Harmonious Subsequence
// https://leetcode.com/problems/longest-harmonious-subsequence/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func LongestHarmoniousSubsequence(nums []int) int {
  // Membuat map untuk pencarian O(1): key → value
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	maxLen := 0
	for v, c := range count {
		if c2, ok := count[v+1]; ok {
			if c+c2 > maxLen {
				maxLen = c + c2
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 2, 3, 4}))
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 1, 1, 1}))
}
```
