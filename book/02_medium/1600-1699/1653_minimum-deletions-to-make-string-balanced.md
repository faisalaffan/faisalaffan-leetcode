# 1653 — Minimum Deletions To Make String Balanced

## Deskripsi

**Soal:** [1653. Minimum Deletions To Make String Balanced](https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1653: Minimum Deletions to Make String Balanced
// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimumDeletions("aababbab"))
	fmt.Println(MinimumDeletions("bbaaaaabb"))
	fmt.Println(MinimumDeletions("a"))
}

func MinimumDeletions(s string) int {
	// Time: O(N), Space: O(1)
	// Count 'a's on the right
	aCount := 0
	for _, ch := range s {
		if ch == 'a' {
			aCount++
		}
	}

	bCount := 0
	minDeletions := len(s)

	for _, ch := range s {
		if ch == 'a' {
			aCount--
		}

		// Deletions needed: remove all 'b's before this point + remove all 'a's after
		deletions := bCount + aCount
		if deletions < minDeletions {
			minDeletions = deletions
		}

		if ch == 'b' {
			bCount++
		}
	}

	return minDeletions
}
```
