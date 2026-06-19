package main

// LeetCode #2520: Count the Digits That Divide a Number
// https://leetcode.com/problems/count-the-digits-that-divide-a-number/
// Difficulty: Easy
// Time O(log n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountTheDigitsThatDivideANumber(7))    // 1
	fmt.Println(CountTheDigitsThatDivideANumber(121))  // 2
	fmt.Println(CountTheDigitsThatDivideANumber(1248)) // 4
}

func CountTheDigitsThatDivideANumber(num int) int {
	count := 0
	n := num
	for n > 0 {
		digit := n % 10
		if digit != 0 && num%digit == 0 {
			count++
		}
		n /= 10
	}
	return count
}
