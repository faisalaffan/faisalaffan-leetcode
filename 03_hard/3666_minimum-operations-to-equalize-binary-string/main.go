package main

// LeetCode #3666: Minimum Operations to Equalize Binary String
// https://leetcode.com/problems/minimum-operations-to-equalize-binary-string/
// Difficulty: Hard
//
// Given binary string s and integer k, in one operation flip exactly k indices.
// Return min operations to make all chars '1', or -1 if impossible.
//
// Approach: Track number of zeros. Each operation flips exactly k bits.
// If we have z zeros, we need to solve z + a*k - b*(n-z) = 0 mod 2 and reachable.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minOperationsBinary("110", 1))
	// Example 2
	fmt.Println(minOperationsBinary("0101", 3))
	// Example 3
	fmt.Println(minOperationsBinary("101", 2))
	// Edge: already all ones
	fmt.Println(minOperationsBinary("111", 2))
}

func minOperationsBinary(s string, k int) int {
	n := len(s)
	zeros := 0
	for _, ch := range s {
		if ch == '0' {
			zeros++
		}
	}
	if zeros == 0 {
		return 0
	}
	if k == 0 {
		return -1
	}

	// BFS over number of zeros
	visited := make([]bool, n+1)
	queue := make([]int, 0, n+1)
	queue = append(queue, zeros)
	visited[zeros] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			z := queue[i]
			if z == 0 {
				return steps
			}
			// Flip k bits: we flip x zeros and k-x ones
			// New zeros = z - x + (k - x) = z + k - 2x
			for x := 0; x <= k && x <= z; x++ {
				if k-x > n-z {
					continue
				}
				newZ := z + k - 2*x
				if newZ >= 0 && newZ <= n && !visited[newZ] {
					visited[newZ] = true
					queue = append(queue, newZ)
				}
			}
		}
		queue = queue[size:]
		steps++
	}

	return -1
}
