package main

// LeetCode #1082: Sales Analysis I
// https://leetcode.com/problems/sales-analysis-i/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT seller_id FROM Sales GROUP BY seller_id HAVING SUM(price) = (SELECT SUM(price) FROM Sales GROUP BY seller_id ORDER BY SUM(price) DESC LIMIT 1)")
}

// This is a SQL problem. The answer is the SQL query above.
