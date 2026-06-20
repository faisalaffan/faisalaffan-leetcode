# 2744 — Find Maximum Number Of String Pairs

## Deskripsi

**Soal:** [2744. Find Maximum Number Of String Pairs](https://leetcode.com/problems/find-maximum-number-of-string-pairs/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2744: Find Maximum Number of String Pairs
// https://leetcode.com/problems/find-maximum-number-of-string-pairs/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindMaximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"}))
	fmt.Println(FindMaximumNumberOfStringPairs([]string{"ab", "ba", "cc"}))
}

func FindMaximumNumberOfStringPairs(words []string) int {
	seen := map[string]bool{}
	count := 0
	for _, w := range words {
		rev := reverse(w)
		if seen[rev] {
			count++
		}
		seen[w] = true
	}
	return count
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
```
