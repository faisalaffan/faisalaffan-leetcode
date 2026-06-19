package main

// LeetCode #1881: Maximum Value After Insertion
// https://leetcode.com/problems/maximum-value-after-insertion/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxValue("99", 9))
	fmt.Println(MaxValue("-13", 2))
	fmt.Println(MaxValue("73", 6))
}

// Time: O(n), Space: O(n)
func MaxValue(n string, x int) string {
	result := make([]byte, 0, len(n)+1)
	negative := n[0] == '-'

	if negative {
		result = append(result, '-')
		inserted := false
		for i := 1; i < len(n); i++ {
			digit := int(n[i] - '0')
			if !inserted && x < digit {
				result = append(result, byte(x+'0'))
				inserted = true
			}
			result = append(result, n[i])
		}
		if !inserted {
			result = append(result, byte(x+'0'))
		}
	} else {
		inserted := false
		for i := 0; i < len(n); i++ {
			digit := int(n[i] - '0')
			if !inserted && x > digit {
				result = append(result, byte(x+'0'))
				inserted = true
			}
			result = append(result, n[i])
		}
		if !inserted {
			result = append(result, byte(x+'0'))
		}
	}
	return string(result)
}
