package main

// LeetCode #970: Powerful Integers
// https://leetcode.com/problems/powerful-integers/
// Difficulty: Medium

import "fmt"

// Time: O(log_x(bound) * log_y(bound)) | Space: O(log_x(bound) * log_y(bound))
func powerfulIntegers(x int, y int, bound int) []int {
	seen := make(map[int]bool)

	for a := 1; a <= bound; a *= x {
		for b := 1; a+b <= bound; b *= y {
			seen[a+b] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}

	ans := make([]int, 0, len(seen))
	for v := range seen {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Println(powerfulIntegers(2, 3, 10))
	fmt.Println(powerfulIntegers(3, 5, 15))
	fmt.Println(powerfulIntegers(2, 1, 10))
}
