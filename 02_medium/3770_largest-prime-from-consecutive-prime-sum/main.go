package main

// LeetCode #3770: Largest Prime from Consecutive Prime Sum
// https://leetcode.com/problems/largest-prime-from-consecutive-prime-sum/
// Difficulty: Medium
// Time: O(n log log n) | Space: O(n)

import "fmt"

func largestPrimeFromConsecutivePrimeSum(n int) int {
	if n < 2 {
		return 0
	}
	// Sieve up to n
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Collect primes
	primes := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	// Prefix sums
	pref := make([]int, len(primes)+1)
	for i, p := range primes {
		pref[i+1] = pref[i] + p
	}

	ans := 0
	for i := 0; i < len(primes); i++ {
		for j := i; j < len(primes); j++ {
			sum := pref[j+1] - pref[i]
			if sum > n {
				break
			}
			if isPrime[sum] && sum > ans {
				ans = sum
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(largestPrimeFromConsecutivePrimeSum(20))
	fmt.Println(largestPrimeFromConsecutivePrimeSum(2))
	fmt.Println(largestPrimeFromConsecutivePrimeSum(50))
}
