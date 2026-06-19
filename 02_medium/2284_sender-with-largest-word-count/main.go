package main

// LeetCode #2284: Sender With Largest Word Count
// https://leetcode.com/problems/sender-with-largest-word-count/
// Difficulty: Medium
// Time: O(n * m) | Space: O(n)

import (
	"fmt"
	"strings"
)

func largestWordCount(messages []string, senders []string) string {
	count := make(map[string]int)
	maxCount := 0
	maxSender := ""

	for i, msg := range messages {
		sender := senders[i]
		words := len(strings.Fields(msg))
		count[sender] += words
		if count[sender] > maxCount || (count[sender] == maxCount && sender > maxSender) {
			maxCount = count[sender]
			maxSender = sender
		}
	}
	return maxSender
}

func main() {
	// Test case 1
	fmt.Println(largestWordCount([]string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"}, []string{"Alice", "userTwo", "userThree", "Alice"}))
	// Expected: "Alice"

	// Test case 2
	fmt.Println(largestWordCount([]string{"t", "e", "s", "t"}, []string{"a", "b", "c", "d"}))
	// Expected: "d" (first with max when equal, by lexicographical)
}
