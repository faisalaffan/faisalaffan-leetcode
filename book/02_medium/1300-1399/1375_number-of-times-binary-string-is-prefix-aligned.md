# 1375 — Number Of Times Binary String Is Prefix Aligned

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func numTimesAllBlue(flips []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n) where n = length of flips  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1375: Number of Times Binary String Is Prefix-Aligned
// https://leetcode.com/problems/number-of-times-binary-string-is-prefix-aligned/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numTimesAllBlue([]int{3, 2, 4, 1, 5})) // 2

	// Test case 2
	fmt.Println(numTimesAllBlue([]int{4, 1, 2, 3})) // 1

	// Test case 3
	fmt.Println(numTimesAllBlue([]int{2, 1, 3})) // 1
}

// Time: O(n) where n = length of flips
// Space: O(1)
func numTimesAllBlue(flips []int) int {
	count := 0
	maxFlip := 0

	for i, f := range flips {
		if f > maxFlip {
			maxFlip = f
		}
		if maxFlip == i+1 {
			count++
		}
	}

	return count
}
```
