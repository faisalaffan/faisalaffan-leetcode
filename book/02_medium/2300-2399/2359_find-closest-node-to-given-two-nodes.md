# 2359 — Find Closest Node To Given Two Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func closestMeetingNode(edges []int, node1 int, node2 int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2359: Find Closest Node to Given Two Nodes
// https://leetcode.com/problems/find-closest-node-to-given-two-nodes/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
  // Alokasi slice
	dist1 := make([]int, n)
  // Alokasi slice
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
