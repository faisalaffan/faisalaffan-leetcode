# 0096 — Unique Binary Search Trees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func numTrees(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #96: Unique Binary Search Trees
// https://leetcode.com/problems/unique-binary-search-trees/
// Difficulty: Medium

import "fmt"

func numTrees(n int) int {
  // Alokasi slice integer
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println(numTrees(3)) // 5

	// Test case 2
	fmt.Println(numTrees(1)) // 1

	// Test case 3
	fmt.Println(numTrees(4)) // 14
}

// Time: O(n^2) | Space: O(n)
```
