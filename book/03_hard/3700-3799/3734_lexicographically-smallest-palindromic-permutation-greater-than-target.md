# 3734 — Lexicographically Smallest Palindromic Permutation Greater Than Target

## Deskripsi

**Soal:** [3734. Lexicographically Smallest Palindromic Permutation Greater Than Target](https://leetcode.com/problems/lexicographically-smallest-palindromic-permutation-greater-than-target/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Generate palindrome permutations by permuting the left

## Solusi Go

```go
package main

// LeetCode #3734: Lexicographically Smallest Palindromic Permutation
// Greater Than Target
// https://leetcode.com/problems/lexicographically-smallest-palindromic-permutation-greater-than-target/
// Difficulty: Hard
//
// Find the smallest palindrome permutation of s that is strictly
// greater than target. If none exists, return "".
//
// Approach: Generate palindrome permutations by permuting the left
// half. Use next-permutation on the left half to find the next valid
// palindrome greater than target.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(smallestPalindrome("aabb", "abba"))
	// Example 2
	fmt.Println(smallestPalindrome("baba", "bbaa"))
	// Edge: no greater palindrome
	fmt.Println(smallestPalindrome("ab", "ba"))
	// Edge: single char
	fmt.Println(smallestPalindrome("a", "a"))
}

func smallestPalindrome(s string, target string) string {
	n := len(s)

	// Count character frequencies
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	// Check palindrome possible: at most one odd count
	oddCount := 0
	for i := 0; i < 26; i++ {
		if freq[i]%2 == 1 {
			oddCount++
		}
	}
	if oddCount > 1 {
		return ""
	}

	// Build the left half multiset
  // Membuat slice untuk menyimpan hasil
	leftHalf := make([]byte, 0, n/2)
	mid := byte(0)
	hasMid := n%2 == 1
	for i := 0; i < 26; i++ {
		cnt := freq[i] / 2
		for j := 0; j < cnt; j++ {
			leftHalf = append(leftHalf, byte('a'+i))
		}
		if freq[i]%2 == 1 {
			mid = byte('a' + i)
		}
	}

	// Generate palindrome from a given left half
	makePalindrome := func(left []byte) string {
  // Membuat slice untuk menyimpan hasil
		buf := make([]byte, n)
		l := len(left)
		for i := 0; i < l; i++ {
			buf[i] = left[i]
			buf[n-1-i] = left[i]
		}
		if hasMid {
			buf[l] = mid
		}
		return string(buf)
	}

	// Check if we can generate a palindrome > target by trying
	// permutations of leftHalf in lexicographic order.
	// Start from the smallest permutation (already sorted leftHalf).
	// Use iterative next-permutation to find the first one > target.

  // Membuat slice untuk menyimpan hasil
	sortedLeft := make([]byte, len(leftHalf))
	copy(sortedLeft, leftHalf)
	sort.Slice(sortedLeft, func(i, j int) bool {
		return sortedLeft[i] < sortedLeft[j]
	})

	// If smallest palindrome > target, return it
	pal := makePalindrome(sortedLeft)
	if pal > target {
		return pal
	}

	// Generate next permutations and check
	for {
		// Find next permutation of leftHalf
  // Membuat slice untuk menyimpan hasil
		nextLeft := make([]byte, len(sortedLeft))
		copy(nextLeft, sortedLeft)
		if !nextPermutation(nextLeft) {
			break
		}
		sortedLeft = nextLeft
		pal = makePalindrome(sortedLeft)
		if pal > target {
			return pal
		}
	}

	return ""
}

func nextPermutation(a []byte) bool {
	n := len(a)
	if n <= 1 {
		return false
	}

	// Find first decreasing element from right
	i := n - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false
	}

	// Find element just larger than a[i]
	j := n - 1
	for a[j] <= a[i] {
		j--
	}

	a[i], a[j] = a[j], a[i]

	// Reverse the suffix
	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}

	return true
}
```
