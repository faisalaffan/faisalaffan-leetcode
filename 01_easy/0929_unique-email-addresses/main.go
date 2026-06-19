package main

// LeetCode #929: Unique Email Addresses
// https://leetcode.com/problems/unique-email-addresses/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})) // 2
	fmt.Println(numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))                                                    // 3
}

// numUniqueEmails counts unique email addresses after applying normalization rules.
// Time: O(n * m). Space: O(n).
func numUniqueEmails(emails []string) int {
	set := make(map[string]bool)
	for _, email := range emails {
		parts := strings.Split(email, "@")
		local := parts[0]
		domain := parts[1]

		// Remove dots and everything after '+'
		cleaned := strings.ReplaceAll(local, ".", "")
		if plusIdx := strings.Index(cleaned, "+"); plusIdx != -1 {
			cleaned = cleaned[:plusIdx]
		}
		set[cleaned+"@"+domain] = true
	}
	return len(set)
}
