package main

// LeetCode #3145: Find Products of Elements of Big Array
// https://leetcode.com/problems/find-products-of-elements-of-big-array/
// Difficulty: Hard
//
// The "powerful array" of x is the shortest sorted array of powers of two that
// sum to x (the binary representation). The "big array" is the concatenation of
// powerful arrays for all positive integers. E.g., [1, 2, 1, 2, 4, 1, 4, ...].
// Each query [from, to, mod] asks for the product of big array elements in that
// range modulo mod. Since every element is a power of 2, product = 2^(exponent_sum).
// Use binary search + bit counting to compute prefix exponent sums.

import (
	"fmt"
	"math/big"
	"sort"
)

func main() {
	// queries with mod: [[from,to,mod], ...]
	queries := [][]int{
		{1, 3, 1000000007},
		{5, 7, 1000000007},
	}
	res := findProductsOfElementsOfBigArray(queries)
	fmt.Println(res)
}

// cnt1 returns the total number of set bits in numbers 1..num.
// This equals the total count of elements in the big array contributed by 1..num.
func cnt1(num int) int {
	if num <= 0 {
		return 0
	}
	res := 0
	for i := 0; 1<<uint(i) <= num; i++ {
		cycle := 1 << uint(i+1)
		cur := (num + 1) % cycle
		res += ((num + 1) / cycle) * (1 << uint(i))
		if cur > (1 << uint(i)) {
			res += cur - (1 << uint(i))
		}
	}
	return res
}

// acc0 returns the total sum of bit-position exponents for set bits in 1..num.
// This equals the prefix exponent sum of the big array contributed by 1..num.
func acc0(num int) int {
	if num <= 0 {
		return 0
	}
	res := 0
	for i := 0; 1<<uint(i) <= num; i++ {
		cycle := 1 << uint(i+1)
		cur := (num + 1) % cycle
		res += ((num + 1) / cycle) * (1 << uint(i)) * i
		if cur > (1 << uint(i)) {
			res += (cur - (1 << uint(i))) * i
		}
	}
	return res
}

// prefixExpSum returns the total exponent sum for big array elements [0..bound-1].
func prefixExpSum(bound int) int {
	if bound <= 0 {
		return 0
	}
	// Binary search for the smallest num where cnt1(num) >= bound
	target := sort.Search(bound, func(n int) bool {
		return cnt1(n) >= bound
	})

	// Number of elements before target's powerful array
	prevCnt := cnt1(target - 1)
	rest := bound - prevCnt // how many elements we need from target's array

	// Start with exponent sum from numbers before target
	expSum := acc0(target - 1)

	// Add exponents from target's set bits
	for i := 0; rest > 0; i++ {
		if target&(1<<uint(i)) != 0 {
			expSum += i
			rest--
		}
	}
	return expSum
}

func findProductsOfElementsOfBigArray(queries [][]int) []int {
	ans := make([]int, len(queries))
	for idx, q := range queries {
		from, to, mod := q[0], q[1], q[2]
		exp := prefixExpSum(to+1) - prefixExpSum(from)
		// Compute 2^exp % mod using big.Int for large exponents
		ans[idx] = int(new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(exp)), big.NewInt(int64(mod))).Int64())
	}
	return ans
}
