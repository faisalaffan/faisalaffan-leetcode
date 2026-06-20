# 1750 — Minimum Length Of String After Deleting Similar Ends

## Deskripsi

**Soal:** [1750. Minimum Length Of String After Deleting Similar Ends](https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func minimumLength(s string) int`

## Solusi Go

```go
package main

// LeetCode #1750: Minimum Length of String After Deleting Similar Ends
// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minimumLength(s string) int {
	left, right := 0, len(s)-1

  // Loop two-pointer: kiri vs kanan
	for left < right && s[left] == s[right] {
		ch := s[left]
		// Delete from left
		for left <= right && s[left] == ch {
			left++
		}
		// Delete from right
		for left <= right && s[right] == ch {
			right--
		}
	}
	return right - left + 1
}

func main() {
	fmt.Println(minimumLength("ca"))            // Expected: 2
	fmt.Println(minimumLength("cabaabac"))      // Expected: 0
	fmt.Println(minimumLength("aabccabba"))     // Expected: 3
}
```
