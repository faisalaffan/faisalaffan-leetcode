package main

// LeetCode #1869: Longer Contiguous Segments of Ones Than Zeros
// https://leetcode.com/problems/longer-contiguous-segments-of-ones-than-zeros/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CheckZeroOnes(s string) bool {
	maxOnes, maxZeros := 0, 0
	curOnes, curZeros := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			curOnes++
			curZeros = 0
			if curOnes > maxOnes {
				maxOnes = curOnes
			}
		} else {
			curZeros++
			curOnes = 0
			if curZeros > maxZeros {
				maxZeros = curZeros
			}
		}
	}
	return maxOnes > maxZeros
}

func main() {
	fmt.Println(CheckZeroOnes("1101"))
	fmt.Println(CheckZeroOnes("111000"))
	fmt.Println(CheckZeroOnes("110100010"))
}
