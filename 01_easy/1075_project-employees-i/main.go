package main

// LeetCode #1075: Project Employees I
// https://leetcode.com/problems/project-employees-i/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT p.project_id, ROUND(AVG(e.experience_years), 2) AS average_years FROM Project p JOIN Employee e ON p.employee_id = e.employee_id GROUP BY p.project_id")
}

// This is a SQL problem. The answer is the SQL query above.
