package main

// LeetCode #1211: Queries Quality and Percentage
// https://leetcode.com/problems/queries-quality-and-percentage/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT query_name, ROUND(AVG(rating/position), 2) AS quality, ROUND(100.0 * SUM(CASE WHEN rating < 3 THEN 1 ELSE 0 END) / COUNT(*), 2) AS poor_query_percentage FROM Queries GROUP BY query_name")
}

// This is a SQL problem. The answer is the SQL query above.
