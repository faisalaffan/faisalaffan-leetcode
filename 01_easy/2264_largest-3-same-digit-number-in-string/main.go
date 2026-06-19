package main

// LeetCode #2264: Largest 3-Same-Digit Number in String
// https://leetcode.com/problems/largest-3-same-digit-number-in-string/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(LargestThreeSameDigitNumberInString("6777133339")) // "777"
	fmt.Println(LargestThreeSameDigitNumberInString("2300011114")) // "000"
	fmt.Println(LargestThreeSameDigitNumberInString("42352338"))   // ""
}

func LargestThreeSameDigitNumberInString(num string) string {
	best := byte(0)
	for i := 2; i < len(num); i++ {
		if num[i] == num[i-1] && num[i] == num[i-2] && num[i] > best {
			best = num[i]
		}
	}
	if best == 0 {
		return ""
	}
	return string([]byte{best, best, best})
}
