package main

// LeetCode #1421: NPV Queries
// https://leetcode.com/problems/npv-queries/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: NPV (id, year, npv), Queries (id, year)

import "fmt"

func main() {
	fmt.Println(NpvQueries())
}

// Time: N/A (SQL query), Space: N/A
func NpvQueries() string {
	return `SELECT q.id, q.year, IFNULL(n.npv, 0) AS npv
FROM Queries q
LEFT JOIN NPV n ON q.id = n.id AND q.year = n.year
ORDER BY q.id, q.year;`
}
