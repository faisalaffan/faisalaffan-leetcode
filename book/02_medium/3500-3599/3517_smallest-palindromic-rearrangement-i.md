# 3517 — Smallest Palindromic Rearrangement I

## Deskripsi

**Soal:** [3517. Smallest Palindromic Rearrangement I](https://leetcode.com/problems/smallest-palindromic-rearrangement-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3517: Smallest Palindromic Rearrangement I
// https://leetcode.com/problems/smallest-palindromic-rearrangement-i/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", SmallestPalindromicRearrangementI("abba"))
	// Test case 2
	fmt.Println("Test 2:", SmallestPalindromicRearrangementI("cbaabc"))
	// Test case 3
	fmt.Println("Test 3:", SmallestPalindromicRearrangementI("a"))
}

func SmallestPalindromicRearrangementI(s string) string {
	// Count character frequencies
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	// Check if palindrome is possible
	oddCount := 0
	oddChar := byte(0)
	for i := 0; i < 26; i++ {
		if freq[i]%2 == 1 {
			oddCount++
			oddChar = byte(i + 'a')
		}
	}
	if oddCount > 1 {
		return ""
	}

	// Build first half
	var half []byte
	for i := 0; i < 26; i++ {
		for j := 0; j < freq[i]/2; j++ {
			half = append(half, byte(i+'a'))
		}
	}

	// To get smallest lexicographically, we want smallest characters first
	sort.Slice(half, func(i, j int) bool { return half[i] < half[j] })

	var result []byte
	result = append(result, half...)
	if oddCount == 1 {
		result = append(result, oddChar)
	}
	// Reverse the second half
	for i := len(half) - 1; i >= 0; i-- {
		result = append(result, half[i])
	}

	return string(result)
}
```
