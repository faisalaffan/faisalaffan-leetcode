package main

// LeetCode #3226: Number of Bit Changes to Make Two Integers Equal
// https://leetcode.com/problems/number-of-bit-changes-to-make-two-integers-equal/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(NumberOfBitChangesToMakeTwoIntegersEqual(13, 4))
	fmt.Println(NumberOfBitChangesToMakeTwoIntegersEqual(21, 21))
	fmt.Println(NumberOfBitChangesToMakeTwoIntegersEqual(14, 13))
}

// NumberOfBitChangesToMakeTwoIntegersEqual returns the number of bit changes needed to make n equal to k, or -1 if impossible.
// Time: O(log n). Space: O(1).
func NumberOfBitChangesToMakeTwoIntegersEqual(n int, k int) int {
	// n must have all set bits that k has, since we can only change 1->0, not 0->1
	if n&k != k {
		return -1
	}
	// Count bits where n has 1 and k has 0
	diff := n ^ k
	count := 0
	for diff > 0 {
		count += diff & 1
		diff >>= 1
	}
	return count
}
