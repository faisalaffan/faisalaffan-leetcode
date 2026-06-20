# 1120 — Maximum Average Subtree

## Deskripsi

**Soal:** [1120. Maximum Average Subtree](https://leetcode.com/problems/maximum-average-subtree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

> **Ide Kunci:** DFS post-order. Return sum and count for each subtree.

## Solusi Go

```go
package main

// LeetCode #1120: Maximum Average Subtree
// https://leetcode.com/problems/maximum-average-subtree/
// Difficulty: Medium
//
// Approach: DFS post-order. Return sum and count for each subtree.
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
		Val: 5,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(maximumAverageSubtree(root)) // 6.0

	root2 := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(maximumAverageSubtree(root2)) // 1.5
}

func maximumAverageSubtree(root *TreeNode) float64 {
	maxAvg := 0.0
	dfs(root, &maxAvg)
	return maxAvg
}

func dfs(node *TreeNode, maxAvg *float64) (int, int) {
	if node == nil {
		return 0, 0
	}

	leftSum, leftCount := dfs(node.Left, maxAvg)
	rightSum, rightCount := dfs(node.Right, maxAvg)

	sum := node.Val + leftSum + rightSum
	count := 1 + leftCount + rightCount

	avg := float64(sum) / float64(count)
	if avg > *maxAvg {
		*maxAvg = avg
	}

	return sum, count
}
```
