# 2330 — Valid Palindrome Iv

## Deskripsi

**Soal:** [2330. Valid Palindrome Iv](https://leetcode.com/problems/valid-palindrome-iv/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func makePalindrome(s string) bool`

## Solusi Go

```go
package main

// LeetCode #2330: Valid Palindrome IV
// https://leetcode.com/problems/valid-palindrome-iv/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func makePalindrome(s string) bool {
	diff := 0
	left, right := 0, len(s)-1

  // Loop two-pointer: kiri vs kanan
	for left < right {
		if s[left] != s[right] {
			diff++
			if diff > 2 {
				return false
			}
		}
		left++
		right--
	}
	return diff <= 2
}

func main() {
	// Test case 1
	fmt.Println(makePalindrome("abcdba"))
	// Expected: true

	// Test case 2
	fmt.Println(makePalindrome("abccba"))
	// Expected: false (already palindrome, need exactly 2 changes)

	// Test case 3
	fmt.Println(makePalindrome("abcdeba"))
	// Expected: false (need 3 changes)
}
```
