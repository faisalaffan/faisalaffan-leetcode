package main

// LeetCode #1831: Maximum Transaction Each Day
// https://leetcode.com/problems/maximum-transaction-each-day/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

type Transaction struct {
	ID     int
	Day    int
	Amount int
}

func maxTransactionPerDay(transactions []Transaction) []int {
	// Group by day, find max amount
	dayMax := make(map[int]int)
	for _, t := range transactions {
		if t.Amount > dayMax[t.Day] {
			dayMax[t.Day] = t.Amount
		}
	}

	// Find transaction IDs that have max amount for their day
	result := make([]int, 0)
	for _, t := range transactions {
		if t.Amount == dayMax[t.Day] {
			result = append(result, t.ID)
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	transactions := []Transaction{
		{1, 1, 100},
		{2, 1, 200},
		{3, 2, 150},
		{4, 2, 100},
	}
	fmt.Println(maxTransactionPerDay(transactions)) // Expected: [2, 3]
}
