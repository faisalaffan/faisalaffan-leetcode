# 1689 — Partitioning Into Minimum Number Of Deci Binary Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minPartitions(n string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1689: Partitioning Into Minimum Number Of Deci-Binary Numbers
// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minPartitions(n string) int {
	maxDigit := 0
	for _, ch := range n {
		digit := int(ch - '0')
		if digit > maxDigit {
			maxDigit = digit
		}
		if maxDigit == 9 {
			break // can't get higher than 9
		}
	}
	return maxDigit
}

func main() {
	fmt.Println(minPartitions("32"))     // Expected: 3
	fmt.Println(minPartitions("82734"))  // Expected: 8
	fmt.Println(minPartitions("27346209830709182346")) // Expected: 9
}
```
