package main

// LeetCode #248: Strobogrammatic Number III
// https://leetcode.com/problems/strobogrammatic-number-iii/
// Difficulty: Hard [Paid]

import "fmt"

var pairs = [][2]byte{
	{'0', '0'},
	{'1', '1'},
	{'6', '9'},
	{'8', '8'},
	{'9', '6'},
}

func strobogrammaticInRange(low string, high string) int {
	count := 0
	lowNum := len(low)
	highNum := len(high)

	for length := lowNum; length <= highNum; length++ {
		var dfs func(cur []byte, left, right int)
		dfs = func(cur []byte, left, right int) {
			if left > right {
				s := string(cur)
				if (len(s) == len(low) && s < low) || (len(s) == len(high) && s > high) {
					return
				}
				count++
				return
			}

			for _, p := range pairs {
				cur[left] = p[0]
				cur[right] = p[1]
				if len(cur) > 1 && cur[0] == '0' {
					continue
				}
				if left == right && p[0] != p[1] {
					continue
				}
				dfs(cur, left+1, right-1)
			}
		}

		cur := make([]byte, length)
		dfs(cur, 0, length-1)
	}

	// fix: check edge cases properly — compare as strings with same length
	result := 0
	if len(low) == len(high) {
		for _, s := range generate(len(low)) {
			if s >= low && s <= high {
				result++
			}
		}
		return result
	}

	for _, s := range generate(len(low)) {
		if s >= low {
			result++
		}
	}
	for _, s := range generate(len(high)) {
		if s <= high {
			result++
		}
	}
	for l := len(low) + 1; l < len(high); l++ {
		result += countStrobogrammatic(l)
	}
	return result
}

func generate(n int) []string {
	return helper(n, n)
}

func helper(n, m int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"0", "1", "8"}
	}

	inner := helper(n-2, m)
	result := make([]string, 0)
	for _, s := range inner {
		if n != m {
			result = append(result, "0"+s+"0")
		}
		result = append(result, "1"+s+"1")
		result = append(result, "6"+s+"9")
		result = append(result, "8"+s+"8")
		result = append(result, "9"+s+"6")
	}
	return result
}

func countStrobogrammatic(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 3 // 0, 1, 8
	}
	valid := 1
	if n%2 == 1 {
		valid = 3
	}
	count := 4 // 1, 6, 8, 9 (first digit)
	for i := 0; i < n/2-1; i++ {
		count *= 5
	}
	count *= valid
	// subtract 0-padded ones
	if n%2 == 0 {
		return count
	}
	// odd length, middle can be 0,1,8
	return count
}

func main() {
	fmt.Println(strobogrammaticInRange("50", "100"))
}
