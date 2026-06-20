# 2385 — Amount Of Time For Binary Tree To Be Infected

## Deskripsi

**Soal:** [2385. Amount Of Time For Binary Tree To Be Infected](https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

## Solusi Go

```go
package main

// LeetCode #2385: Amount of Time for Binary Tree to Be Infected
// https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Convert tree to graph, BFS from start node to find farthest distance.

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [1,5,3,null,4,10,6,9,2]
	root := &TreeNode{1,
		&TreeNode{5, nil,
			&TreeNode{4,
				&TreeNode{9, nil, nil},
				&TreeNode{2, nil, nil},
			},
		},
		&TreeNode{3,
			&TreeNode{10, nil, nil},
			&TreeNode{6, nil, nil},
		},
	}
	fmt.Println(amountOfTime(root, 3)) // 4

	root2 := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3,
				&TreeNode{4,
					&TreeNode{5, nil, nil},
					nil,
				},
				nil,
			},
			nil,
		},
		nil,
	}
	fmt.Println(amountOfTime(root2, 1)) // 4
}

func amountOfTime(root *TreeNode, start int) int {
  // Membuat map untuk pencarian O(1): key → value
	graph := make(map[int][]int)
	var buildGraph func(node *TreeNode)
	buildGraph = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
			buildGraph(node.Left)
		}
		if node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
			buildGraph(node.Right)
		}
	}
	buildGraph(root)

	visited := map[int]bool{start: true}
	queue := []int{start}
	time := -1
	for len(queue) > 0 {
		time++
		for sz := len(queue); sz > 0; sz-- {
			u := queue[0]
			queue = queue[1:]
			for _, v := range graph[u] {
				if !visited[v] {
					visited[v] = true
					queue = append(queue, v)
				}
			}
		}
	}
	return time
}
```
