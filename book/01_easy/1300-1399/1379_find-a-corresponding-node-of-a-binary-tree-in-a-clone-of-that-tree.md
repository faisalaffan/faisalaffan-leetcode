# 1379 — Find A Corresponding Node Of A Binary Tree In A Clone Of That Tree

## Deskripsi

**Soal:** [1379. Find A Corresponding Node Of A Binary Tree In A Clone Of That Tree](https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(h) where h is tree height  
**Kompleksitas Ruang:** O(h) where h is tree height

**Algoritma:** —

**Fungsi Solusi:** `func getTargetCopy(original, cloned *TreeNode, target *TreeNode) *TreeNode`

## Solusi Go

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
