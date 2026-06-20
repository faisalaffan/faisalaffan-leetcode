# 2662 — Minimum Cost Of A Path With Special Roads

## Deskripsi

**Soal:** [2662. Minimum Cost Of A Path With Special Roads](https://leetcode.com/problems/minimum-cost-of-a-path-with-special-roads/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dijkstra (lintasan terpendek)

**Fungsi Solusi:** `func minimumCost(start []int, target []int, specialRoads [][]int) int`

## Solusi Go

```go
package main

// LeetCode #2662: Minimum Cost of a Path With Special Roads
// https://leetcode.com/problems/minimum-cost-of-a-path-with-special-roads/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func minimumCost(start []int, target []int, specialRoads [][]int) int {
	// Use Dijkstra: dist[i] = min cost to reach special road i's end
	n := len(specialRoads)
  // Membuat slice untuk menyimpan hasil
	dist := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	visited := make([]bool, n)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	minDist := abs(target[0]-start[0]) + abs(target[1]-start[1])

	// Initialize dist with cost to reach each special road's start from (start)
	for i, road := range specialRoads {
		costToStart := abs(road[0]-start[0]) + abs(road[1]-start[1])
		dist[i] = costToStart + road[4] // cost to travel the special road
	}

	for {
		// Find unvisited node with minimum dist
		u := -1
		best := int(1e18)
		for i := 0; i < n; i++ {
			if !visited[i] && dist[i] < best {
				best = dist[i]
				u = i
			}
		}
		if u == -1 {
			break
		}
		visited[u] = true

		// From road u's end, try going directly to target
		costToTarget := dist[u] + abs(target[0]-specialRoads[u][2]) + abs(target[1]-specialRoads[u][3])
		if costToTarget < minDist {
			minDist = costToTarget
		}

		// From road u's end, try going to another special road's start
		for v, road := range specialRoads {
			if visited[v] {
				continue
			}
			cost := dist[u] + abs(road[0]-specialRoads[u][2]) + abs(road[1]-specialRoads[u][3]) + road[4]
			if cost < dist[v] {
				dist[v] = cost
			}
		}
	}

	return minDist
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumCost([]int{1, 1}, []int{4, 5}, [][]int{{1, 2, 3, 3, 2}, {3, 4, 4, 5, 1}}))
	// Expected: 5

	// Test case 2: no special roads
	fmt.Println("Test 2:", minimumCost([]int{0, 0}, []int{3, 4}, [][]int{}))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", minimumCost([]int{1, 1}, []int{10, 10}, [][]int{{5, 5, 6, 6, 1}}))
	// Expected: 18
}
```
