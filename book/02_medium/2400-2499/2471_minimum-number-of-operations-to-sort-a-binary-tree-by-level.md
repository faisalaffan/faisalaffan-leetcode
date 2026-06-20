# 2471 — Minimum Number Of Operations To Sort A Binary Tree By Level

## Deskripsi

**Soal:** [2471. Minimum Number Of Operations To Sort A Binary Tree By Level](https://leetcode.com/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

## Solusi Go

```go
package main

// LeetCode #2471: Minimum Number of Operations to Sort a Binary Tree by Level
// https://leetcode.com/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// BFS level-order. For each level, count min swaps to sort (cycle decomposition).

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{1,
		&TreeNode{4,
			&TreeNode{7, nil, nil},
			&TreeNode{6, nil, nil},
		},
		&TreeNode{3,
			&TreeNode{8, nil, nil},
			&TreeNode{5, nil, nil},
		},
	}
	// Level 1: [1] sorted. Level 2: [4,3] -> swap, 1 op. Level 3: [7,6,8,5] -> 2 ops
	fmt.Println(minimumOperations(root)) // 3

	root2 := &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}
	fmt.Println(minimumOperations(root2)) // 0
}

func minimumOperations(root *TreeNode) int {
	q := []*TreeNode{root}
	ans := 0
	for len(q) > 0 {
		n := len(q)
  // Membuat slice untuk menyimpan hasil
		vals := make([]int, n)
		for i := 0; i < n; i++ {
			vals[i] = q[i].Val
		}

		// Count min swaps to sort vals
		ans += minSwaps(vals)

  // Membuat slice untuk menyimpan hasil
		next := make([]*TreeNode, 0)
		for _, node := range q {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		q = next
	}
	return ans
}

func minSwaps(arr []int) int {
	n := len(arr)
  // Membuat slice untuk menyimpan hasil
	sorted := make([]int, n)
	copy(sorted, arr)
	sort.Ints(sorted)

  // Membuat map untuk pencarian O(1): key → value
	pos := make(map[int]int)
	for i, v := range arr {
		pos[v] = i
	}

  // Membuat slice untuk menyimpan hasil
	visited := make([]bool, n)
	swaps := 0
	for i := 0; i < n; i++ {
		if visited[i] || arr[i] == sorted[i] {
			continue
		}
		cycle := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = pos[sorted[j]]
			cycle++
		}
		swaps += cycle - 1
	}
	return swaps
}
```
