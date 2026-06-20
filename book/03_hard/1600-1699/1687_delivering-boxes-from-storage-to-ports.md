# 1687 — Delivering Boxes From Storage To Ports

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sliding Window, DP, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sliding Window** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1687: Delivering Boxes from Storage to Ports
// https://leetcode.com/problems/delivering-boxes-from-storage-to-ports/
// Difficulty: Hard
// Strategy: DP with deque optimization (sliding window min).

import (
	"fmt"
)

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)

	// diffTrips[i] = number of port changes between box i-1 and box i
  // Alokasi slice
	diffTrips := make([]int, n+2)
	for i := 1; i < n; i++ {
		if boxes[i][0] != boxes[i-1][0] {
			diffTrips[i] = 1
		}
	}
	// prefix sum of diffTrips (size n+2 to allow safe access at index n+1)
  // Alokasi slice
	prefDiff := make([]int, n+2)
	for i := 1; i <= n; i++ {
		prefDiff[i] = prefDiff[i-1] + diffTrips[i-1]
	}
	prefDiff[n+1] = prefDiff[n] // diffTrips[n] is always 0

	// prefix sum of weights
  // Alokasi slice
	prefW := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefW[i] = prefW[i-1] + boxes[i-1][1]
	}

	// DP: dp[i] = min trips to deliver first i boxes
  // Alokasi slice
	dp := make([]int, n+1)
	// deque stores indices j, maintaining dp[j] - prefDiff[j+1] in increasing order
  // Alokasi slice
	deque := make([]int, 0, n+1)
	deque = append(deque, 0)

	for i := 1; i <= n; i++ {
		// Remove boxes that exceed maxBoxes or maxWeight from front
		for len(deque) > 0 {
			j := deque[0]
			if i-j > maxBoxes || prefW[i]-prefW[j] > maxWeight {
				deque = deque[1:]
			} else {
				break
			}
		}

		// dp[i] = dp[deque[0]] + prefDiff[i] - prefDiff[deque[0]+1] + 2
		j := deque[0]
		// Each trip: 1 for storage->first port + 1 for last port->storage = 2,
		// plus port changes between boxes
		dp[i] = dp[j] + prefDiff[i] - prefDiff[j+1] + 2

		// Insert i into deque
		val := dp[i] - prefDiff[i+1]
		for len(deque) > 0 {
			last := deque[len(deque)-1]
			if dp[last]-prefDiff[last+1] >= val {
				deque = deque[:len(deque)-1]
			} else {
				break
			}
		}
		deque = append(deque, i)
	}

	return dp[n]
}

func main() {
	// Example 1
	boxes1 := [][]int{{1, 1}, {2, 1}, {1, 1}}
	portsCount1 := 2
	maxBoxes1 := 3
	maxWeight1 := 3
	fmt.Printf("boxDelivering(%v, %d, %d, %d) = %d (expected 4)\n",
		boxes1, portsCount1, maxBoxes1, maxWeight1,
		boxDelivering(boxes1, portsCount1, maxBoxes1, maxWeight1))

	// Example 2
	boxes2 := [][]int{{1, 2}, {3, 3}, {3, 1}, {3, 1}, {2, 4}}
	portsCount2 := 3
	maxBoxes2 := 3
	maxWeight2 := 6
	fmt.Printf("boxDelivering(%v, %d, %d, %d) = %d (expected 6)\n",
		boxes2, portsCount2, maxBoxes2, maxWeight2,
		boxDelivering(boxes2, portsCount2, maxBoxes2, maxWeight2))

	// Example 3
	boxes3 := [][]int{{1, 4}, {1, 2}, {2, 1}, {2, 1}, {3, 2}, {3, 4}}
	portsCount3 := 3
	maxBoxes3 := 6
	maxWeight3 := 7
	fmt.Printf("boxDelivering(%v, %d, %d, %d) = %d (expected 6)\n",
		boxes3, portsCount3, maxBoxes3, maxWeight3,
		boxDelivering(boxes3, portsCount3, maxBoxes3, maxWeight3))

	// Example 4
	boxes4 := [][]int{{2, 4}, {2, 5}, {3, 1}, {3, 2}, {3, 7}, {3, 1}, {4, 4}, {1, 3}, {5, 2}}
	portsCount4 := 5
	maxBoxes4 := 5
	maxWeight4 := 7
	fmt.Printf("boxDelivering(%v, %d, %d, %d) = %d (expected 14)\n",
		boxes4, portsCount4, maxBoxes4, maxWeight4,
		boxDelivering(boxes4, portsCount4, maxBoxes4, maxWeight4))
}
```
