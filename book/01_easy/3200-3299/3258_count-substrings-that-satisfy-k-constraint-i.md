# 3258 — Count Substrings That Satisfy K Constraint I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountSubstringsThatSatisfyKConstraintI(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3258: Count Substrings That Satisfy K-Constraint I
// https://leetcode.com/problems/count-substrings-that-satisfy-k-constraint-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("10101", 1))
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("1010101", 2))
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("11111", 1))
}

// CountSubstringsThatSatisfyKConstraintI counts substrings where both number of 0s and 1s <= k.
// Time: O(n^2). Space: O(1).
func CountSubstringsThatSatisfyKConstraintI(s string, k int) int {
	n := len(s)
	count := 0
	for i := 0; i < n; i++ {
		zeros, ones := 0, 0
		for j := i; j < n; j++ {
			if s[j] == '0' {
				zeros++
			} else {
				ones++
			}
			if zeros <= k || ones <= k {
				count++
			} else {
				break
			}
		}
	}
	return count
}
```
