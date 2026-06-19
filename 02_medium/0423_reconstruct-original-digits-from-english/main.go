package main

// LeetCode #423: Reconstruct Original Digits from English
// https://leetcode.com/problems/reconstruct-original-digits-from-english/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strings"
)

func originalDigits(s string) string {
	count := [26]int{}
	for _, ch := range s {
		count[ch-'a']++
	}

	// Unique identifying letters: z(0), w(2), u(4), x(6), g(8)
	digits := make([]int, 10)
	digits[0] = count['z'-'a']
	digits[2] = count['w'-'a']
	digits[4] = count['u'-'a']
	digits[6] = count['x'-'a']
	digits[8] = count['g'-'a']

	// Deduce remaining
	digits[1] = count['o'-'a'] - digits[0] - digits[2] - digits[4]
	digits[3] = count['h'-'a'] - digits[8]
	digits[5] = count['f'-'a'] - digits[4]
	digits[7] = count['s'-'a'] - digits[6]
	digits[9] = count['i'-'a'] - digits[5] - digits[6] - digits[8]

	var sb strings.Builder
	for d := 0; d <= 9; d++ {
		for i := 0; i < digits[d]; i++ {
			sb.WriteByte(byte('0' + d))
		}
	}
	return sb.String()
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", originalDigits("owoztneoer"))
	// Expected: "012"

	// Test case 2
	fmt.Println("Test 2:", originalDigits("fviefuro"))
	// Expected: "45"

	// Test case 3
	fmt.Println("Test 3:", originalDigits("zerozero"))
	// Expected: "00"
}
