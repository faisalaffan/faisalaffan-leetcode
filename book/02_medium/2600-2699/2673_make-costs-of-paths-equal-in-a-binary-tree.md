# 2673 — Make Costs Of Paths Equal In A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minIncrements(n int, cost []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2673: Make Costs of Paths Equal in a Binary Tree
// https://leetcode.com/problems/make-costs-of-paths-equal-in-a-binary-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minIncrements(n int, cost []int) int {
	ans := 0
	// Process from leaves to root (bottom-up)
	for i := n/2 - 1; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		// Make both child subtrees have equal max path sum
		if cost[left] > cost[right] {
			ans += cost[left] - cost[right]
			cost[i] += cost[left]
		} else {
			ans += cost[right] - cost[left]
			cost[i] += cost[right]
		}
	}
	return ans
}

func main() {
	// Test case 1: n=7, cost=[1,5,2,2,3,3,1]
	fmt.Println("Test 1:", minIncrements(7, []int{1, 5, 2, 2, 3, 3, 1}))
	// Expected: 6

	// Test case 2: n=3, cost=[5,3,3]
	fmt.Println("Test 2:", minIncrements(3, []int{5, 3, 3}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minIncrements(3, []int{10, 5, 5}))
	// Expected: 0
}
```
