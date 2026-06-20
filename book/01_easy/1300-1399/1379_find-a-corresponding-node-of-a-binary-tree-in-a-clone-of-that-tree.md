# 1379 — Find A Corresponding Node Of A Binary Tree In A Clone Of That Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func getTargetCopy(original, cloned *TreeNode, target *TreeNode) *TreeNode

import "fmt"

type TreeNode struct
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(h) where h is tree height  
**Kompleksitas Ruang:** O(h) where h is tree height

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1379: Find a Corresponding Node of a Binary Tree in a Clone of That Tree
// https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/
// Difficulty: Easy
//
// LeetCode submission: func getTargetCopy(original, cloned *TreeNode, target *TreeNode) *TreeNode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Original: [7,4,3,null,null,6,19]
	original := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 19},
		},
	}
	cloned := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 19},
		},
	}
	target := original.Right // node with value 3
	result := FindACorrespondingNodeOfABinaryTreeInACloneOfThatTree(original, cloned, target)
	fmt.Println(result.Val) // 3
}

// Time: O(n), Space: O(h) where h is tree height
func FindACorrespondingNodeOfABinaryTreeInACloneOfThatTree(original, cloned *TreeNode, target *TreeNode) *TreeNode {
	if original == nil {
		return nil
	}
	if original == target {
		return cloned
	}
	left := FindACorrespondingNodeOfABinaryTreeInACloneOfThatTree(original.Left, cloned.Left, target)
	if left != nil {
		return left
	}
	return FindACorrespondingNodeOfABinaryTreeInACloneOfThatTree(original.Right, cloned.Right, target)
}
```
