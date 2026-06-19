package main

// LeetCode #975: Odd Even Jump
// https://leetcode.com/problems/odd-even-jump/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

func oddEvenJumps(A []int) int {
	n := len(A)
	if n == 0 {
		return 0
	}

	// odd[i] = can reach end from i with odd-numbered jump
	// even[i] = can reach end from i with even-numbered jump
	odd := make([]bool, n)
	even := make([]bool, n)
	odd[n-1] = true
	even[n-1] = true

	// TreeMap simulation: map from value to index
	// We process from right to left
	type pair struct {
		val int
		idx int
	}
	pairs := make([]pair, n)
	for i, v := range A {
		pairs[i] = pair{v, i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].val != pairs[j].val {
			return pairs[i].val < pairs[j].val
		}
		return pairs[i].idx < pairs[j].idx
	})

	// For each index, find the next odd jump (smallest value >= current)
	// and next even jump (largest value <= current)
	nextOdd := make([]int, n)
	nextEven := make([]int, n)
	for i := range nextOdd {
		nextOdd[i] = -1
		nextEven[i] = -1
	}

	// Use TreeMap structure: we process sorted values and maintain a set of indices
	// Actually, we can use the sorted list of values and use a balanced BST (simulated with sorted slice + binary search)

	// Simpler approach: for each index, find the smallest value >= A[i] to the right
	// and the largest value <= A[i] to the right
	for i := 0; i < n; i++ {
		// Find smallest value >= A[i] among indices > i
		minVal := int(1e9 + 1)
		minIdx := -1
		for j := i + 1; j < n; j++ {
			if A[j] >= A[i] && A[j] < minVal {
				minVal = A[j]
				minIdx = j
			}
		}
		nextOdd[i] = minIdx

		// Find largest value <= A[i] among indices > i
		maxVal := -1
		maxIdx := -1
		for j := i + 1; j < n; j++ {
			if A[j] <= A[i] && A[j] > maxVal {
				maxVal = A[j]
				maxIdx = j
			}
		}
		nextEven[i] = maxIdx
	}

	// DP from right to left
	for i := n - 2; i >= 0; i-- {
		if nextOdd[i] != -1 {
			odd[i] = even[nextOdd[i]]
		}
		if nextEven[i] != -1 {
			even[i] = odd[nextEven[i]]
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if odd[i] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println("Example 1:")
	fmt.Println(oddEvenJumps([]int{10, 13, 12, 14, 15}))
	// Expected: 2

	fmt.Println("Example 2:")
	fmt.Println(oddEvenJumps([]int{2, 3, 1, 1, 4}))
	// Expected: 3

	fmt.Println("Example 3:")
	fmt.Println(oddEvenJumps([]int{5, 1, 3, 4, 2}))
	// Expected: 3
}
