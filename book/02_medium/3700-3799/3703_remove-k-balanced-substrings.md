# 3703 — Remove K Balanced Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func removeKBalancedSubstrings(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3703: Remove K-Balanced Substrings
// https://leetcode.com/problems/remove-k-balanced-substrings/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type pair struct {
	ch    byte
	count int
}

func removeKBalancedSubstrings(s string, k int) string {
	var stack []pair

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		c := s[i]
		if len(stack) > 0 && stack[len(stack)-1].ch == c {
			stack[len(stack)-1].count++
		} else {
			stack = append(stack, pair{ch: c, count: 1})
		}

		// Check for k-balanced pattern
		if c == ')' && stack[len(stack)-1].count == k {
			if len(stack) >= 2 {
				prev := &stack[len(stack)-2]
				if prev.ch == '(' && prev.count >= k {
					// Remove closing part
					stack = stack[:len(stack)-1]
					// Remove or reduce opening part
					if prev.count == k {
						stack = stack[:len(stack)-1]
					} else {
						prev.count -= k
					}
				}
			}
		}
	}

	res := make([]byte, 0, len(s))
	for _, p := range stack {
		for j := 0; j < p.count; j++ {
			res = append(res, p.ch)
		}
	}
	return string(res)
}

func main() {
	fmt.Println(removeKBalancedSubstrings("(())", 1))
	fmt.Println(removeKBalancedSubstrings("(()(", 1))
	fmt.Println(removeKBalancedSubstrings("((()))()()()", 3))
}
```
