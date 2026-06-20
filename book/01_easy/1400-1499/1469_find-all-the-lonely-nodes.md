# 1469 — Find All The Lonely Nodes

## Deskripsi

**Soal:** [1469. Find All The Lonely Nodes](https://leetcode.com/problems/find-all-the-lonely-nodes/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(h) recursive + O(n) output  
**Kompleksitas Ruang:** O(h) recursive + O(n) output

**Algoritma:** —

**Fungsi Solusi:** `func getLonelyNodes(root *TreeNode) []int`

## Solusi Go

```go
package main

// LeetCode #1469: Find All The Lonely Nodes
// https://leetcode.com/problems/find-all-the-lonely-nodes/
// Difficulty: Easy [Paid]
//
// LeetCode submission: func getLonelyNodes(root *TreeNode) []int

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [1,2,3,null,4]
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(FindAllTheLonelyNodes(root)) // [4]

	// Tree: [7,1,4,6,null,5,3,null,null,null,null,null,2]
	root2 := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 6}},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}},
	}
	fmt.Println(FindAllTheLonelyNodes(root2)) // [6 5 2]
}

// Time: O(n), Space: O(h) recursive + O(n) output
func FindAllTheLonelyNodes(root *TreeNode) []int {
  // Membuat slice untuk menyimpan hasil
	res := make([]int, 0)
	collectLonely(root, &res)
	return res
}

func collectLonely(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil && node.Right == nil {
		*res = append(*res, node.Left.Val)
	}
	if node.Left == nil && node.Right != nil {
		*res = append(*res, node.Right.Val)
	}
	collectLonely(node.Left, res)
	collectLonely(node.Right, res)
}
```
