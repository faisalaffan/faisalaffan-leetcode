# 3722 — Lexicographically Smallest String After Reverse

## Deskripsi

**Soal:** [3722. Lexicographically Smallest String After Reverse](https://leetcode.com/problems/lexicographically-smallest-string-after-reverse/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func reverseSubstring(s string, start int, end int) string`

## Solusi Go

```go
package main

// LeetCode #3722: Lexicographically Smallest String After Reverse
// https://leetcode.com/problems/lexicographically-smallest-string-after-reverse/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func reverseSubstring(s string, start int, end int) string {
	b := []byte(s)
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func lexicographicallySmallestStringAfterReverse(s string) string {
	n := len(s)
	best := s

	// Reverse first k
	for k := 1; k <= n; k++ {
		candidate := reverseSubstring(s, 0, k-1)
		if candidate < best {
			best = candidate
		}
	}

	// Reverse last k
	for k := 1; k <= n; k++ {
		candidate := reverseSubstring(s, n-k, n-1)
		if candidate < best {
			best = candidate
		}
	}

	return best
}

func main() {
	fmt.Println(lexicographicallySmallestStringAfterReverse("dcab"))
	fmt.Println(lexicographicallySmallestStringAfterReverse("abba"))
	fmt.Println(lexicographicallySmallestStringAfterReverse("zxy"))
}
```
