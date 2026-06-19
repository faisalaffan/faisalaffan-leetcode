package main

// LeetCode #2844: Minimum Operations to Make a Special Number
// https://leetcode.com/problems/minimum-operations-to-make-a-special-number/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math"
)

func MinimumOperationsToMakeASpecialNumber(num string) int {
	n := len(num)
	best := n // Remove all

	// Find "00"
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if num[i] == '0' && num[j] == '0' {
				ops := (n - 1 - j) + (j - i - 1)
				if ops < best {
					best = ops
				}
			}
			if num[i] == '2' && num[j] == '5' {
				ops := (n - 1 - j) + (j - i - 1)
				if ops < best {
					best = ops
				}
			}
			if num[i] == '5' && num[j] == '0' {
				ops := (n - 1 - j) + (j - i - 1)
				if ops < best {
					best = ops
				}
			}
			if num[i] == '7' && num[j] == '5' {
				ops := (n - 1 - j) + (j - i - 1)
				if ops < best {
					best = ops
				}
			}
		}
	}

	// Also check for single "0"
	for i := 0; i < n; i++ {
		if num[i] == '0' {
			ops := n - 1
			if ops < best {
				best = ops
			}
		}
	}

	if best == math.MaxInt32 {
		return n
	}
	return best
}

func main() {
	fmt.Println(MinimumOperationsToMakeASpecialNumber("2245047"))
	fmt.Println(MinimumOperationsToMakeASpecialNumber("2908305"))
	fmt.Println(MinimumOperationsToMakeASpecialNumber("10"))
}
