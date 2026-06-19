package main

// LeetCode #182: Duplicate Emails
// https://leetcode.com/problems/duplicate-emails/
// Difficulty: Easy

import "fmt"

func DuplicateEmails() string {
	return "SELECT email FROM Person GROUP BY email HAVING COUNT(email) > 1"
}

func main() {
	fmt.Println(DuplicateEmails())
}
