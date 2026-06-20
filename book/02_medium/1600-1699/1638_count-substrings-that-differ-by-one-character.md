# 1638 — Count Substrings That Differ By One Character

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountSubstrings(s string, t string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N*M*min(N,M)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1638: Count Substrings That Differ by One Character
// https://leetcode.com/problems/count-substrings-that-differ-by-one-character/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountSubstrings("aba", "baba"))
	fmt.Println(CountSubstrings("ab", "bb"))
	fmt.Println(CountSubstrings("abe", "bbc"))
}

func CountSubstrings(s string, t string) int {
	// Time: O(N*M*min(N,M)), Space: O(1)
	count := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			diff := 0
			k := 0
			for i+k < len(s) && j+k < len(t) && diff <= 1 {
				if s[i+k] != t[j+k] {
					diff++
				}
				if diff == 1 {
					count++
				}
				k++
			}
		}
	}

	return count
}
```
