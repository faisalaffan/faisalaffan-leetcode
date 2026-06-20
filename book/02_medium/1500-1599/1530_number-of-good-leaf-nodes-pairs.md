# 1530 — Number Of Good Leaf Nodes Pairs

## Deskripsi

**Soal:** [1530. Number Of Good Leaf Nodes Pairs](https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N * distance^2), Space: O(N * distance)  
**Kompleksitas Ruang:** O(N * distance)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1530: Number of Good Leaf Nodes Pairs
// https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/
// Difficulty: Medium

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [1,2,3,null,4]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(CountPairs(root, 3))

	// Tree: [1,2,3,4,5,6,7]
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	root2.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}
	fmt.Println(CountPairs(root2, 3))

	// Single node
	fmt.Println(CountPairs(&TreeNode{Val: 1}, 1))
}

func CountPairs(root *TreeNode, distance int) int {
	// Time: O(N * distance^2), Space: O(N * distance)
	pairs := 0

	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			return []int{1} // leaf, distance 1 from here
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if left == nil && right == nil {
			return nil
		}

		// Count pairs between left and right subtrees
		if left != nil && right != nil {
			for _, ld := range left {
				for _, rd := range right {
					if ld+rd <= distance {
						pairs++
					}
				}
			}
		}

		// Merge distances, incrementing by 1 (edge to parent)
  // Membuat slice untuk menyimpan hasil
		result := make([]int, 0)
		if left != nil {
			for _, d := range left {
				if d+1 < distance {
					result = append(result, d+1)
				}
			}
		}
		if right != nil {
			for _, d := range right {
				if d+1 < distance {
					result = append(result, d+1)
				}
			}
		}

		return result
	}

	dfs(root)
	return pairs
}
```
