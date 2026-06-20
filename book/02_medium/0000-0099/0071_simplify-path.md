# 0071 — Simplify Path

## Deskripsi

**Soal:** [0071. Simplify Path](https://leetcode.com/problems/simplify-path/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func simplifyPath(path string) string`

## Solusi Go

```go
package main

// LeetCode #71: Simplify Path
// https://leetcode.com/problems/simplify-path/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stack := []string{}

	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
}

func main() {
	// Test case 1
	fmt.Println(simplifyPath("/home/")) // "/home"

	// Test case 2
	fmt.Println(simplifyPath("/home//foo/")) // "/home/foo"

	// Test case 3
	fmt.Println(simplifyPath("/../")) // "/"

	// Test case 4
	fmt.Println(simplifyPath("/a/./b/../../c/")) // "/c"
}

// Time: O(n) | Space: O(n)
```
