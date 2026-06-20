# 2236 — Root Equals Sum Of Children

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func RootEqualsSumOfChildren(root *TreeNode) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2236: Root Equals Sum of Children
// https://leetcode.com/problems/root-equals-sum-of-children/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{10, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}}
	fmt.Println(RootEqualsSumOfChildren(root)) // true

	root2 := &TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}
	fmt.Println(RootEqualsSumOfChildren(root2)) // false
}

// Time: O(1), Space: O(1)
func RootEqualsSumOfChildren(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
```
