# 3856 — Trim Trailing Vowels

## Deskripsi

**Soal:** [3856. Trim Trailing Vowels](https://leetcode.com/problems/trim-trailing-vowels/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3856: Trim Trailing Vowels
// https://leetcode.com/problems/trim-trailing-vowels/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TrimTrailingVowels("idea"))
	fmt.Println(TrimTrailingVowels("day"))
	fmt.Println(TrimTrailingVowels("aeiou"))
}

// Time: O(n)
// Space: O(n)
func TrimTrailingVowels(s string) string {
	i := len(s) - 1
	for i >= 0 && (s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u') {
		i--
	}
	return s[:i+1]
}
```
