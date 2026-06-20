# 2947 — Count Beautiful Substrings I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func beautifulSubstrings(s string, k int) (ans int)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2947: Count Beautiful Substrings I
// https://leetcode.com/problems/count-beautiful-substrings-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(beautifulSubstrings("baeyh", 2))
	fmt.Println(beautifulSubstrings("ab", 1))
	fmt.Println(beautifulSubstrings("a", 1))
}

func beautifulSubstrings(s string, k int) (ans int) {
	n := len(s)
	vowels := [26]bool{}
	for _, c := range "aeiou" {
		vowels[c-'a'] = true
	}
	for i := 0; i < n; i++ {
		v := 0
		for j := i; j < n; j++ {
			if vowels[s[j]-'a'] {
				v++
			}
			c := j - i + 1 - v
			if v == c && v*c%k == 0 {
				ans++
			}
		}
	}
	return
}
```
