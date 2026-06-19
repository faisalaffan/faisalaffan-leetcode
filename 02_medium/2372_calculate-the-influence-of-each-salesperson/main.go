package main

// LeetCode #2372: Calculate the Influence of Each Salesperson
// https://leetcode.com/problems/calculate-the-influence-of-each-salesperson/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)
// Simulate: track total price by salesperson from sales and product tables.

import "fmt"

func main() {
	sales := [][]int{
		{1, 1, 10}, // salesperson_id, product_id, quantity
		{2, 1, 5},
		{1, 2, 8},
		{3, 2, 3},
		{2, 2, 2},
	}
	products := map[int]int{
		1: 100, // product_id -> price
		2: 50,
	}
	fmt.Println(calculateInfluence(sales, products)) // {1: 1400, 2: 600, 3: 150}

	sales2 := [][]int{
		{1, 1, 1},
		{2, 1, 1},
	}
	products2 := map[int]int{1: 500}
	fmt.Println(calculateInfluence(sales2, products2)) // {1: 500, 2: 500}
}

func calculateInfluence(sales [][]int, priceMap map[int]int) map[int]int {
	result := make(map[int]int)
	for _, s := range sales {
		sp, prodID, qty := s[0], s[1], s[2]
		result[sp] += qty * priceMap[prodID]
	}
	return result
}
