package main

// LeetCode #3896: Minimum Operations to Transform Array into Alternating Prime
// https://leetcode.com/problems/minimum-operations-to-transform-array-into-alternating-prime/
// Difficulty: Medium
// Time: O(N log log M + N) | Space: O(M) where M = 200000
// Approach: Sieve primes. Even indices need next prime >= val (binary search).
// Odd indices need non-prime: if prime, inc to next non-prime (2->4 cost 2, rest cost 1).

import "fmt"

const MAX_VAL = 200000

func MinimumOperationsToTransformArrayIntoAlternatingPrime(nums []int) int64 {
	isPrime := make([]bool, MAX_VAL+1)
	for i := 2; i <= MAX_VAL; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= MAX_VAL; i++ {
		if isPrime[i] {
			for j := i * i; j <= MAX_VAL; j += i {
				isPrime[j] = false
			}
		}
	}

	// nextPrime[v] = smallest prime >= v
	nextPrime := make([]int, MAX_VAL+1)
	np := -1
	for i := MAX_VAL; i >= 2; i-- {
		if isPrime[i] {
			np = i
		}
		nextPrime[i] = np
	}

	var ans int64 = 0
	for i, v := range nums {
		if i%2 == 0 {
			// Even index: must be prime
			need := nextPrime[v]
			if need == -1 {
				// No prime >= v within MAX_VAL (shouldn't happen per constraints)
				ans += int64(MAX_VAL + 1 - v)
			} else {
				ans += int64(need - v)
			}
		} else {
			// Odd index: must be non-prime
			if isPrime[v] {
				if v == 2 {
					ans += 2 // 2 -> 4
				} else {
					ans += 1 // odd prime +1 -> even non-prime
				}
			}
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumOperationsToTransformArrayIntoAlternatingPrime([]int{1, 2, 3, 4})) // Expected: 3

	// Example 2
	fmt.Println(MinimumOperationsToTransformArrayIntoAlternatingPrime([]int{5, 6, 7, 8})) // Expected: 0

	// Example 3
	fmt.Println(MinimumOperationsToTransformArrayIntoAlternatingPrime([]int{4, 4})) // Expected: 1
}
