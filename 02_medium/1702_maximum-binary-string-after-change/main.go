package main

// LeetCode #1702: Maximum Binary String After Change
// https://leetcode.com/problems/maximum-binary-string-after-change/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maximumBinaryString(binary string) string {
	n := len(binary)
	zeros := 0
	firstZero := -1

	for i, ch := range binary {
		if ch == '0' {
			zeros++
			if firstZero == -1 {
				firstZero = i
			}
		}
	}

	if zeros <= 1 {
		return binary
	}

	// Result: all 1s except position (firstZero + zeros - 1)
	result := make([]byte, n)
	for i := range result {
		if i == firstZero+zeros-1 {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}
	return string(result)
}

func main() {
	fmt.Println(maximumBinaryString("000110")) // Expected: "111011"
	fmt.Println(maximumBinaryString("01"))     // Expected: "01"
	fmt.Println(maximumBinaryString("10"))     // Expected: "10"
}
