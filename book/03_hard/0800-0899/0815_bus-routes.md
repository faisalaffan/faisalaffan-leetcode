# 0815 — Bus Routes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func numBusesToDestination(routes [][]int, source int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, BFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #815: Bus Routes
// https://leetcode.com/problems/bus-routes/
// Difficulty: Hard
//
// Given an array routes where routes[i] is a bus route (list of stops), and a
// source and target stop, find the minimum number of buses needed to travel
// from source to target. Return -1 if impossible.
//
// Approach: BFS at the bus level.
//   - Build a mapping from stop → list of bus indices that serve that stop.
//   - BFS from source, tracking visited buses and visited stops.
//   - At each stop, we can board any bus that serves it. Taking that bus
//     visits all stops on its route. Increment bus count once per BFS level
//     (each bus taken adds 1).

import "fmt"

func main() {
	// Example: routes = [[1,2,7],[3,6,7]], source=1, target=6 → 2 (bus 0 to stop 7, then bus 1)
	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))

	// Same route: source and target both on bus 0 → 1
	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 2))

	// Source == target → 0
	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 5, 5))

	// No connection → -1
	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 100))

	// Single route
	fmt.Println(numBusesToDestination([][]int{{1, 2, 3, 4, 5, 6, 7}}, 1, 7))

	// Large: three routes
	fmt.Println(numBusesToDestination([][]int{{1, 9, 10}, {2, 3, 4}, {5, 6, 7, 8, 10}}, 1, 10))
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	// Map: stop → list of bus indices
  // HashMap: O(1) lookup
	stopToBuses := make(map[int][]int)
	for busIdx, stops := range routes {
		for _, stop := range stops {
			stopToBuses[stop] = append(stopToBuses[stop], busIdx)
		}
	}

	visitedBus := make([]bool, len(routes))
  // HashMap: O(1) lookup
	visitedStop := make(map[int]bool)
	queue := []int{source}
	visitedStop[source] = true
	buses := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			stop := queue[i]

			for _, busIdx := range stopToBuses[stop] {
				if visitedBus[busIdx] {
					continue
				}
				visitedBus[busIdx] = true

				for _, nextStop := range routes[busIdx] {
					if nextStop == target {
						return buses + 1
					}
					if !visitedStop[nextStop] {
						visitedStop[nextStop] = true
						queue = append(queue, nextStop)
					}
				}
			}
		}
		queue = queue[size:]
		buses++
	}

	return -1
}
```
