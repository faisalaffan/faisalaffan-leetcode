package main

// LeetCode #670: Maximum Swap
// https://leetcode.com/problems/maximum-swap/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximumSwap(2736))
	fmt.Println(maximumSwap(9973))
	fmt.Println(maximumSwap(98368))
}

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)

	// Track last occurrence of each digit
	last := make([]int, 10)
	for i := 0; i < n; i++ {
		last[s[i]-'0'] = i
	}

	for i := 0; i < n; i++ {
		for d := 9; d > int(s[i]-'0'); d-- {
			if last[d] > i {
				s[i], s[last[d]] = s[last[d]], s[i]
				result, _ := strconv.Atoi(string(s))
				return result
			}
		}
	}

	return num
}
