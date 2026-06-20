# 2096 — Step By Step Directions From A Binary Tree Node To Another

## Deskripsi

**Soal:** [2096. Step By Step Directions From A Binary Tree Node To Another](https://leetcode.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func getDirections(root *TreeNode, startValue int, destValue int) string`

## Solusi Go

```go
package main

// LeetCode #2096: Step-By-Step Directions From a Binary Tree Node to Another
// https://leetcode.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	var findPath func(node *TreeNode, target int, path []byte) ([]byte, bool)
	findPath = func(node *TreeNode, target int, path []byte) ([]byte, bool) {
		if node == nil {
			return nil, false
		}
		if node.Val == target {
			return path, true
		}
		if p, ok := findPath(node.Left, target, append(path, 'L')); ok {
			return p, true
		}
		if p, ok := findPath(node.Right, target, append(path, 'R')); ok {
			return p, true
		}
		return nil, false
	}

	pathToStart, _ := findPath(root, startValue, []byte{})
	pathToDest, _ := findPath(root, destValue, []byte{})

	// Find common prefix length
	i := 0
	for i < len(pathToStart) && i < len(pathToDest) && pathToStart[i] == pathToDest[i] {
		i++
	}

	// Go up from start to LCA (U for each remaining step)
  // Membuat slice untuk menyimpan hasil
	result := make([]byte, len(pathToStart)-i)
	for j := range result {
		result[j] = 'U'
	}

	// Go down from LCA to dest
	result = append(result, pathToDest[i:]...)

	return string(result)
}

func main() {
	// Test case 1
	root1 := &TreeNode{5,
		&TreeNode{1, &TreeNode{3, nil, nil}, nil},
		&TreeNode{2, &TreeNode{6, nil, nil}, &TreeNode{4, nil, nil}},
	}
	fmt.Println("Test 1:", getDirections(root1, 3, 6))
	// Expected: "UURL"

	// Test case 2
	root2 := &TreeNode{2, &TreeNode{1, nil, nil}, nil}
	fmt.Println("Test 2:", getDirections(root2, 2, 1))
	// Expected: "L"

	// Test case 3
	root3 := &TreeNode{1, nil, nil}
	fmt.Println("Test 3:", getDirections(root3, 1, 1))
	// Expected: ""
}
```
