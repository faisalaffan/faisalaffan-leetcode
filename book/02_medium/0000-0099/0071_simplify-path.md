# 0071 — Simplify Path

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func simplifyPath(path string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

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
