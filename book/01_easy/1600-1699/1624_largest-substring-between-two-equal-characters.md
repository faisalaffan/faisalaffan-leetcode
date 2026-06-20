# 1624 — Largest Substring Between Two Equal Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxLengthBetweenEqualCharacters(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1) (since only 26 letters)  
**Kompleksitas Ruang:** O(1) (since only 26 letters)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1624: Largest Substring Between Two Equal Characters
// https://leetcode.com/problems/largest-substring-between-two-equal-characters/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1) (since only 26 letters)
func MaxLengthBetweenEqualCharacters(s string) int {
  // Membuat map (HashMap) — pencarian O(1)
	firstIndex := make(map[rune]int)
	maxLen := -1
	for i, ch := range s {
		if idx, exists := firstIndex[ch]; exists {
			if i-idx-1 > maxLen {
				maxLen = i - idx - 1
			}
		} else {
			firstIndex[ch] = i
		}
	}
	return maxLen
}

func main() {
	fmt.Println(MaxLengthBetweenEqualCharacters("aa"))
	fmt.Println(MaxLengthBetweenEqualCharacters("abca"))
	fmt.Println(MaxLengthBetweenEqualCharacters("cbzxy"))
}
```
