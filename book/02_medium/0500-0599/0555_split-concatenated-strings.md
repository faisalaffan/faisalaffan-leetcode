# 0555 — Split Concatenated Strings

## Deskripsi

**Soal:** [0555. Split Concatenated Strings](https://leetcode.com/problems/split-concatenated-strings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * L) where L = total length of all strings  
**Kompleksitas Ruang:** O(L)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #555: Split Concatenated Strings
// https://leetcode.com/problems/split-concatenated-strings/
// Difficulty: Medium [Paid]
// Time: O(n * L) where L = total length of all strings
// Space: O(L)

import "fmt"

func main() {
	fmt.Println(SplitLoopedString([]string{"abc", "xyz"}))
	fmt.Println(SplitLoopedString([]string{"lc", "evol", "cdy"}))
}

func SplitLoopedString(strs []string) string {
	n := len(strs)
	// For each string, use the lexicographically larger of itself and its reverse
  // Membuat slice untuk menyimpan hasil
	reversed := make([]string, n)
	for i, s := range strs {
		rev := reverse(s)
		if s > rev {
			reversed[i] = s
		} else {
			reversed[i] = rev
		}
	}

	result := ""
	for i := 0; i < n; i++ {
		s := strs[i]
		rev := reverse(s)
		// Try both original and reversed
		for _, candidate := range []string{s, rev} {
			for j := 0; j <= len(candidate); j++ {
				// split candidate into candidate[:j] + candidate[j:]
				left := candidate[:j]
				right := candidate[j:]
				var sb string
				sb += right
				for k := i + 1; k < n; k++ {
					sb += reversed[k]
				}
				for k := 0; k < i; k++ {
					sb += reversed[k]
				}
				sb += left
				if sb > result {
					result = sb
				}
			}
		}
	}

	return result
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
```
