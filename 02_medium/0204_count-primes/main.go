package main

// LeetCode #204: Count Primes
// https://leetcode.com/problems/count-primes/
// Difficulty: Medium
// Time: O(n log log n), Space: O(n)

import "fmt"

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(0))
	fmt.Println(countPrimes(1))
}
