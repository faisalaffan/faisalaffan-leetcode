package main

// LeetCode #2806: Account Balance After Rounded Purchase
// https://leetcode.com/problems/account-balance-after-rounded-purchase/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(AccountBalanceAfterRoundedPurchase(9))
	fmt.Println(AccountBalanceAfterRoundedPurchase(15))
}

func AccountBalanceAfterRoundedPurchase(purchaseAmount int) int {
	return 100 - ((purchaseAmount+5)/10)*10
}
