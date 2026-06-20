# 0656 — Coin Path

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func coinPath(coins []int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #656: Coin Path
// https://leetcode.com/problems/coin-path/
// Difficulty: Hard [Paid]
//
// Given an array coins where coins[i] is the cost of landing on position i
// (or -1 if position i is blocked), and a jump length k, find the lexicographically
// smallest path from position 0 to position n-1 with minimum total cost.
// The path always starts at 0 and ends at n-1 (both must be valid).

// Complexity: O(N*k) time, O(N) space where N = len(coins)

// coinPath returns the lexicographically smallest minimum-cost path from 0 to n-1.
// Returns empty slice if no path exists.
func coinPath(coins []int, k int) []int {
	n := len(coins)
	if n == 0 || coins[0] == -1 || coins[n-1] == -1 {
		return nil
	}

	// dp[i] = minimum total cost from i to n-1 (including coins[i]).
	// next[i] = next position after i on the optimal path.
  // Alokasi slice integer
	dp := make([]int, n)
  // Alokasi slice integer
	next := make([]int, n)

	// Initialize.
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt32
		next[i] = -1
	}

	// Base case: last position.
	dp[n-1] = coins[n-1]

	// Right-to-left DP.
	for i := n - 2; i >= 0; i-- {
		if coins[i] == -1 {
			continue
		}
		// Try all valid jumps from i.
		for j := i + 1; j <= i+k && j < n; j++ {
			if coins[j] == -1 || dp[j] == math.MaxInt32 {
				continue
			}
			cost := coins[i] + dp[j]
			if cost < dp[i] || (cost == dp[i] && j < next[i]) {
				// Lexicographically smallest: when costs are equal,
				// pick the smaller next index (since we're going R-to-L
				// and comparing j < next[i] picks the smaller index).
				dp[i] = cost
				next[i] = j
			}
		}
	}

	// No valid path.
	if dp[0] == math.MaxInt32 {
		return nil
	}

	// Reconstruct path.
	path := []int{0} // 0-indexed positions (1-indexed as per problem spec)
	curr := 0
	for curr < n-1 {
		curr = next[curr]
		path = append(path, curr)
	}
	return path
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0656 Coin Path ===")

	// Test 1: Basic case.
	// coins: [1,2,4,-1,2], k=2
	// Possible paths from 0 to 4:
	//   [0,2,4]: cost = 1+4+2 = 7
	//   [0,1,3] - can't, 3 is -1
	//   [0,1,4]: cost = 1+2+2 = 5
	//   [0,1,2,4]: cost = 1+2+4+2 = 9
	//   [0,2,4]: cost = 7
	// Min cost = 5 via [0,1,4]
	coins1 := []int{1, 2, 4, -1, 2}
	fmt.Printf("Test 1 - coins=%v, k=2\n", coins1)
	path1 := coinPath(coins1, 2)
	fmt.Printf("  Path: %v (expected [0 1 4])\n", path1)

	// Test 2: All positive, k=2.
	// coins: [1,2,3,4,5], k=2
	// From 0 -> 1 -> 3 -> 4: 1+2+4+5 = 12
	// From 0 -> 2 -> 4: 1+3+5 = 9
	// Min = 9 via [0,2,4]
	coins2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Test 2 - coins=%v, k=2\n", coins2)
	path2 := coinPath(coins2, 2)
	fmt.Printf("  Path: %v (expected [0 2 4])\n", path2)

	// Test 3: Lexicographically smallest tie-breaking.
	// If [0,3,5] and [0,4,5] have same cost, pick [0,3,5] (smaller first diff).
	coins3 := []int{1, 1, 1, 1, 1, 1}
	fmt.Printf("Test 3 - Tie-breaking, k=3, coins=%v\n", coins3)
	path3 := coinPath(coins3, 3)
	fmt.Printf("  Path: %v\n", path3)

	// Test 4: No path (blocked first position).
	coins4 := []int{-1, 2, 3}
	fmt.Printf("Test 4 - Blocked start: %v\n", coins4)
	path4 := coinPath(coins4, 2)
	fmt.Printf("  Path: %v (expected nil)\n", path4)

	// Test 5: No path (blocked last position).
	coins5 := []int{1, 2, -1}
	fmt.Printf("Test 5 - Blocked end: %v\n", coins5)
	path5 := coinPath(coins5, 2)
	fmt.Printf("  Path: %v (expected nil)\n", path5)

	// Test 6: Single element (already at end).
	coins6 := []int{10}
	fmt.Printf("Test 6 - Single element: %v\n", coins6)
	path6 := coinPath(coins6, 2)
	fmt.Printf("  Path: %v (expected [0])\n", path6)

	// Test 7: k larger than array.
	coins7 := []int{5, 3, 1}
	fmt.Printf("Test 7 - k larger than array: %v, k=10\n", coins7)
	path7 := coinPath(coins7, 10)
	fmt.Printf("  Path: %v (expected [0 2])\n", path7)

	// Test 8: Must step carefully.
	// coins: [1, -1, 1, -1, 1], k=2
	//  0 -> 2 -> 4: cost = 1+1+1 = 3
	//  0 -> 1 blocked
	//  So must be [0,2,4]
	coins8 := []int{1, -1, 1, -1, 1}
	fmt.Printf("Test 8 - Must skip blocks: %v, k=2\n", coins8)
	path8 := coinPath(coins8, 2)
	fmt.Printf("  Path: %v (expected [0 2 4])\n", path8)

	// Test 9: No path (can't reach end).
	// [1, -1, -1, 1], k=2
	// 0 -> 2 blocked, 0 -> 1 blocked. No path.
	coins9 := []int{1, -1, -1, 1}
	fmt.Printf("Test 9 - No reachable path: %v, k=2\n", coins9)
	path9 := coinPath(coins9, 2)
	fmt.Printf("  Path: %v (expected nil)\n", path9)

	// Test 10: Empty input.
	fmt.Printf("Test 10 - Empty: %v\n", coinPath(nil, 2))
}
```
