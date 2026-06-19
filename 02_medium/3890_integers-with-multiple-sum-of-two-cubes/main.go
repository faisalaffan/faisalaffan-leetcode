package main

// LeetCode #3890: Integers With Multiple Sum of Two Cubes
// https://leetcode.com/problems/integers-with-multiple-sum-of-two-cubes/
// Difficulty: Medium
// Time: O(C^2) precompute + O(N log N) sort | Space: O(C^2)
// Approach: Enumerate a,b up to 1000 (since 1000^3 = 1e9). Count frequency
// of each sum. Return sums with >= 2 representations, sorted.

import (
	"fmt"
	"sort"
)

func IntegersWithMultipleSumOfTwoCubes(n int) []int {
	cubeCount := make(map[int]int)
	limit := 1000
	for a := 1; a <= limit; a++ {
		a3 := a * a * a
		if a3 > n {
			break
		}
		for b := a; b <= limit; b++ {
			b3 := b * b * b
			sum := a3 + b3
			if sum > n {
				break
			}
			cubeCount[sum]++
		}
	}

	ans := []int{}
	for sum, cnt := range cubeCount {
		if cnt >= 2 {
			ans = append(ans, sum)
		}
	}
	sort.Ints(ans)
	return ans
}

func main() {
	// Example 1
	fmt.Println(IntegersWithMultipleSumOfTwoCubes(4104)) // Expected: [1729 4104]

	// Example 2
	fmt.Println(IntegersWithMultipleSumOfTwoCubes(578)) // Expected: []

	// Extra
	fmt.Println(IntegersWithMultipleSumOfTwoCubes(1729)) // Expected: [1729]
}
