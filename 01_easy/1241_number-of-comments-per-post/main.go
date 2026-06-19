package main

// LeetCode #1241: Number of Comments per Post
// https://leetcode.com/problems/number-of-comments-per-post/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT DISTINCT p.sub_id AS post_id, (SELECT COUNT(DISTINCT c.sub_id) FROM Submissions c WHERE c.parent_id = p.sub_id) AS number_of_comments FROM Submissions p WHERE p.parent_id IS NULL ORDER BY p.sub_id")
}

// This is a SQL problem. The answer is the SQL query above.
