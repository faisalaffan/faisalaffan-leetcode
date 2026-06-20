# 3146 — Permutation Difference Between Two Strings

## Deskripsi

**Soal:** [3146. Permutation Difference Between Two Strings](https://leetcode.com/problems/permutation-difference-between-two-strings/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #3146: Permutation Difference between Two Strings
// https://leetcode.com/problems/permutation-difference-between-two-strings/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findPermutationDifference
	fmt.Println(PermutationDifferenceBetweenTwoStrings("abc", "bac")) // 2
	fmt.Println(PermutationDifferenceBetweenTwoStrings("abcde", "edcba")) // 12
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: findPermutationDifference
func PermutationDifferenceBetweenTwoStrings(s string, t string) int {
  // Membuat map untuk pencarian O(1): key → value
	pos := make(map[byte]int)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(t); i++ {
		pos[t[i]] = i
	}
	diff := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		d := i - pos[s[i]]
		if d < 0 {
			d = -d
		}
		diff += d
	}
	return diff
}
```
