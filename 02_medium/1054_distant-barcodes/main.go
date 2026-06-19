package main

// LeetCode #1054: Distant Barcodes
// https://leetcode.com/problems/distant-barcodes/
// Difficulty: Medium
//
// Approach: Count frequencies, place most frequent in even indices, then odd
// Time: O(n log n)
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rearrangeBarcodes([]int{1, 1, 1, 2, 2, 2})) // [1,2,1,2,1,2] or similar
	fmt.Println(rearrangeBarcodes([]int{1, 1, 1, 1, 2, 2, 3, 3})) // valid rearrangement
}

func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)
	freq := make(map[int]int)
	for _, b := range barcodes {
		freq[b]++
	}

	type pair struct {
		val   int
		count int
	}
	pairs := make([]pair, 0, len(freq))
	for val, count := range freq {
		pairs = append(pairs, pair{val, count})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	result := make([]int, n)
	idx := 0

	for _, p := range pairs {
		for k := 0; k < p.count; k++ {
			result[idx] = p.val
			idx += 2
			if idx >= n {
				idx = 1
			}
		}
	}

	return result
}
