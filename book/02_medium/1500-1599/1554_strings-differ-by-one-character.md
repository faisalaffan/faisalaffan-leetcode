# 1554 — Strings Differ By One Character

## Deskripsi

**Soal:** [1554. Strings Differ By One Character](https://leetcode.com/problems/strings-differ-by-one-character/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N*M^2) where N = len(dict), M = string length  
**Kompleksitas Ruang:** O(N*M)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1554: Strings Differ by One Character
// https://leetcode.com/problems/strings-differ-by-one-character/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(DifferByOne([]string{"abcd", "acbd", "aacd"}))
	fmt.Println(DifferByOne([]string{"ab", "cd", "yz"}))
	fmt.Println(DifferByOne([]string{"abcd", "cccc", "abxd", "abzd"}))
}

func DifferByOne(dict []string) bool {
	// Time: O(N*M^2) where N = len(dict), M = string length
	// Space: O(N*M)
	// Use rolling hash: for each position, check if any two strings
	// become identical when that position is skipped.

	n := len(dict)
	if n < 2 {
		return false
	}
	m := len(dict[0])

	for skipIdx := 0; skipIdx < m; skipIdx++ {
  // Membuat map untuk pencarian O(1): key → value
		seen := make(map[string]bool)
		for i := 0; i < n; i++ {
			// Create string without char at skipIdx
			key := dict[i][:skipIdx] + dict[i][skipIdx+1:]
			if seen[key] {
				return true
			}
			seen[key] = true
		}
	}

	return false
}
```
