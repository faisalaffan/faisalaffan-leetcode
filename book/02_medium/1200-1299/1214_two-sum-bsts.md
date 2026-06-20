# 1214 — Two Sum Bsts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n + m)  |  **Ruang:** O(n + m)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1214: Two Sum BSTs
// https://leetcode.com/problems/two-sum-bsts/
// Difficulty: Medium [Paid]

// Given two BSTs and a target, return true if there exists
// a node from each tree whose values sum to target.

// Time: O(n + m)
// Space: O(n + m)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
  // HashMap: O(1) lookup
	vals := make(map[int]bool)

	var collect func(node *TreeNode)
	collect = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals[node.Val] = true
		collect(node.Left)
		collect(node.Right)
	}
	collect(root1)

	var find func(node *TreeNode) bool
	find = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if vals[target-node.Val] {
			return true
		}
		return find(node.Left) || find(node.Right)
	}
	return find(root2)
}

func main() {
	// Tree1: [2,1,4], Tree2: [1,0,3], target=5
	r1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4}}
	r2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 3}}
	fmt.Printf("%t (expected: true)\n", twoSumBSTs(r1, r2, 5))

	// target=10 -> false
	fmt.Printf("%t (expected: false)\n", twoSumBSTs(r1, r2, 10))
}
```
