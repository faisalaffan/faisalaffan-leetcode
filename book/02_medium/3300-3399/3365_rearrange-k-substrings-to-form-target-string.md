# 3365 — Rearrange K Substrings To Form Target String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isPossibleToRearrange(s string, t string, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3365: Rearrange K Substrings to Form Target String
// https://leetcode.com/problems/rearrange-k-substrings-to-form-target-string/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(isPossibleToRearrange("abcd", "cdab", 2)) // true
	fmt.Println(isPossibleToRearrange("aabb", "bbaa", 2)) // true
	fmt.Println(isPossibleToRearrange("abcd", "acbd", 2)) // false
}

func isPossibleToRearrange(s string, t string, k int) bool {
	n := len(s)
	m := n / k

  // Membuat map (HashMap) — pencarian O(1)
	cnt := make(map[string]int)
	for i := 0; i < n; i += m {
		cnt[s[i:i+m]]++
		cnt[t[i:i+m]]--
	}

	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
```
