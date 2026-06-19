package main

// LeetCode #1148: Article Views I
// https://leetcode.com/problems/article-views-i/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT DISTINCT author_id AS id FROM Views WHERE author_id = viewer_id ORDER BY author_id")
}

// This is a SQL problem. The answer is the SQL query above.
