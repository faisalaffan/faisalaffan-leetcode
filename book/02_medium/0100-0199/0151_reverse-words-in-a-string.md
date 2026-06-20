# 0151 — Reverse Words In A String

## Deskripsi

**Soal:** [0151. Reverse Words In A String](https://leetcode.com/problems/reverse-words-in-a-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func reverseWords(s string) string`

## Solusi Go

```go
package main

// LeetCode #151: Reverse Words in a String
// https://leetcode.com/problems/reverse-words-in-a-string/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	fields := strings.Fields(s)
	for i, j := 0, len(fields)-1; i < j; i, j = i+1, j-1 {
		fields[i], fields[j] = fields[j], fields[i]
	}
	return strings.Join(fields, " ")
}

func main() {
	// Test case 1
	fmt.Println(reverseWords("the sky is blue")) // "blue is sky the"

	// Test case 2
	fmt.Println(reverseWords("  hello world  ")) // "world hello"

	// Test case 3
	fmt.Println(reverseWords("a good   example")) // "example good a"
}

// Time: O(n) | Space: O(n)
```
