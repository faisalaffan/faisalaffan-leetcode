package main

// LeetCode #875: Koko Eating Bananas
// https://leetcode.com/problems/koko-eating-bananas/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(KokoEatingBananas([]int{3, 6, 7, 11}, 8))
	fmt.Println(KokoEatingBananas([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(KokoEatingBananas([]int{30, 11, 23, 4, 20}, 6))
}

// Time: O(n log m) where m = max pile | Space: O(1)
func KokoEatingBananas(piles []int, h int) int {
	maxPile := 0
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}

	left, right := 1, maxPile
	for left < right {
		mid := left + (right-left)/2
		hours := 0
		for _, p := range piles {
			hours += (p + mid - 1) / mid
		}
		if hours <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
