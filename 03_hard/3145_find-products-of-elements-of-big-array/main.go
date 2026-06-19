package main

// LeetCode #3145: Find Products of Elements of Big Array
// https://leetcode.com/problems/find-products-of-elements-of-big-array/
// Difficulty: Hard
//
// A "big array" is formed by concatenating binary representations of all
// non-negative integers in order. For each query [from, to], compute the
// product of the VALUES at those positions. The value at a position is the
// sum of powers of 2 for each set bit in the (0-indexed) position within
// the current number's binary representation. This simplifies to: the big array
// equals [1,2,3,4,5,6,7,8,...] (the natural numbers). For queries [from,to],
// compute the product nums[from]*nums[from+1]*...*nums[to] % MOD.

import (
	"fmt"
	"math/big"
)

func main() {
	// queries=[[1,3],[5,7]]
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	queries := [][]int{{1, 3}, {5, 7}}
	res := findProductsOfElementsOfBigArray(nums, queries)
	fmt.Println(res)
}

func findProductsOfElementsOfBigArray(nums []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	prefixProd := make([]*big.Int, len(nums)+1)
	prefixProd[0] = big.NewInt(1)
	for i, v := range nums {
		prefixProd[i+1] = new(big.Int).Mul(prefixProd[i], big.NewInt(int64(v)))
	}
	for idx, q := range queries {
		from, to := q[0], q[1]
		prod := new(big.Int).Quo(prefixProd[to+1], prefixProd[from])
		ans[idx] = int(prod.Int64())
	}
	return ans
}
