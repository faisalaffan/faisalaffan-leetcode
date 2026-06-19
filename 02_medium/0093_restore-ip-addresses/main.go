package main

// LeetCode #93: Restore IP Addresses
// https://leetcode.com/problems/restore-ip-addresses/
// Difficulty: Medium

import "fmt"

func restoreIpAddresses(s string) []string {
	result := []string{}
	if len(s) < 4 || len(s) > 12 {
		return result
	}

	var backtrack func(start int, parts []string)
	backtrack = func(start int, parts []string) {
		if len(parts) == 4 && start == len(s) {
			ip := parts[0] + "." + parts[1] + "." + parts[2] + "." + parts[3]
			result = append(result, ip)
			return
		}
		if len(parts) == 4 || start == len(s) {
			return
		}

		for i := 1; i <= 3 && start+i <= len(s); i++ {
			segment := s[start : start+i]
			if (len(segment) > 1 && segment[0] == '0') || (i == 3 && segment > "255") {
				continue
			}
			parts = append(parts, segment)
			backtrack(start+i, parts)
			parts = parts[:len(parts)-1]
		}
	}

	backtrack(0, []string{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(restoreIpAddresses("25525511135"))
	// ["255.255.11.135","255.255.111.35"]

	// Test case 2
	fmt.Println(restoreIpAddresses("0000")) // ["0.0.0.0"]

	// Test case 3
	fmt.Println(restoreIpAddresses("101023"))
	// ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
}

// Time: O(3^4) = O(1) | Space: O(1)
