# 3438 — Find Valid Pair Of Adjacent Digits In String

## Deskripsi

**Soal:** [3438. Find Valid Pair Of Adjacent Digits In String](https://leetcode.com/problems/find-valid-pair-of-adjacent-digits-in-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #3438: Find Valid Pair of Adjacent Digits in String
// https://leetcode.com/problems/find-valid-pair-of-adjacent-digits-in-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindValidPairOfAdjacentDigitsInString("2523533"))
	fmt.Println(FindValidPairOfAdjacentDigitsInString("111"))
}

// FindValidPairOfAdjacentDigitsInString finds the first pair of adjacent equal digits where the digit's frequency > the digit.
// Time: O(n). Space: O(n).
func FindValidPairOfAdjacentDigitsInString(s string) string {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[byte]int)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cnt := freq[s[i]]
			if cnt > int(s[i]-'0') {
				return s[i : i+2]
			}
		}
	}
	return ""
}
```
