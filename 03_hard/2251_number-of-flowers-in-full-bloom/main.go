package main

// LeetCode #2251: Number of Flowers in Full Bloom
// https://leetcode.com/problems/number-of-flowers-in-full-bloom/
// Difficulty: Hard
//
// You are given a 0-indexed 2D integer array flowers where flowers[i] = [start_i, end_i]
// means the i-th flower will be in full bloom from start_i to end_i (inclusive).
// You are also given a 0-indexed integer array people of size n.
// For each person, return the number of flowers in full bloom at the time people[i].

import (
	"fmt"
	"sort"
)

// fullBloomFlowers returns for each person the count of flowers in bloom at their arrival time.
func fullBloomFlowers(flowers [][]int, people []int) []int {
	n := len(flowers)
	m := len(people)

	starts := make([]int, n)
	ends := make([]int, n)
	for i, f := range flowers {
		starts[i] = f[0]
		ends[i] = f[1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	result := make([]int, m)
	for i, p := range people {
		// number of flowers that started blooming <= p
		bloomed := sort.SearchInts(starts, p+1) // first index > p

		// number of flowers that ended blooming < p
		faded := sort.SearchInts(ends, p) // first index >= p

		result[i] = bloomed - faded
	}

	return result
}

func main() {
	// Example 1
	flowers1 := [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}
	people1 := []int{2, 3, 7, 11}
	fmt.Println(fullBloomFlowers(flowers1, people1)) // Expected: [1,2,2,2]

	// Example 2
	flowers2 := [][]int{{1, 10}, {3, 3}}
	people2 := []int{3, 3, 2}
	fmt.Println(fullBloomFlowers(flowers2, people2)) // Expected: [2,2,1]
}
