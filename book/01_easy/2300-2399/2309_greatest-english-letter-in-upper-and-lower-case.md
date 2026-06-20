# 2309 — Greatest English Letter In Upper And Lower Case

## Deskripsi

**Soal:** [2309. Greatest English Letter In Upper And Lower Case](https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #2309: Greatest English Letter in Upper and Lower Case
// https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(GreatestEnglishLetterInUpperAndLowerCase("lEeTcOdE")) // "E"
	fmt.Println(GreatestEnglishLetterInUpperAndLowerCase("arRAzFif")) // "R"
	fmt.Println(GreatestEnglishLetterInUpperAndLowerCase("AbCdEfGhIjK")) // ""
}

func GreatestEnglishLetterInUpperAndLowerCase(s string) string {
	seen := [26]bool{}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			seen[s[i]-'a'] = true
		}
	}
	best := byte(0)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			idx := s[i] - 'A'
			if seen[idx] && s[i] > best {
				best = s[i]
			}
		}
	}
	if best == 0 {
		return ""
	}
	return string(best)
}
```
