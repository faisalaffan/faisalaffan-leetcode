# 2311 — Longest Binary Subsequence Less Than Or Equal To K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestSubsequence(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2311: Longest Binary Subsequence Less Than or Equal to K
// https://leetcode.com/problems/longest-binary-subsequence-less-than-or-equal-to-k/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func longestSubsequence(s string, k int) int {
	val := 0
	count := 0
	pow := 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			count++
		} else if s[i] == '1' {
			if val+pow <= k {
				val += pow
				count++
			}
		}
		if pow <= k {
			pow <<= 1
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(longestSubsequence("1001010", 5))
	// Expected: 5

	// Test case 2
	fmt.Println(longestSubsequence("0010101010110101001001011010100010101010101010111111", 93))
	// Expected: 44
}
```
