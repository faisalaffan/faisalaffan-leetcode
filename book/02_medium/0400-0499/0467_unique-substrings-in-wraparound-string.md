# 0467 — Unique Substrings In Wraparound String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func UniqueSubstringsInWraparoundString(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (26 letters)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #467: Unique Substrings in Wraparound String
// https://leetcode.com/problems/unique-substrings-in-wraparound-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(1) (26 letters)

import "fmt"

func main() {
	fmt.Println(UniqueSubstringsInWraparoundString("a"))
	fmt.Println(UniqueSubstringsInWraparoundString("cac"))
	fmt.Println(UniqueSubstringsInWraparoundString("zab"))
}

func UniqueSubstringsInWraparoundString(s string) int {
  // Alokasi slice integer
	maxLen := make([]int, 26)
	curLen := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if i > 0 && (s[i]-s[i-1] == 1 || (s[i-1] == 'z' && s[i] == 'a')) {
			curLen++
		} else {
			curLen = 1
		}
		idx := s[i] - 'a'
		if curLen > maxLen[idx] {
			maxLen[idx] = curLen
		}
	}

	total := 0
	for _, v := range maxLen {
		total += v
	}
	return total
}
```
