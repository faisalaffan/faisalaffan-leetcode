# 1668 — Maximum Repeating Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxRepeating(sequence string, word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m), Space: O(n) where n = len(sequence), m = len(word)  
**Kompleksitas Ruang:** O(n) where n = len(sequence), m = len(word)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1668: Maximum Repeating Substring
// https://leetcode.com/problems/maximum-repeating-substring/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n * m), Space: O(n) where n = len(sequence), m = len(word)
func MaxRepeating(sequence string, word string) int {
	k := 0
	repeated := word
	for strings.Contains(sequence, repeated) {
		k++
		repeated += word
	}
	return k
}

func main() {
	fmt.Println(MaxRepeating("ababc", "ab"))
	fmt.Println(MaxRepeating("ababc", "ba"))
	fmt.Println(MaxRepeating("ababc", "ac"))
}
```
