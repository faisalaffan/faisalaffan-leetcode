# 3083 — Existence Of A Substring In A String And Its Reverse

## Deskripsi

**Soal:** [3083. Existence Of A Substring In A String And Its Reverse](https://leetcode.com/problems/existence-of-a-substring-in-a-string-and-its-reverse/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3083: Existence of a Substring in a String and Its Reverse
// https://leetcode.com/problems/existence-of-a-substring-in-a-string-and-its-reverse/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isSubstringPresent
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("leetcode")) // true
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("abcba"))   // true
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("abcd"))    // false
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: isSubstringPresent
func ExistenceOfASubstringInAStringAndItsReverse(s string) bool {
	// Build set of all substrings of length 2
  // Membuat map untuk pencarian O(1): key → value
	substrings := make(map[string]bool)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s)-1; i++ {
		substrings[s[i:i+2]] = true
	}

	// Check reverse for any of those substrings
	for i := len(s) - 1; i > 0; i-- {
		if substrings[string(s[i])+string(s[i-1])] {
			return true
		}
	}
	return false
}
```
