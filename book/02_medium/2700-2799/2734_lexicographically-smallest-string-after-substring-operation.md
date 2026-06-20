# 2734 — Lexicographically Smallest String After Substring Operation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func LexicographicallySmallestStringAfterSubstringOperation(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2734: Lexicographically Smallest String After Substring Operation
// https://leetcode.com/problems/lexicographically-smallest-string-after-substring-operation/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func LexicographicallySmallestStringAfterSubstringOperation(s string) string {
	n := len(s)
	b := []byte(s)

	start := -1
	for i := 0; i < n; i++ {
		if b[i] > 'a' {
			start = i
			break
		}
	}

	if start == -1 {
		b[n-1] = 'z'
		return string(b)
	}

	for i := start; i < n; i++ {
		if b[i] == 'a' {
			break
		}
		b[i]--
	}

	return string(b)
}

func main() {
	fmt.Println(LexicographicallySmallestStringAfterSubstringOperation("cbabc"))
	fmt.Println(LexicographicallySmallestStringAfterSubstringOperation("acbbc"))
	fmt.Println(LexicographicallySmallestStringAfterSubstringOperation("a"))
}
```
