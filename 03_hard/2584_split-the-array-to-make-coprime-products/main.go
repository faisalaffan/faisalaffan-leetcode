package main

// LeetCode #2584: Split the Array to Make Coprime Products
// https://leetcode.com/problems/split-the-array-to-make-coprime-products/
// Difficulty: Hard

import "fmt"

// findValidSplit finds the smallest index i such that
// gcd(product(nums[0..i]), product(nums[i+1..n-1])) == 1.
//
// For each prime factor, track its last occurrence. Scan left to right,
// tracking the furthest last occurrence of any prime seen so far.
// When furthest == current position, we found a valid split.
//
// Complexity: O(n * sqrt(max(nums))) time, O(distinct primes) space
func findValidSplit(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	// Track last occurrence of each prime factor
	last := make(map[int]int)

	// Factorize a number into its distinct prime factors
	factorize := func(x int) []int {
		if x == 1 {
			return nil
		}
		var factors []int
		v := x
		for p := 2; p*p <= v; p++ {
			if v%p == 0 {
				factors = append(factors, p)
				for v%p == 0 {
					v /= p
				}
			}
		}
		if v > 1 {
			factors = append(factors, v)
		}
		return factors
	}

	// First pass: compute last occurrence of each prime factor
	for i, x := range nums {
		primes := factorize(x)
		for _, p := range primes {
			last[p] = i
		}
	}

	// Second pass: track furthest last occurrence of primes seen so far
	furthest := 0
	for i, x := range nums {
		if i > furthest {
			return i - 1
		}
		primes := factorize(x)
		for _, p := range primes {
			if last[p] > furthest {
				furthest = last[p]
			}
		}
	}

	return -1
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: [2,3,4,5] ->", findValidSplit([]int{2, 3, 4, 5})) // 2

	// Additional test cases
	fmt.Println("Test 2: [4,7,15,8,3,5] ->", findValidSplit([]int{4, 7, 15, 8, 3, 5}))
	fmt.Println("Test 3: [1,1,1] ->", findValidSplit([]int{1, 1, 1})) // 0
	fmt.Println("Test 4: [6,10,15] ->", findValidSplit([]int{6, 10, 15}))
	fmt.Println("Test 5: [2,4,8] ->", findValidSplit([]int{2, 4, 8})) // -1
	fmt.Println("Test 6: [1] ->", findValidSplit([]int{1}))           // -1
}
