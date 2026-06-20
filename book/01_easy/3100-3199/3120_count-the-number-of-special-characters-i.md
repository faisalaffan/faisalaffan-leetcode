# 3120 — Count The Number Of Special Characters I

## Deskripsi

**Soal:** [3120. Count The Number Of Special Characters I](https://leetcode.com/problems/count-the-number-of-special-characters-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3120: Count the Number of Special Characters I
// https://leetcode.com/problems/count-the-number-of-special-characters-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: numberOfSpecialChars
	fmt.Println(CountTheNumberOfSpecialCharactersI("aaAbcBC")) // 3
	fmt.Println(CountTheNumberOfSpecialCharactersI("abcd"))    // 0
	fmt.Println(CountTheNumberOfSpecialCharactersI("abAB"))   // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: numberOfSpecialChars
func CountTheNumberOfSpecialCharactersI(word string) int {
  // Membuat map untuk pencarian O(1): key → value
	lower := make(map[byte]bool)
  // Membuat map untuk pencarian O(1): key → value
	upper := make(map[byte]bool)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= 'a' && c <= 'z' {
			lower[c] = true
		} else if c >= 'A' && c <= 'Z' {
			upper[c] = true
		}
	}
	count := 0
	for c := byte('a'); c <= 'z'; c++ {
		if lower[c] && upper[c-'a'+'A'] {
			count++
		}
	}
	return count
}
```
