# 0133 — Clone Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func cloneGraph(node *Node) *Node
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(V+E)  
**Kompleksitas Ruang:** O(V)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #133: Clone Graph
// https://leetcode.com/problems/clone-graph/
// Difficulty: Medium

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

  // Membuat map (HashMap) — pencarian O(1)
	visited := make(map[*Node]*Node)

	var dfs func(n *Node) *Node
	dfs = func(n *Node) *Node {
		if clone, ok := visited[n]; ok {
			return clone
		}

		clone := &Node{Val: n.Val, Neighbors: make([]*Node, len(n.Neighbors))}
		visited[n] = clone

		for i, neighbor := range n.Neighbors {
			clone.Neighbors[i] = dfs(neighbor)
		}

		return clone
	}

	return dfs(node)
}

func printGraph(node *Node, visited map[*Node]bool) {
	if node == nil || visited[node] {
		return
	}
	visited[node] = true
	fmt.Printf("Node %d: [", node.Val)
	for i, n := range node.Neighbors {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(n.Val)
	}
	fmt.Println("]")
	for _, n := range node.Neighbors {
		printGraph(n, visited)
	}
}

func main() {
	// Test case 1: [[2,4],[1,3],[2,4],[1,3]]
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	clone := cloneGraph(n1)
	printGraph(clone, make(map[*Node]bool))

	// Test case 2
	fmt.Println(cloneGraph(nil)) // nil
}

// Time: O(V+E) | Space: O(V)
```
