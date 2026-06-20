# 2506 — Count Pairs Of Similar Strings

## Deskripsi

**Soal:** [2506. Count Pairs Of Similar Strings](https://leetcode.com/problems/count-pairs-of-similar-strings/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2506: Count Pairs Of Similar Strings
// https://leetcode.com/problems/count-pairs-of-similar-strings/
// Difficulty: Easy
// Time O(n * m) | Space O(n)

import "fmt"

func main() {
	fmt.Println(CountPairsOfSimilarStrings([]string{"aba", "aabb", "abcd", "bac", "aabc"})) // 2
	fmt.Println(CountPairsOfSimilarStrings([]string{"aabb", "ab", "ba"}))                    // 3
}

func charMask(s string) int {
	mask := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		mask |= 1 << (s[i] - 'a')
	}
	return mask
}

func CountPairsOfSimilarStrings(words []string) int {
	count := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(words); i++ {
		maskI := charMask(words[i])
		for j := i + 1; j < len(words); j++ {
			if maskI == charMask(words[j]) {
				count++
			}
		}
	}
	return count
}
```
