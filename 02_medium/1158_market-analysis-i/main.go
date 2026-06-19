package main

import (
	"fmt"
	"sort"
)

// LeetCode #1158: Market Analysis I
// https://leetcode.com/problems/market-analysis-i/
// Difficulty: Medium

// For each user, find total orders placed in 2019 and their join date.
// Sort by buyer_id.

// Time: O(n log n)
// Space: O(n)

type order struct {
	orderID    int
	orderDate  string
	productID  int
	buyerID    int
	sellerID   int
}

type user struct {
	userID    int
	joinDate  string
}

func marketAnalysisI(users []user, orders []order) [][]int {
	orderCount := make(map[int]int)
	for _, o := range orders {
		if len(o.orderDate) >= 4 && o.orderDate[:4] == "2019" {
			orderCount[o.buyerID]++
		}
	}

	result := make([][]int, 0, len(users))
	for _, u := range users {
		count := orderCount[u.userID]
		result = append(result, []int{u.userID, count})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

func main() {
	users := []user{
		{1, "2018-01-01"},
		{2, "2018-01-01"},
	}
	orders := []order{
		{1, "2019-01-01", 1, 1, 2},
		{2, "2020-01-01", 2, 1, 2},
		{3, "2019-05-01", 3, 2, 1},
	}
	fmt.Printf("%v (expected: [[1 1] [2 1]])\n", marketAnalysisI(users, orders))

	users2 := []user{
		{1, "2019-01-01"},
		{2, "2019-01-01"},
	}
	orders2 := []order{
		{1, "2020-01-01", 1, 1, 2},
		{2, "2020-05-01", 2, 2, 1},
	}
	fmt.Printf("%v (expected: [[1 0] [2 0]])\n", marketAnalysisI(users2, orders2))
}
