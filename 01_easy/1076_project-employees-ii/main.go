package main

// LeetCode #1076: Project Employees II
// https://leetcode.com/problems/project-employees-ii/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT project_id FROM Project GROUP BY project_id HAVING COUNT(employee_id) = (SELECT COUNT(employee_id) FROM Project GROUP BY project_id ORDER BY COUNT(employee_id) DESC LIMIT 1)")
}

// This is a SQL problem. The answer is the SQL query above.
