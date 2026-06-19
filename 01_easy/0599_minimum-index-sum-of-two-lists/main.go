package main

// LeetCode #599: Minimum Index Sum of Two Lists
// https://leetcode.com/problems/minimum-index-sum-of-two-lists/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

// Time: O(n+m), Space: O(n)
func MinimumIndexSumOfTwoLists(list1, list2 []string) []string {
	index := make(map[string]int)
	for i, s := range list1 {
		index[s] = i
	}
	minSum := math.MaxInt32
	var result []string
	for j, s := range list2 {
		if i, ok := index[s]; ok {
			sum := i + j
			if sum < minSum {
				minSum = sum
				result = []string{s}
			} else if sum == minSum {
				result = append(result, s)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(MinimumIndexSumOfTwoLists(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
	))
	fmt.Println(MinimumIndexSumOfTwoLists(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"},
	))
}
