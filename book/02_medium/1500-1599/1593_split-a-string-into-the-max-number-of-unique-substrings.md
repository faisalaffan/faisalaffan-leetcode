# 1593 — Split A String Into The Max Number Of Unique Substrings

## Deskripsi

**Soal:** [1593. Split A String Into The Max Number Of Unique Substrings](https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(2^N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1593: Split a String Into the Max Number of Unique Substrings
// https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxUniqueSplit("ababccc"))
	fmt.Println(MaxUniqueSplit("aba"))
	fmt.Println(MaxUniqueSplit("aa"))
}

func MaxUniqueSplit(s string) int {
	// Time: O(2^N), Space: O(N)
  // Membuat map untuk pencarian O(1): key → value
	used := make(map[string]bool)
	maxCount := 0

	var backtrack func(start int, count int)
	backtrack = func(start int, count int) {
		if start == len(s) {
			if count > maxCount {
				maxCount = count
			}
			return
		}

		// Pruning: remaining chars <= max possible new substrings
		if count+(len(s)-start) <= maxCount {
			return
		}

		for end := start + 1; end <= len(s); end++ {
			sub := s[start:end]
			if !used[sub] {
				used[sub] = true
				backtrack(end, count+1)
				used[sub] = false
			}
		}
	}

	backtrack(0, 0)
	return maxCount
}
```
