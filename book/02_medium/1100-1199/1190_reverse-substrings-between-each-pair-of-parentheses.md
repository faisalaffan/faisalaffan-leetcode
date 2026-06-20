# 1190 — Reverse Substrings Between Each Pair Of Parentheses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func reverseParentheses(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1190: Reverse Substrings Between Each Pair of Parentheses
// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/
// Difficulty: Medium

// Find matching parentheses, then process from outside-in.
// Use wormhole approach: when hitting '(', jump to matching ')' and reverse direction.

// Time: O(n)
// Space: O(n)

func reverseParentheses(s string) string {
	n := len(s)
  // Alokasi slice integer
	pair := make([]int, n)
  // Alokasi slice integer
	stack := make([]int, 0)

	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else if ch == ')' {
			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[open] = i
			pair[i] = open
		}
	}

	result := make([]byte, 0, n)
	dir := 1
	for i := 0; i < n; i += dir {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			dir = -dir
		} else {
			result = append(result, s[i])
		}
	}

	return string(result)
}

func main() {
	fmt.Printf("%q (expected: %q)\n", reverseParentheses("(abcd)"), "dcba")
	fmt.Printf("%q (expected: %q)\n", reverseParentheses("(u(love)i)"), "iloveu")
	fmt.Printf("%q (expected: %q)\n", reverseParentheses("(ed(et(oc))el)"), "leetcode")
}
```
