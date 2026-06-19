package main

// LeetCode #1969: Minimum Non-Zero Product of the Array Elements
// https://leetcode.com/problems/minimum-non-zero-product-of-the-array-elements/
// Difficulty: Medium

import "fmt"

const mod1969 = 1000000007

func main() {
	fmt.Println(MinNonZeroProduct(1))
	fmt.Println(MinNonZeroProduct(2))
	fmt.Println(MinNonZeroProduct(3))
}

// Time: O(p), Space: O(1)
func MinNonZeroProduct(p int) int {
	maxVal := (int64(1) << uint(p)) - 1
	base := maxVal - 1
	exp := (int64(1) << uint(p-1)) - 1
	result := int(maxVal % mod1969)
	result = int(int64(result) * powMod1969(base%int64(mod1969), exp) % mod1969)
	return result
}

func powMod1969(base int64, exp int64) int64 {
	result := int64(1)
	b := base % int64(mod1969)
	e := exp
	for e > 0 {
		if e&1 == 1 {
			result = (result * b) % int64(mod1969)
		}
		b = (b * b) % int64(mod1969)
		e >>= 1
	}
	return result
}
