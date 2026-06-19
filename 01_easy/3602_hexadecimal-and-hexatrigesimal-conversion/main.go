package main

// LeetCode #3602: Hexadecimal and Hexatrigesimal Conversion
// https://leetcode.com/problems/hexadecimal-and-hexatrigesimal-conversion/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

const digits = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	fmt.Println(HexadecimalAndHexatrigesimalConversion(4))
}

// Time: O(log n)
// Space: O(log n)
func HexadecimalAndHexatrigesimalConversion(n int) string {
	sq := n * n
	cube := n * n * n
	return toBase(sq, 16) + toBase(cube, 36)
}

func toBase(num, base int) string {
	if num == 0 {
		return "0"
	}
	var res strings.Builder
	for num > 0 {
		res.WriteByte(digits[num%base])
		num /= base
	}

	// reverse
	s := []byte(res.String())
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
