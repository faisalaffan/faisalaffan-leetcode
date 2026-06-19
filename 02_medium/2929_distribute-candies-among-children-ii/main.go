package main

// LeetCode #2929: Distribute Candies Among Children II
// https://leetcode.com/problems/distribute-candies-among-children-ii/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(distributeCandies(5, 2))
	fmt.Println(distributeCandies(3, 3))
	fmt.Println(distributeCandies(10, 3))
}

func distributeCandies(n int, limit int) int64 {
	if n > 3*limit {
		return 0
	}
	ans := int64(n+2) * int64(n+1) / 2
	if n > limit {
		ans -= 3 * int64(n-limit+1) * int64(n-limit) / 2
	}
	if n-2 >= 2*limit {
		ans += 3 * int64(n-2*limit) * int64(n-2*limit+1) / 2
	}
	return ans
}
