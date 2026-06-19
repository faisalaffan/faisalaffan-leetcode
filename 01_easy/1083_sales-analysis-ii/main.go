package main

// LeetCode #1083: Sales Analysis II
// https://leetcode.com/problems/sales-analysis-ii/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT DISTINCT s.buyer_id FROM Sales s JOIN Product p ON s.product_id = p.product_id WHERE p.product_name = 'S8' AND s.buyer_id NOT IN (SELECT s2.buyer_id FROM Sales s2 JOIN Product p2 ON s2.product_id = p2.product_id WHERE p2.product_name = 'iPhone')")
}

// This is a SQL problem. The answer is the SQL query above.
