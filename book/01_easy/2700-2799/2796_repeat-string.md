# 2796 — Repeat String

## Deskripsi

**Soal:** [2796. Repeat String](https://leetcode.com/problems/repeat-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2796: Repeat String
// https://leetcode.com/problems/repeat-string/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: JS problem, adapted to Go. Repeats string n times.

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RepeatString("abc", 3))
	fmt.Println(RepeatString("x", 5))
}

func RepeatString(s string, n int) string {
	return strings.Repeat(s, n)
}
```
