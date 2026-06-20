# 1717 — Maximum Score From Removing Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumGain(s string, x int, y int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1717: Maximum Score From Removing Substrings
// https://leetcode.com/problems/maximum-score-from-removing-substrings/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maximumGain(s string, x int, y int) int {
	// Ensure we always process the higher-scoring pair first
	if y > x {
		s = reverse(s)
		x, y = y, x
	}

	ans := 0

	// First pass: remove "ab" for x points
	stack := make([]byte, 0, len(s))
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == 'a' && s[i] == 'b' {
			stack = stack[:len(stack)-1]
			ans += x
		} else {
			stack = append(stack, s[i])
		}
	}

	// Second pass: remove "ba" for y points from remaining
	stack2 := make([]byte, 0, len(stack))
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(stack); i++ {
		if len(stack2) > 0 && stack2[len(stack2)-1] == 'b' && stack[i] == 'a' {
			stack2 = stack2[:len(stack2)-1]
			ans += y
		} else {
			stack2 = append(stack2, stack[i])
		}
	}

	return ans
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	fmt.Println(maximumGain("cdbcbbaaabab", 4, 5)) // Expected: 19
	fmt.Println(maximumGain("aabbaaxybbaabb", 5, 4)) // Expected: 20
	fmt.Println(maximumGain("ab", 1, 2)) // Expected: 1
}
```
