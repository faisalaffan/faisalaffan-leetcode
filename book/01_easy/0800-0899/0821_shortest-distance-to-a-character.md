# 0821 — Shortest Distance To A Character

## Deskripsi

**Soal:** [0821. Shortest Distance To A Character](https://leetcode.com/problems/shortest-distance-to-a-character/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1) excluding output.  
**Kompleksitas Ruang:** O(1) excluding output.

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #821: Shortest Distance to a Character
// https://leetcode.com/problems/shortest-distance-to-a-character/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e')) // [3,2,1,0,1,0,0,1,2,2,1,0]
	fmt.Println(shortestToChar("aaab", 'b'))          // [3,2,1,0]
}

// shortestToChar returns the shortest distance from each character to the target character c.
// Time: O(n). Space: O(1) excluding output.
func shortestToChar(s string, c byte) []int {
	n := len(s)
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
	// Initialize with large value
  // Iterasi seluruh elemen
	for i := range result {
		result[i] = n
	}

	// Left to right
	last := -n
	for i := 0; i < n; i++ {
		if s[i] == c {
			last = i
		}
		result[i] = min(result[i], i-last)
	}

	// Right to left
	last = 2 * n
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			last = i
		}
		result[i] = min(result[i], last-i)
	}
	return result
}
```
