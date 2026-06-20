# 0753 — Cracking The Safe

## Deskripsi

**Soal:** [0753. Cracking The Safe](https://leetcode.com/problems/cracking-the-safe/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func crackSafe(n int, k int) string`

## Solusi Go

```go
package main

// LeetCode #753: Cracking the Safe
// https://leetcode.com/problems/cracking-the-safe/
// Difficulty: Hard
//
// Algorithm: De Bruijn Sequence / Hierholzer (Eulerian Path)
// Construct a de Bruijn graph where:
// - Nodes are all (n-1)-digit k-ary strings
// - Edges are n-digit k-ary strings labeled with the last digit
// - Find an Eulerian path covering all edges exactly once
// - The concatenation of edge labels gives the shortest password

import (
	"fmt"
	"strconv"
	"strings"
)

func crackSafe(n int, k int) string {
	if n == 1 {
		// Return all single digits
		var sb strings.Builder
		for i := 0; i < k; i++ {
			sb.WriteString(strconv.Itoa(i))
		}
		return sb.String()
	}

	// Total nodes = k^(n-1)
	totalNodes := 1
	for i := 0; i < n-1; i++ {
		totalNodes *= k
	}

	// Adjacency list: for each node, track which edges (digits) we've used
  // Membuat map untuk pencarian O(1): key → value
	visited := make(map[int]map[int]bool)

	// The result path
	var result []int

	// Hierholzer's algorithm (DFS)
	var dfs func(node int)
	dfs = func(node int) {
		for d := 0; d < k; d++ {
			if !visited[node][d] {
				if visited[node] == nil {
					visited[node] = make(map[int]bool)
				}
				visited[node][d] = true
				// Next node: shift left, add new digit
				nextNode := (node*k + d) % totalNodes
				dfs(nextNode)
				result = append(result, d)
			}
		}
	}

	// Start from node 0
	dfs(0)

	// Build the password
	// Start with the initial node (n-1 zeros)
	var sb strings.Builder
	for i := 0; i < n-1; i++ {
		sb.WriteByte('0')
	}
	// Append edges in reverse order (since DFS appends after recursive call)
	for i := len(result) - 1; i >= 0; i-- {
		sb.WriteString(strconv.Itoa(result[i]))
	}

	return sb.String()
}

func main() {
	// Example from problem
	n1, k1 := 1, 2
	result1 := crackSafe(n1, k1)
	fmt.Printf("Input: n=%d, k=%d\nOutput: %q (expected: 01 or 10)\n\n", n1, k1, result1)

	// Example from problem
	n2, k2 := 2, 2
	result2 := crackSafe(n2, k2)
	fmt.Printf("Input: n=%d, k=%d\nOutput: %q (expected: 00110 or 01100 or 10011 or 11001)\n\n", n2, k2, result2)

	// Test case 3
	n3, k3 := 2, 3
	result3 := crackSafe(n3, k3)
	fmt.Printf("Input: n=%d, k=%d\nOutput: %q (length=%d, expected length=%d)\n\n", n3, k3, result3, len(result3), 10)

	// Test case 4: n=2, k=1
	n4, k4 := 2, 1
	result4 := crackSafe(n4, k4)
	fmt.Printf("Input: n=%d, k=%d\nOutput: %q\n\n", n4, k4, result4)

	// Test case 5: n=3, k=2
	n5, k5 := 3, 2
	result5 := crackSafe(n5, k5)
	fmt.Printf("Input: n=%d, k=%d\nOutput: %q (length=%d, expected length=%d)\n\n", n5, k5, result5, len(result5), 10)

	// Verify coverage for n=2,k=2
	fmt.Println("Verification: n=2, k=2 passwords covered:")
	inputs2 := []string{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			inputs2 = append(inputs2, strconv.Itoa(i)+strconv.Itoa(j))
		}
	}
	covered := 0
	for _, pw := range inputs2 {
		if strings.Contains(result2, pw) {
			covered++
		}
	}
	fmt.Printf("  Covered %d/4 passwords: %v\n", covered, covered == 4)

	// Verify coverage for n=3,k=2
	inputs3 := []string{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for kk := 0; kk < 2; kk++ {
				inputs3 = append(inputs3, strconv.Itoa(i)+strconv.Itoa(j)+strconv.Itoa(kk))
			}
		}
	}
	covered3 := 0
	for _, pw := range inputs3 {
		if strings.Contains(result5, pw) {
			covered3++
		}
	}
	fmt.Printf("  Covered %d/8 passwords: %v\n", covered3, covered3 == 8)
}
```
