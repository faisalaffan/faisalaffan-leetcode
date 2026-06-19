package main

// LeetCode #2107: Number of Unique Flavors After Sharing K Candies
// https://leetcode.com/problems/number-of-unique-flavors-after-sharing-k-candies/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func shareCandies(candies []int, k int) int {
	if k >= len(candies) {
		return 0
	}

	freq := make(map[int]int)
	for i := k; i < len(candies); i++ {
		freq[candies[i]]++
	}

	maxUnique := len(freq)
	for i := k; i < len(candies); i++ {
		// Add candies[i-k] back (give it away)
		freq[candies[i-k]]++
		// Remove candies[i] from the kept set
		freq[candies[i]]--
		if freq[candies[i]] == 0 {
			delete(freq, candies[i])
		}
		if len(freq) > maxUnique {
			maxUnique = len(freq)
		}
	}

	return maxUnique
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", shareCandies([]int{1, 2, 2, 3, 4, 3}, 3))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", shareCandies([]int{1, 1, 2, 3}, 2))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", shareCandies([]int{1, 1, 1, 1}, 2))
	// Expected: 1
}
