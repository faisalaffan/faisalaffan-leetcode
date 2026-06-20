# 2689 — Extract Kth Character From The Rope Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func ExtractKthCharacterFromTheRopeTree(root *RopeNode, k int) byte`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(h)


## 💻 Solusi Go

```go
package main

// LeetCode #2689: Extract Kth Character From The Rope Tree
// https://leetcode.com/problems/extract-kth-character-from-the-rope-tree/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(h)
// Note: A rope tree is a binary tree where leaves contain substrings and internal nodes concatenate children.

import "fmt"

func main() {
	// Rope tree representing "abc" + "defg" = "abcdefg"
	root := &RopeNode{left: &RopeNode{val: "abc"}, right: &RopeNode{val: "defg"}}
	fmt.Println(string(ExtractKthCharacterFromTheRopeTree(root, 5))) // 'e'
}

type RopeNode struct {
	val   string
	left  *RopeNode
	right *RopeNode
}

func ExtractKthCharacterFromTheRopeTree(root *RopeNode, k int) byte {
	var dfs func(*RopeNode) string
	dfs = func(node *RopeNode) string {
		if node == nil {
			return ""
		}
		if node.left == nil && node.right == nil {
			return node.val
		}
		return dfs(node.left) + dfs(node.right)
	}
	return dfs(root)[k-1]
}
```
