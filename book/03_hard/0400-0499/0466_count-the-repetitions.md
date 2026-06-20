# 0466 — Count The Repetitions

## Deskripsi

**Soal:** [0466. Count The Repetitions](https://leetcode.com/problems/count-the-repetitions/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Greedy (pemilihan optimal lokal)

## Solusi Go

```go
package main

import "fmt"

// LeetCode #466: Count The Repetitions
// https://leetcode.com/problems/count-the-repetitions/
// Difficulty: Hard
//
// Find the maximum m such that [s2, m] is a subsequence of [s1, n1].
// Greedy matching through repeated s1, counting how many full s2 sequences
// are matched as a subsequence.

func main() {
	// Example: s1="acb", n1=4 => "acbacbacbacb"
	// s2="ab", n2=2 => "abab"
	// "abab" is a subsequence of "acbacbacbacb" => m=2
	fmt.Println("m:", getMaxRepetitions("acb", 4, "ab", 2)) // 2

	// Example 2
	fmt.Println("m:", getMaxRepetitions("abc", 4, "ab", 2)) // 2

	// Repeated char
	fmt.Println("m:", getMaxRepetitions("aaa", 3, "aa", 1)) // 4

	// No match
	fmt.Println("m:", getMaxRepetitions("a", 1, "b", 1)) // 0
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}

	len1, len2 := len(s1), len(s2)
	totalChars := int64(n1) * int64(len1)

	matchCount := 0
	s2idx := 0

	for i := int64(0); i < totalChars; i++ {
		if s1[i%int64(len1)] == s2[s2idx] {
			s2idx++
			if s2idx == len2 {
				matchCount++
				s2idx = 0
			}
		}
	}
	return matchCount / n2
}
```
