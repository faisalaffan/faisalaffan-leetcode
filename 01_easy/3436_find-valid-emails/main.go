package main

// LeetCode #3436: Find Valid Emails
// https://leetcode.com/problems/find-valid-emails/
// Difficulty: Easy

import (
	"fmt"
	"regexp"
)

func main() {
	users := []User{
		{UserID: 1, Email: "alice@example.com"},
		{UserID: 2, Email: "bob@example"},
		{UserID: 3, Email: "@example.com"},
		{UserID: 4, Email: "charlie@example.com"},
	}
	result := FindValidEmails(users)
	for _, r := range result {
		fmt.Printf("%d: %s\n", r.UserID, r.Email)
	}
}

// User represents a user with an email.
type User struct {
	UserID int
	Email  string
}

// FindValidEmails returns users with valid email addresses (alphanumeric prefix, letter-only domain, .com suffix).
// Time: O(n). Space: O(n).
func FindValidEmails(users []User) []User {
	re := regexp.MustCompile(`^[A-Za-z0-9_]+@[A-Za-z]+\.com$`)
	result := []User{}
	for _, u := range users {
		if re.MatchString(u.Email) {
			result = append(result, u)
		}
	}
	return result
}
