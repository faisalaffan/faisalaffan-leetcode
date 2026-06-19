package main

// LeetCode #1281: Subtract the Product and Sum of Digits of an Integer
// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
// Difficulty: Easy
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(234))  // 15
	fmt.Println(subtractProductAndSum(4421)) // 21
}

// LeetCode submission: subtractProductAndSum
func subtractProductAndSum(n int) int {
	product := 1
	sum := 0
	for x := n; x > 0; x /= 10 {
		d := x % 10
		product *= d
		sum += d
	}
	return product - sum
}
