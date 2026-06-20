# 1933 — Check If String Is Decomposable Into Value Equal Substrings

## Deskripsi

**Soal:** [1933. Check If String Is Decomposable Into Value Equal Substrings](https://leetcode.com/problems/check-if-string-is-decomposable-into-value-equal-substrings/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1933: Check if String Is Decomposable Into Value-Equal Substrings
// https://leetcode.com/problems/check-if-string-is-decomposable-into-value-equal-substrings/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("000111000"))   // false
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("00011111222"))  // true
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("011100022233")) // false
}

// Time: O(n), Space: O(1)
func CheckIfStringIsDecomposableIntoValueEqualSubstrings(s string) bool {
	hasGroupOfTwo := false
	i := 0
	for i < len(s) {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		count := j - i
		if count%3 == 1 {
			return false
		}
		if count%3 == 2 {
			if hasGroupOfTwo {
				return false
			}
			hasGroupOfTwo = true
		}
		i = j
	}
	return hasGroupOfTwo
}
```
