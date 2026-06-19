package main

// LeetCode #1303: Find the Team Size
// https://leetcode.com/problems/find-the-team-size/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT employee_id, COUNT(team_id) OVER (PARTITION BY team_id) AS team_size FROM Employee")
}

// This is a SQL problem. The answer is the SQL query above.
