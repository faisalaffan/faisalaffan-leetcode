# 1784 — Check If Binary String Has At Most One Segment Of Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckOnesSegment(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1784: Check if Binary String Has at Most One Segment of Ones
// https://leetcode.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(1)
func CheckOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}

func main() {
	fmt.Println(CheckOnesSegment("1001"))
	fmt.Println(CheckOnesSegment("110"))
	fmt.Println(CheckOnesSegment("1"))
}
```
