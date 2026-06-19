package main

// LeetCode #3564: Seasonal Sales Analysis
// https://leetcode.com/problems/seasonal-sales-analysis/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	sales := []int{100, 200, 150, 300, 250}
	fmt.Println("Test 1:", SeasonalSalesAnalysis(sales))
	// Test case 2
	sales2 := []int{50, 60, 70, 80}
	fmt.Println("Test 2:", SeasonalSalesAnalysis(sales2))
	// Test case 3
	sales3 := []int{100}
	fmt.Println("Test 3:", SeasonalSalesAnalysis(sales3))
}

func SeasonalSalesAnalysis(sales []int) int {
	if len(sales) <= 1 {
		return 0
	}
	// Find max difference between any two sales that is positive
	minPrice := sales[0]
	maxProfit := 0
	for i := 1; i < len(sales); i++ {
		if sales[i]-minPrice > maxProfit {
			maxProfit = sales[i] - minPrice
		}
		if sales[i] < minPrice {
			minPrice = sales[i]
		}
	}
	return maxProfit
}
