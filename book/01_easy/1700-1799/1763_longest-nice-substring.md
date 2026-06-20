# 1763 — Longest Nice Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestNiceSubstring(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1763: Longest Nice Substring
// https://leetcode.com/problems/longest-nice-substring/
// Difficulty: Easy

import "fmt"
import "unicode"

// Time: O(n^2), Space: O(n)
func LongestNiceSubstring(s string) string {
	result := ""
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		lower := 0
		upper := 0
		for j := i; j < len(s); j++ {
			ch := rune(s[j])
			if unicode.IsUpper(ch) {
				upper |= 1 << (unicode.ToLower(ch) - 'a')
			} else {
				lower |= 1 << (ch - 'a')
			}
			if lower == upper && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func main() {
	fmt.Println(LongestNiceSubstring("YazaAay"))
	fmt.Println(LongestNiceSubstring("Bb"))
	fmt.Println(LongestNiceSubstring("c"))
}
```
