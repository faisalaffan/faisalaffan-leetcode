package main

// LeetCode #2927: Distribute Candies Among Children III
// https://leetcode.com/problems/distribute-candies-among-children-iii/
// Difficulty: Hard [Paid]
//
// Distribute n identical candies to 3 distinct children, each child gets at most
// `limit` candies. Count the number of ways.
//
// Combinatorics with inclusion-exclusion:
// Total = C(n+2, 2) (nonnegative integer solutions to x+y+z=n)
// Subtract: cases where child 1 > limit, child 2 > limit, child 3 > limit
// Add back: cases where two children > limit
// Subtract: cases where all three > limit
//
// For a child exceeding limit, give them limit+1 candies first,
// then distribute the rest: C(n-(limit+1)+2, 2)

import "fmt"

func distributeCandies(n int, limit int) int {
	// C(n+2, 2) = (n+2)*(n+1)/2
	total := comb2(n + 2)
	if total == 0 {
		return 0
	}

	// Subtract: one child exceeds limit (gets limit+1 fixed, then distribute rest)
	// C(n-(limit+1)+2, 2) = C(n-limit+1, 2)
	oneExceed := 3 * comb2(n-limit+1)
	if oneExceed < 0 {
		oneExceed = 0
	}
	total -= oneExceed

	// Add back: two children exceed limit
	// C(n-2(limit+1)+2, 2) = C(n-2*limit, 2)
	twoExceed := 3 * comb2(n-2*limit)
	if twoExceed < 0 {
		twoExceed = 0
	}
	total += twoExceed

	// Subtract: all three exceed limit
	// C(n-3(limit+1)+2, 2) = C(n-3*limit-1, 2)
	threeExceed := comb2(n - 3*limit - 1)
	if threeExceed < 0 {
		threeExceed = 0
	}
	total -= threeExceed

	return total
}

// C(x, 2) = x*(x-1)/2 for x >= 2, 0 for x <= 1
// Assumes n >= 0 is the argument to C(n, 2) which equals 0 for n < 2.
// But we call comb2(n+2) so the result is always >= 1 for n >= 0.
func comb2(n int) int {
	if n < 2 {
		return 0
	}
	return n * (n - 1) / 2
}

func main() {
	// Example: n=3, limit=3
	// All partitions of 3 with each <= 3: (0,0,3),(0,1,2),(0,2,1),(0,3,0),(1,0,2),
	// (1,1,1),(1,2,0),(2,0,1),(2,1,0),(3,0,0) = 10
	fmt.Println(distributeCandies(3, 3))

	// n=5, limit=2 => only (2,2,1) permutations = 3
	fmt.Println(distributeCandies(5, 2))

	// n=0, limit=3 => (0,0,0) = 1
	fmt.Println(distributeCandies(0, 3))

	// n=6, limit=2 => only (2,2,2) = 1
	fmt.Println(distributeCandies(6, 2))

	// n=10, limit=5
	fmt.Println(distributeCandies(10, 5))

	// n=100, limit=50
	fmt.Println(distributeCandies(100, 50))

	// n=1000, limit=500
	fmt.Println(distributeCandies(1000, 500))
}
