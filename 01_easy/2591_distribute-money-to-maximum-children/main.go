package main

// LeetCode #2591: Distribute Money to Maximum Children
// https://leetcode.com/problems/distribute-money-to-maximum-children/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DistributeMoneyToMaximumChildren(20, 3)) // 1
	fmt.Println(DistributeMoneyToMaximumChildren(16, 2)) // 2
}

func DistributeMoneyToMaximumChildren(money int, children int) int {
	if money < children {
		return -1
	}

	// Give each child 1 dollar first
	money -= children

	// Now we have 7-dollar increments (to make 8) for as many children as possible
	count := money / 7
	money %= 7

	// If count > children, we over-assigned
	if count > children {
		return children - 1
	}

	// If we have 3 children left and money = 3, we can't give it optimally
	remaining := children - count
	if remaining == 0 && money > 0 {
		count--
	} else if remaining == 1 && money == 3 {
		count--
	}

	return count
}
