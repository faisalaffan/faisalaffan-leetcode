# 1124 — Longest Well Performing Interval

## Deskripsi

**Soal:** [1124. Longest Well Performing Interval](https://leetcode.com/problems/longest-well-performing-interval/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Prefix Sum (jumlah kumulatif)

> **Ide Kunci:** Prefix sum. Map first occurrence of each prefix sum.

## Solusi Go

```go
package main

// LeetCode #1124: Longest Well-Performing Interval
// https://leetcode.com/problems/longest-well-performing-interval/
// Difficulty: Medium
//
// Approach: Prefix sum. Map first occurrence of each prefix sum.
//           hours[i] > 8 -> +1, else -1. Find longest subarray with sum > 0.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9})) // 3
	fmt.Println(longestWPI([]int{6, 6, 6}))              // 0
}

func longestWPI(hours []int) int {
	prefix := 0
  // Membuat map untuk pencarian O(1): key → value
	firstSeen := make(map[int]int)
	result := 0

	for i, h := range hours {
		if h > 8 {
			prefix++
		} else {
			prefix--
		}

		if prefix > 0 {
			result = i + 1
		} else {
			if _, ok := firstSeen[prefix]; !ok {
				firstSeen[prefix] = i
			}
			if j, ok := firstSeen[prefix-1]; ok {
				if i-j > result {
					result = i - j
				}
			}
		}
	}

	return result
}
```
