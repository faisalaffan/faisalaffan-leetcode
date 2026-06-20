# 0389 — Find The Difference

## Deskripsi

**Soal:** [0389. Find The Difference](https://leetcode.com/problems/find-the-difference/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func FindTheDifference(s, t string) byte`

## Solusi Go

```go
package main

// LeetCode #389: Find the Difference
// https://leetcode.com/problems/find-the-difference/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FindTheDifference(s, t string) byte {
	var diff byte
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		diff ^= s[i]
	}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(t); i++ {
		diff ^= t[i]
	}
	return diff
}

func main() {
	fmt.Printf("%c\n", FindTheDifference("abcd", "abcde"))
	fmt.Printf("%c\n", FindTheDifference("", "y"))
	fmt.Printf("%c\n", FindTheDifference("a", "aa"))
}
```
