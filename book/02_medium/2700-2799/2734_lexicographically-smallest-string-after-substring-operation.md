# 2734 — Lexicographically Smallest String After Substring Operation

## Deskripsi

**Soal:** [2734. Lexicographically Smallest String After Substring Operation](https://leetcode.com/problems/lexicographically-smallest-string-after-substring-operation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func LexicographicallySmallestStringAfterSubstringOperation(s string) string`

## Solusi Go

```go
package main

// LeetCode #2734: Lexicographically Smallest String After Substring Operation
// https://leetcode.com/problems/lexicographically-smallest-string-after-substring-operation/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func LexicographicallySmallestStringAfterSubstringOperation(s string) string {
	n := len(s)
	b := []byte(s)

	start := -1
	for i := 0; i < n; i++ {
		if b[i] > 'a' {
			start = i
			break
		}
	}

	if start == -1 {
		b[n-1] = 'z'
		return string(b)
	}

	for i := start; i < n; i++ {
		if b[i] == 'a' {
			break
		}
		b[i]--
	}

	return string(b)
}

func main() {
	fmt.Println(LexicographicallySmallestStringAfterSubstringOperation("cbabc"))
	fmt.Println(LexicographicallySmallestStringAfterSubstringOperation("acbbc"))
	fmt.Println(LexicographicallySmallestStringAfterSubstringOperation("a"))
}
```
