# 1123 — Lowest Common Ancestor Of Deepest Leaves

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func lcaDeepestLeaves(root *TreeNode) *TreeNode
```

> **💡 Hint:** DFS. Return depth and LCA candidate for each subtree.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1123: Lowest Common Ancestor of Deepest Leaves
// https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/
// Difficulty: Medium
//
// Approach: DFS. Return depth and LCA candidate for each subtree.
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7, Left: nil, Right: nil}, Right: &TreeNode{Val: 4, Left: nil, Right: nil}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
			Right: &TreeNode{Val: 8, Left: nil, Right: nil},
		},
	}
	result := lcaDeepestLeaves(root)
	fmt.Println(result.Val) // 2
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := dfs(root)
	return lca
}

func dfs(node *TreeNode) (int, *TreeNode) {
	if node == nil {
		return 0, nil
	}

	leftDepth, leftLCA := dfs(node.Left)
	rightDepth, rightLCA := dfs(node.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1, leftLCA
	} else if rightDepth > leftDepth {
		return rightDepth + 1, rightLCA
	}
	return leftDepth + 1, node
}
```
