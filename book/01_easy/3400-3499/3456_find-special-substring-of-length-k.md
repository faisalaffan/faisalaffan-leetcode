# 3456 — Find Special Substring Of Length K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindSpecialSubstringOfLengthK(s string, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3456: Find Special Substring of Length K
// https://leetcode.com/problems/find-special-substring-of-length-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindSpecialSubstringOfLengthK("aaabaaa", 3))
	fmt.Println(FindSpecialSubstringOfLengthK("abc", 2))
}

// FindSpecialSubstringOfLengthK returns true if there is a substring of length k consisting of a single character, surrounded by different characters (or boundaries).
// Time: O(n). Space: O(1).
func FindSpecialSubstringOfLengthK(s string, k int) bool {
	n := len(s)
	for i := 0; i <= n-k; i++ {
		same := true
		for j := i; j < i+k-1; j++ {
			if s[j] != s[j+1] {
				same = false
				break
			}
		}
		if !same {
			continue
		}
		if i > 0 && s[i-1] == s[i] {
			continue
		}
		if i+k < n && s[i+k] == s[i] {
			continue
		}
		return true
	}
	return false
}
```
