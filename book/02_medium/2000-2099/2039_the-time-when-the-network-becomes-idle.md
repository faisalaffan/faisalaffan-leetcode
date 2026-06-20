# 2039 — The Time When The Network Becomes Idle

## Deskripsi

**Soal:** [2039. The Time When The Network Becomes Idle](https://leetcode.com/problems/the-time-when-the-network-becomes-idle/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n + m)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

**Fungsi Solusi:** `func networkBecomesIdle(edges [][]int, patience []int) int`

## Solusi Go

```go
package main

// LeetCode #2039: The Time When the Network Becomes Idle
// https://leetcode.com/problems/the-time-when-the-network-becomes-idle/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)

import "fmt"

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// BFS to find shortest distance from node 0
  // Membuat slice untuk menyimpan hasil
	dist := make([]int, n)
  // Iterasi seluruh elemen
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0
	queue := []int{0}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if dist[v] == -1 {
				dist[v] = dist[u] + 1
				queue = append(queue, v)
			}
		}
	}

	maxTime := 0
	for i := 1; i < n; i++ {
		// Distance to master server and back
		roundTrip := dist[i] * 2
		p := patience[i]

		// Last message sent at: floor((roundTrip-1)/p) * p
		// Message arrives at: last_send + roundTrip
		lastSend := ((roundTrip - 1) / p) * p
		lastArrival := lastSend + roundTrip

		if lastArrival > maxTime {
			maxTime = lastArrival
		}
	}

	// Network becomes idle 1 ms after last message
	return maxTime + 1
}

func main() {
	// Test case 1
	edges1 := [][]int{{0, 1}, {1, 2}}
	patience1 := []int{0, 2, 1}
	fmt.Println("Test 1:", networkBecomesIdle(edges1, patience1))
	// Expected: 8

	// Test case 2
	edges2 := [][]int{{0, 1}, {0, 2}, {1, 2}}
	patience2 := []int{0, 10, 10}
	fmt.Println("Test 2:", networkBecomesIdle(edges2, patience2))
	// Expected: 3
}
```
