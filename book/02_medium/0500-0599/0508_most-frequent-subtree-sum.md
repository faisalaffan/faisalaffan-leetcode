# 0508 — Most Frequent Subtree Sum

## Deskripsi

**Soal:** [0508. Most Frequent Subtree Sum](https://leetcode.com/problems/most-frequent-subtree-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #508: Most Frequent Subtree Sum
// https://leetcode.com/problems/most-frequent-subtree-sum/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1: root = [5,2,-3]
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: -3}
	fmt.Println(FindFrequentTreeSum(root1))

	// Test case 2: root = [5,2,-5]
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: -5}
	fmt.Println(FindFrequentTreeSum(root2))
}

func FindFrequentTreeSum(root *TreeNode) []int {
  // Membuat map untuk pencarian O(1): key → value
	sumFreq := make(map[int]int)
	maxFreq := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		sumFreq[sum]++
		if sumFreq[sum] > maxFreq {
			maxFreq = sumFreq[sum]
		}
		return sum
	}

	dfs(root)

	result := []int{}
	for sum, freq := range sumFreq {
		if freq == maxFreq {
			result = append(result, sum)
		}
	}
	return result
}
```
