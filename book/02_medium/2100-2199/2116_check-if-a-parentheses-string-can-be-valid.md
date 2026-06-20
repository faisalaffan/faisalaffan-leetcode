# 2116 — Check If A Parentheses String Can Be Valid

## Deskripsi

**Soal:** [2116. Check If A Parentheses String Can Be Valid](https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func canBeValid(s string, locked string) bool`

## Solusi Go

```go
package main

// LeetCode #2116: Check if a Parentheses String Can Be Valid
// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}

	// Left to right: check for excess ')'
	balance := 0
	flexible := 0
	for i := 0; i < n; i++ {
		if locked[i] == '0' {
			flexible++
		} else if s[i] == '(' {
			balance++
		} else {
			balance--
		}
		if balance+flexible < 0 {
			return false
		}
	}

	// Right to left: check for excess '('
	balance = 0
	flexible = 0
	for i := n - 1; i >= 0; i-- {
		if locked[i] == '0' {
			flexible++
		} else if s[i] == ')' {
			balance++
		} else {
			balance--
		}
		if balance+flexible < 0 {
			return false
		}
	}

	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", canBeValid("))()))", "010100"))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", canBeValid("()()", "0000"))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", canBeValid(")", "0"))
	// Expected: false (odd length)
}
```
