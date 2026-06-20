# 3813 — Vowel Consonant Score

## Deskripsi

**Soal:** [3813. Vowel Consonant Score](https://leetcode.com/problems/vowel-consonant-score/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3813: Vowel-Consonant Score
// https://leetcode.com/problems/vowel-consonant-score/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(VowelConsonantScore("cooear"))
	fmt.Println(VowelConsonantScore("axeyizou"))
	fmt.Println(VowelConsonantScore("au 123"))
}

// Time: O(n)
// Space: O(1)
func VowelConsonantScore(s string) int {
	vowels := 0
	consonants := 0
	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			vowels++
		} else if ch >= 'a' && ch <= 'z' {
			consonants++
		}
	}
	if consonants == 0 {
		return 0
	}
	return vowels / consonants
}
```
