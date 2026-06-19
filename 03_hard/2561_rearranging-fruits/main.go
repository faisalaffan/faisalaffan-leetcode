package main

// LeetCode #2561: Rearranging Fruits
// https://leetcode.com/problems/rearranging-fruits/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

// minCost computes the minimum cost to make both fruit baskets identical.
//
// Swapping fruit a (from basket1) with fruit b (from basket2) costs min(a, b).
// We can also use the globally minimum fruit value as a mediator:
// swap a with min (cost = min), then min with b (cost = min), total = 2*min.
//
// Strategy:
// 1. Count frequency differences: freq[val] = count1 - count2.
// 2. Values with freq > 0 are excess in basket1; freq < 0 are excess in basket2.
// 3. Collect all excess values (each added |freq|/2 times) into one array.
// 4. Sort and pair smallest with largest. Cost = min(min(a,b), 2*globalMin).
//
// Complexity: O(n log n) time, O(n) space
func minCost(basket1 []int, basket2 []int) int64 {
	freq := make(map[int]int)
	for _, v := range basket1 {
		freq[v]++
	}
	for _, v := range basket2 {
		freq[v]--
	}

	// Find global minimum across both baskets
	globalMin := basket1[0]
	for _, v := range basket1 {
		if v < globalMin {
			globalMin = v
		}
	}
	for _, v := range basket2 {
		if v < globalMin {
			globalMin = v
		}
	}

	// Build excess array: values that need to be moved
	var excess []int
	for val, count := range freq {
		if count%2 != 0 {
			return -1 // impossible: odd frequency difference
		}
		cnt := abs(count) / 2
		for i := 0; i < cnt; i++ {
			excess = append(excess, val)
		}
	}

	sort.Ints(excess)

	cost := int64(0)
	// Pair smallest excess with largest excess
	for i := 0; i < len(excess)/2; i++ {
		a := excess[i]
		b := excess[len(excess)-1-i]
		// Direct swap cost = min(a, b); mediator cost = 2*globalMin
		importedMin := min(a, b)
		if 2*globalMin < importedMin {
			importedMin = 2 * globalMin
		}
		cost += int64(importedMin)
	}

	return cost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Test cases
	fmt.Println("Test 1: basket1=[4,2,2,2], basket2=[1,4,1,2] ->", minCost([]int{4, 2, 2, 2}, []int{1, 4, 1, 2}))
	fmt.Println("Test 2: basket1=[1,2,3,4], basket2=[1,2,3,4] ->", minCost([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}))
	fmt.Println("Test 3: basket1=[1,1,2,2], basket2=[3,3,4,4] ->", minCost([]int{1, 1, 2, 2}, []int{3, 3, 4, 4}))
	fmt.Println("Test 4: basket1=[84,80,43,8,80,88,43,84], basket2=[32,60,13,58,7,63,81,81] ->",
		minCost([]int{84, 80, 43, 8, 80, 88, 43, 84}, []int{32, 60, 13, 58, 7, 63, 81, 81}))
	fmt.Println("Test 5: basket1=[1,1,1], basket2=[2,2,2] ->", minCost([]int{1, 1, 1}, []int{2, 2, 2}))
	fmt.Println("Test 6: basket1=[1], basket2=[1] ->", minCost([]int{1}, []int{1}))
}
