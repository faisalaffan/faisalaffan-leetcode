package main

// LeetCode #1742: Maximum Number of Balls in a Box
// https://leetcode.com/problems/maximum-number-of-balls-in-a-box/
// Difficulty: Easy

import "fmt"

// Time: O(n log n), Space: O(n) - but n <= 10^5, fine
func CountBalls(lowLimit int, highLimit int) int {
	boxes := make(map[int]int)
	maxBalls := 0
	for i := lowLimit; i <= highLimit; i++ {
		sum := digitSum(i)
		boxes[sum]++
		if boxes[sum] > maxBalls {
			maxBalls = boxes[sum]
		}
	}
	return maxBalls
}

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(CountBalls(1, 10))
	fmt.Println(CountBalls(5, 15))
	fmt.Println(CountBalls(19, 28))
}
