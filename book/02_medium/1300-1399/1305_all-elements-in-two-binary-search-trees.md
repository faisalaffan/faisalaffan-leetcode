# 1305 — All Elements In Two Binary Search Trees

## Deskripsi

**Soal:** [1305. All Elements In Two Binary Search Trees](https://leetcode.com/problems/all-elements-in-two-binary-search-trees/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m+n) for tree traversal + O(m+n) for merge = O(m+n)  
**Kompleksitas Ruang:** O(m+n) for storing values

**Algoritma:** Binary Search (pencarian biner)

## Solusi Go

```go
package main

// LeetCode #1305: All Elements in Two Binary Search Trees
// https://leetcode.com/problems/all-elements-in-two-binary-search-trees/
// Difficulty: Medium

import "fmt"
import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4}}
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 3}}
	fmt.Println(getAllElements(root1, root2)) // [0,1,1,2,3,4]

	// Test case 2
	root3 := &TreeNode{Val: 1, Right: &TreeNode{Val: 8}}
	root4 := &TreeNode{Val: 8, Left: &TreeNode{Val: 1}}
	fmt.Println(getAllElements(root3, root4)) // [1,1,8,8]

	// Test case 3
	fmt.Println(getAllElements(nil, nil)) // []
}

// Time: O(m+n) for tree traversal + O(m+n) for merge = O(m+n)
// Space: O(m+n) for storing values
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var vals1, vals2 []int
	inorder(root1, &vals1)
	inorder(root2, &vals2)
	return merge(vals1, vals2)
}

func inorder(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, vals)
	*vals = append(*vals, node.Val)
	inorder(node.Right, vals)
}

func merge(a, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

// Below ensures the signature matches expected problem name
func getAllElementsSort(root1 *TreeNode, root2 *TreeNode) []int {
	var vals []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root1)
	dfs(root2)
	sort.Ints(vals)
	return vals
}
```
