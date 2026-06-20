# 3787 — Find Diameter Endpoints Of A Tree

## Deskripsi

**Soal:** [3787. Find Diameter Endpoints Of A Tree](https://leetcode.com/problems/find-diameter-endpoints-of-a-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP), BFS (Breadth-First Search / pencarian lebar)

**Fungsi Solusi:** `func findDiameterEndpointsOfATree(n int, edges [][]int) string`

## Solusi Go

```go
package main

// LeetCode #3787: Find Diameter Endpoints of a Tree
// https://leetcode.com/problems/find-diameter-endpoints-of-a-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func findDiameterEndpointsOfATree(n int, edges [][]int) string {
  // Membuat slice 2D untuk DP/tabel
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	bfs := func(start int) (int, []int) {
  // Membuat slice untuk menyimpan hasil
		dist := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = -1
		}
		dist[start] = 0
		q := []int{start}
		far := start
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			if dist[u] > dist[far] {
				far = u
			}
			for _, v := range g[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					q = append(q, v)
				}
			}
		}
		return far, dist
	}

	a, _ := bfs(0)
	b, distA := bfs(a)
	_, distB := bfs(b)
	diameter := distA[b]

  // Membuat slice untuk menyimpan hasil
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		if distA[i] == diameter || distB[i] == diameter {
			ans[i] = '1'
		} else {
			ans[i] = '0'
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(findDiameterEndpointsOfATree(7, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {3, 5}, {1, 6}}))
	fmt.Println(findDiameterEndpointsOfATree(3, [][]int{{0, 1}, {1, 2}}))
	fmt.Println(findDiameterEndpointsOfATree(1, [][]int{}))
}
```
