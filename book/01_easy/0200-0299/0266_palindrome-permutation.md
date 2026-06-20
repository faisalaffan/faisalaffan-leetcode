# 0266 — Palindrome Permutation

## Deskripsi

**Soal:** [0266. Palindrome Permutation](https://leetcode.com/problems/palindrome-permutation/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (fixed 256 chars)

**Algoritma:** —

**Fungsi Solusi:** `func CanPermutePalindrome(s string) bool`

## Solusi Go

```go
package main

// LeetCode #266: Palindrome Permutation
// https://leetcode.com/problems/palindrome-permutation/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(1) (fixed 256 chars)
func CanPermutePalindrome(s string) bool {
  // Membuat map untuk pencarian O(1): key → value
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	oddCount := 0
	for _, v := range count {
		if v%2 == 1 {
			oddCount++
		}
	}
	return oddCount <= 1
}

func main() {
	fmt.Println(CanPermutePalindrome("code"))
	fmt.Println(CanPermutePalindrome("aab"))
	fmt.Println(CanPermutePalindrome("carerac"))
}
```
