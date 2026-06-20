# 0508 — Most Frequent Subtree Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindFrequentTreeSum(root *TreeNode) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #508: Most Frequent Subtree Sum
// https://leetcode.com/problems/most-frequent-subtree-sum/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1: root = [5,2,-3]
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: -3}
	fmt.Println(FindFrequentTreeSum(root1))

	// Test case 2: root = [5,2,-5]
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: -5}
	fmt.Println(FindFrequentTreeSum(root2))
}

func FindFrequentTreeSum(root *TreeNode) []int {
  // Membuat map (HashMap) — pencarian O(1)
	sumFreq := make(map[int]int)
	maxFreq := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		sumFreq[sum]++
		if sumFreq[sum] > maxFreq {
			maxFreq = sumFreq[sum]
		}
		return sum
	}

	dfs(root)

	result := []int{}
	for sum, freq := range sumFreq {
		if freq == maxFreq {
			result = append(result, sum)
		}
	}
	return result
}
```
