package main

// LeetCode #313: Super Ugly Number
// https://leetcode.com/problems/super-ugly-number/
// Difficulty: Medium
// Time: O(n * len(primes)), Space: O(n + len(primes))

import "fmt"

func nthSuperUglyNumber(n int, primes []int) int {
	ugly := make([]int, n)
	ugly[0] = 1

	pointers := make([]int, len(primes))
	values := make([]int, len(primes))
	for i, p := range primes {
		values[i] = p
	}

	for i := 1; i < n; i++ {
		minVal := values[0]
		for _, v := range values {
			if v < minVal {
				minVal = v
			}
		}
		ugly[i] = minVal

		for j := range values {
			if values[j] == minVal {
				pointers[j]++
				values[j] = ugly[pointers[j]] * primes[j]
			}
		}
	}

	return ugly[n-1]
}

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
	fmt.Println(nthSuperUglyNumber(1, []int{2, 3, 5}))
	fmt.Println(nthSuperUglyNumber(6, []int{2, 3, 5}))
}
