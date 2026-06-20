# 2477 — Minimum Fuel Cost To Report To The Capital

## Deskripsi

**Soal:** [2477. Minimum Fuel Cost To Report To The Capital](https://leetcode.com/problems/minimum-fuel-cost-to-report-to-the-capital/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP), Tree DP (DP pada pohon)

## Solusi Go

```go
package main

// LeetCode #2477: Minimum Fuel Cost to Report to the Capital
// https://leetcode.com/problems/minimum-fuel-cost-to-report-to-the-capital/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Tree DP: Each person travels from leaf towards root (city 0).
// Accumulate people from children, compute fuel for each edge.

import "fmt"

func main() {
	fmt.Println(minimumFuelCost([][]int{{0, 1}, {0, 2}, {0, 3}}, 1)) // 3
	fmt.Println(minimumFuelCost([][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2)) // 7
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
  // Membuat slice 2D untuk DP/tabel
	graph := make([][]int, n)
	for _, r := range roads {
		a, b := r[0], r[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	var ans int64
	var dfs func(u, parent int) int
	dfs = func(u, parent int) int {
		people := 1 // each node has 1 representative
		for _, v := range graph[u] {
			if v != parent {
				people += dfs(v, u)
			}
		}
		if u != 0 {
			// cars needed = ceil(people / seats)
			cars := (people + seats - 1) / seats
			ans += int64(cars)
		}
		return people
	}
	dfs(0, -1)
	return ans
}
```
