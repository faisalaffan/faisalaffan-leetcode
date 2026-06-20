# 1249 — Minimum Remove To Make Valid Parentheses

## Deskripsi

**Soal:** [1249. Minimum Remove To Make Valid Parentheses](https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func minRemoveToMakeValid(s string) string`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1249: Minimum Remove to Make Valid Parentheses
// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
// Difficulty: Medium

// First pass: remove unmatched ')'. Second pass: remove unmatched '('.

// Time: O(n)
// Space: O(n)

func minRemoveToMakeValid(s string) string {
	n := len(s)
  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0)
  // Membuat slice untuk menyimpan hasil
	remove := make([]bool, n)

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				remove[i] = true
			}
		}
	}

	for _, idx := range stack {
		remove[idx] = true
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		if !remove[i] {
			result = append(result, s[i])
		}
	}
	return string(result)
}

func main() {
	fmt.Printf("%q (expected: %q)\n", minRemoveToMakeValid("lee(t(c)o)de)"), "lee(t(c)o)de")
	fmt.Printf("%q (expected: %q)\n", minRemoveToMakeValid("a)b(c)d"), "ab(c)d")
	fmt.Printf("%q (expected: %q)\n", minRemoveToMakeValid("))(("), "")
}
```
