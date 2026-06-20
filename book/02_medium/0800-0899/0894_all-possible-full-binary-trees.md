# 0894 — All Possible Full Binary Trees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func AllPossibleFullBinaryTrees(n int) []*TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS, Dynamic Programming

**Kompleksitas Waktu:** O(2^n)  
**Kompleksitas Ruang:** O(2^n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #894: All Possible Full Binary Trees
// https://leetcode.com/problems/all-possible-full-binary-trees/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(len(AllPossibleFullBinaryTrees(7)))
	fmt.Println(len(AllPossibleFullBinaryTrees(3)))
	fmt.Println(len(AllPossibleFullBinaryTrees(1)))
}

// Time: O(2^n) | Space: O(2^n)
func AllPossibleFullBinaryTrees(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}

  // Membuat map (HashMap) — pencarian O(1)
	memo := make(map[int][]*TreeNode)

	var dfs func(int) []*TreeNode
	dfs = func(count int) []*TreeNode {
		if trees, ok := memo[count]; ok {
			return trees
		}

		if count == 1 {
			return []*TreeNode{{Val: 0}}
		}

		var res []*TreeNode
		for left := 1; left < count; left += 2 {
			right := count - 1 - left
			for _, l := range dfs(left) {
				for _, r := range dfs(right) {
					res = append(res, &TreeNode{Val: 0, Left: l, Right: r})
				}
			}
		}

		memo[count] = res
		return res
	}

	return dfs(n)
}
```
