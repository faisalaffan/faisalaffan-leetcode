# 1003 — Check If Word Is Valid After Substitutions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isValid(s string) bool
```

> **💡 Hint:** Use a stack. When we see "c", check if top two are "a" and "b".

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1003: Check If Word Is Valid After Substitutions
// https://leetcode.com/problems/check-if-word-is-valid-after-substitutions/
// Difficulty: Medium
//
// Approach: Use a stack. When we see "c", check if top two are "a" and "b".
//           If so, they form "abc" and we pop them. Otherwise, push "c".
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(isValid("aabcbc"))  // true
	fmt.Println(isValid("abcabcababcc")) // true
	fmt.Println(isValid("abccba"))  // false
	fmt.Println(isValid("cababc"))  // false
}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] == 'c' {
			n := len(stack)
			if n >= 2 && stack[n-1] == 'b' && stack[n-2] == 'a' {
				stack = stack[:n-2]
				continue
			}
		}
		stack = append(stack, s[i])
	}

	return len(stack) == 0
}
```
