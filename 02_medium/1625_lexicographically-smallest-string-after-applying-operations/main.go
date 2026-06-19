package main

// LeetCode #1625: Lexicographically Smallest String After Applying Operations
// https://leetcode.com/problems/lexicographically-smallest-string-after-applying-operations/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindLexSmallestString("5525", 9, 2))
	fmt.Println(FindLexSmallestString("74", 5, 1))
	fmt.Println(FindLexSmallestString("0011", 4, 2))
}

func FindLexSmallestString(s string, a int, b int) string {
	// Time: O(N^2), Space: O(N)
	// BFS over all possible states
	n := len(s)
	seen := make(map[string]bool)
	queue := []string{s}
	seen[s] = true
	smallest := s

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr < smallest {
			smallest = curr
		}

		// Operation 1: add a to odd-position digits
		bytes := []byte(curr)
		for i := 1; i < len(bytes); i += 2 {
			val := int(bytes[i] - '0')
			val = (val + a) % 10
			bytes[i] = byte('0' + val)
		}
		next1 := string(bytes)

		// Operation 2: rotate right by b
		next2 := curr[n-b:] + curr[:n-b]

		for _, next := range []string{next1, next2} {
			if !seen[next] {
				seen[next] = true
				queue = append(queue, next)
			}
		}
	}

	return smallest
}
