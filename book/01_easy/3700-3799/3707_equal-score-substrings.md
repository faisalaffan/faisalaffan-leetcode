# 3707 — Equal Score Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func EqualScoreSubstrings(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3707: Equal Score Substrings
// https://leetcode.com/problems/equal-score-substrings/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(EqualScoreSubstrings("adcb"))
	fmt.Println(EqualScoreSubstrings("bace"))
}

// Time: O(n)
// Space: O(1)
func EqualScoreSubstrings(s string) bool {
	total := 0
	for _, ch := range s {
		total += int(ch-'a') + 1
	}

	prefix := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s)-1; i++ {
		prefix += int(s[i]-'a') + 1
		if prefix == total-prefix {
			return true
		}
	}
	return false
}
```
