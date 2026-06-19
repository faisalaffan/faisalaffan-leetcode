package main

// LeetCode #3591: Check if Any Element Has Prime Frequency
// https://leetcode.com/problems/check-if-any-element-has-prime-frequency/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAnyElementHasPrimeFrequency([]int{1, 2, 3, 4, 5, 4}))
	fmt.Println(CheckIfAnyElementHasPrimeFrequency([]int{1, 2, 3, 4, 5}))
	fmt.Println(CheckIfAnyElementHasPrimeFrequency([]int{2, 2, 2, 4, 4}))
}

// Time: O(n + sqrt(m)) where m is max frequency
// Space: O(1)
func CheckIfAnyElementHasPrimeFrequency(nums []int) bool {
	freq := [101]int{}
	for _, v := range nums {
		freq[v]++
	}

	for _, f := range freq {
		if f > 1 && isPrime(f) {
			return true
		}
	}
	return false
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
