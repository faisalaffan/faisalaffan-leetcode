# 0009 — Palindrome Number

## Deskripsi

**Soal:** [0009. Palindrome Number](https://leetcode.com/problems/palindrome-number/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func IsPalindrome(x int) bool`

## Solusi Go

```go
package main

// LeetCode #9: Palindrome Number
// https://leetcode.com/problems/palindrome-number/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}
	return x == reverted || x == reverted/10
}

func main() {
	fmt.Println(IsPalindrome(121))  // true
	fmt.Println(IsPalindrome(-121)) // false
	fmt.Println(IsPalindrome(10))   // false
}
```
