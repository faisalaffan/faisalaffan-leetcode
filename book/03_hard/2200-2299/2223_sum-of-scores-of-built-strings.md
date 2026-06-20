# 2223 — Sum Of Scores Of Built Strings

## Deskripsi

**Soal:** [2223. Sum Of Scores Of Built Strings](https://leetcode.com/problems/sum-of-scores-of-built-strings/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2223: Sum of Scores of Built Strings
// https://leetcode.com/problems/sum-of-scores-of-built-strings/
// Difficulty: Hard
//
// Z-algorithm: compute Z-array where Z[i] = longest common prefix of s and s[i:].
// Answer = sum of all Z-values + n (the whole string is its own prefix).

import (
	"fmt"
)

func main() {
	// "babab" => 9
	fmt.Println(sumScores("babab"))
	// "ababa" => 9
	fmt.Println(sumScores("ababa"))
	// "a" => 1
	fmt.Println(sumScores("a"))
	// "aa" => 3
	fmt.Println(sumScores("aa"))
	// "abc" => 3
	fmt.Println(sumScores("abc"))
}

func sumScores(s string) int64 {
	n := len(s)
  // Membuat slice untuk menyimpan hasil
	z := make([]int, n)
	l, r := 0, 0

	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}

	var sum int64 = int64(n)
	for _, v := range z {
		sum += int64(v)
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
