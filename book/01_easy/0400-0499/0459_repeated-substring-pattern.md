# 0459 — Repeated Substring Pattern

## Deskripsi

**Soal:** [0459. Repeated Substring Pattern](https://leetcode.com/problems/repeated-substring-pattern/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func RepeatedSubstringPattern(s string) bool`

## Solusi Go

```go
package main

// LeetCode #459: Repeated Substring Pattern
// https://leetcode.com/problems/repeated-substring-pattern/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func RepeatedSubstringPattern(s string) bool {
	t := s + s
	return strings.Contains(t[1:len(t)-1], s)
}

func main() {
	fmt.Println(RepeatedSubstringPattern("abab"))
	fmt.Println(RepeatedSubstringPattern("aba"))
	fmt.Println(RepeatedSubstringPattern("abcabcabcabc"))
}
```
