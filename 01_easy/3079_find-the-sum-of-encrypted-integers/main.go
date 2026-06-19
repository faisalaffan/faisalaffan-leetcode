package main

// LeetCode #3079: Find the Sum of Encrypted Integers
// https://leetcode.com/problems/find-the-sum-of-encrypted-integers/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: sumOfEncryptedInt
	fmt.Println(FindTheSumOfEncryptedIntegers([]int{10, 21, 31})) // 66
	fmt.Println(FindTheSumOfEncryptedIntegers([]int{1, 2, 3}))    // 6
}

// Time: O(n * d) where d is number of digits | Space: O(1)
// LeetCode submission name: sumOfEncryptedInt
func FindTheSumOfEncryptedIntegers(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += encrypt(v)
	}
	return sum
}

func encrypt(n int) int {
	maxDigit := 0
	digits := 0
	for n > 0 {
		d := n % 10
		if d > maxDigit {
			maxDigit = d
		}
		digits++
		n /= 10
	}
	result := 0
	for i := 0; i < digits; i++ {
		result = result*10 + maxDigit
	}
	return result
}
