package main

// LeetCode #1271: Hexspeak
// https://leetcode.com/problems/hexspeak/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(log n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(toHexspeak("257"))  // "IOI"
	fmt.Println(toHexspeak("3"))    // "ERROR"
	fmt.Println(toHexspeak("619"))  // "ERROR" (619=26B, B not allowed)
}

// LeetCode submission: toHexspeak
func toHexspeak(num string) string {
	n, _ := strconv.Atoi(num)
	hex := strconv.FormatInt(int64(n), 16)
	replacer := map[byte]byte{
		'0': 'O',
		'1': 'I',
	}
	ans := make([]byte, len(hex))
	for i := range hex {
		if r, ok := replacer[hex[i]]; ok {
			ans[i] = r
		} else if hex[i] >= 'a' && hex[i] <= 'f' {
			ans[i] = hex[i] - 'a' + 'A'
		} else {
			return "ERROR"
		}
	}
	return string(ans)
}
