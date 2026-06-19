package main

// LeetCode #2979: Most Expensive Item That Can Not Be Bought
// https://leetcode.com/problems/most-expensive-item-that-can-not-be-bought/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(mostExpensiveItem(2, 5))
	fmt.Println(mostExpensiveItem(3, 7))
}

func mostExpensiveItem(primeOne int, primeTwo int) int {
	return primeOne*primeTwo - primeOne - primeTwo
}
