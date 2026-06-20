# 3575 — Maximum Good Subtree Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func goodSubtreeSum(vals []int, par []int) int
```

> **💡 Hint:** DFS from root. For each node, check if all nodes in its subtree

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3575: Maximum Good Subtree Score
// https://leetcode.com/problems/maximum-good-subtree-score/
// Difficulty: Hard
//
// Given a tree with values on nodes, find the maximum sum of values in a "good"
// subtree. A subtree is good if all nodes in it share the same value.
//
// Approach: DFS from root. For each node, check if all nodes in its subtree
// have the same value. Track the maximum such sum.

import "fmt"

func main() {
	// Example 1
	fmt.Println(goodSubtreeSum([]int{1, 2, 3, 4, 5}, []int{-1, 0, 0, 1, 1}))
	// Example 2: all same values
	fmt.Println(goodSubtreeSum([]int{5, 5, 5}, []int{-1, 0, 0}))
	// Edge: single node
	fmt.Println(goodSubtreeSum([]int{10}, []int{-1}))
	// Edge: values alternate
	fmt.Println(goodSubtreeSum([]int{1, 2, 1}, []int{-1, 0, 0}))
}

func goodSubtreeSum(vals []int, par []int) int {
	n := len(vals)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

  // Membuat matriks/slice 2D untuk DP
	children := make([][]int, n)
	root := -1
	for i := 0; i < n; i++ {
		if par[i] == -1 {
			root = i
		} else {
			children[par[i]] = append(children[par[i]], i)
		}
	}

	maxScore := 0

	// DFS returns sum of subtree if all nodes in subtree have same value, else -1
	var dfs func(u int) int
	dfs = func(u int) int {
		sum := vals[u]
		for _, v := range children[u] {
			childSum := dfs(v)
			if childSum == -1 || vals[v] != vals[u] {
				sum = -1
			} else if sum != -1 {
				sum += childSum
			}
		}
		if sum != -1 && sum > maxScore {
			maxScore = sum
		}
		return sum
	}

	dfs(root)
	return maxScore
}
```
