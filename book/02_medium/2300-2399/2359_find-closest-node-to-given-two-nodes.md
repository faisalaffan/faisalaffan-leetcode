# 2359 — Find Closest Node To Given Two Nodes

## Deskripsi

**Soal:** [2359. Find Closest Node To Given Two Nodes](https://leetcode.com/problems/find-closest-node-to-given-two-nodes/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar), Queue (antrian FIFO)

**Fungsi Solusi:** `func closestMeetingNode(edges []int, node1 int, node2 int) int`

## Solusi Go

```go
package main

// LeetCode #2359: Find Closest Node to Given Two Nodes
// https://leetcode.com/problems/find-closest-node-to-given-two-nodes/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
  // Membuat slice untuk menyimpan hasil
	dist1 := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	dist2 := make([]int, n)
	for i := 0; i < n; i++ {
		dist1[i] = -1
		dist2[i] = -1
	}

	// BFS from node1
	queue := []int{node1}
	dist1[node1] = 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		v := edges[u]
		if v != -1 && dist1[v] == -1 {
			dist1[v] = dist1[u] + 1
			queue = append(queue, v)
		}
	}

	// BFS from node2
	queue = []int{node2}
	dist2[node2] = 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		v := edges[u]
		if v != -1 && dist2[v] == -1 {
			dist2[v] = dist2[u] + 1
			queue = append(queue, v)
		}
	}

	bestNode := -1
	bestDist := 1 << 30
	for i := 0; i < n; i++ {
		if dist1[i] != -1 && dist2[i] != -1 {
			maxDist := dist1[i]
			if dist2[i] > maxDist {
				maxDist = dist2[i]
			}
			if maxDist < bestDist {
				bestDist = maxDist
				bestNode = i
			}
		}
	}
	return bestNode
}

func main() {
	// Test case 1
	fmt.Println(closestMeetingNode([]int{2, 2, 3, -1}, 0, 1))
	// Expected: 2

	// Test case 2
	fmt.Println(closestMeetingNode([]int{1, 2, -1}, 0, 2))
	// Expected: 2

	// Test case 3
	fmt.Println(closestMeetingNode([]int{4, 4, 4, 5, -1, 0}, 1, 3))
	// Expected: 4
}
```
