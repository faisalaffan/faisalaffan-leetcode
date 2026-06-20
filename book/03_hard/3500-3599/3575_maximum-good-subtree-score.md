# 3575 — Maximum Good Subtree Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func goodSubtreeSum(vals []int, par []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

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
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

  // Matriks 2D
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
