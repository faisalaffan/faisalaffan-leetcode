# 3707 — Equal Score Substrings

## Deskripsi

**Soal:** [3707. Equal Score Substrings](https://leetcode.com/problems/equal-score-substrings/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3707: Equal Score Substrings
// https://leetcode.com/problems/equal-score-substrings/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(EqualScoreSubstrings("adcb"))
	fmt.Println(EqualScoreSubstrings("bace"))
}

// Time: O(n)
// Space: O(1)
func EqualScoreSubstrings(s string) bool {
	total := 0
	for _, ch := range s {
		total += int(ch-'a') + 1
	}

	prefix := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s)-1; i++ {
		prefix += int(s[i]-'a') + 1
		if prefix == total-prefix {
			return true
		}
	}
	return false
}
```
