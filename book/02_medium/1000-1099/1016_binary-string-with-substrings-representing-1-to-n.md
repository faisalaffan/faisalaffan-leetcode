# 1016 — Binary String With Substrings Representing 1 To N

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func queryString(s string, n int) bool
```

> **💡 Hint:** Check if binary representation of each number from n down to 1

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * len(s) * log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1016: Binary String With Substrings Representing 1 To N
// https://leetcode.com/problems/binary-string-with-substrings-representing-1-to-n/
// Difficulty: Medium
//
// Approach: Check if binary representation of each number from n down to 1
//           is a substring of s. Start from n and go down for efficiency.
// Time: O(n * len(s) * log n)
// Space: O(log n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(queryString("0110", 3))  // true
	fmt.Println(queryString("0110", 4))  // false
	fmt.Println(queryString("1", 1))     // true
}

func queryString(s string, n int) bool {
	for i := n; i >= 1; i-- {
		binary := fmt.Sprintf("%b", i)
		if !strings.Contains(s, binary) {
			return false
		}
	}
	return true
}
```
