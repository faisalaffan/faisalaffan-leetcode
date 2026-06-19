package main

// LeetCode #299: Bulls and Cows
// https://leetcode.com/problems/bulls-and-cows/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import (
	"fmt"
	"strconv"
)

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	secretCount := [10]int{}
	guessCount := [10]int{}

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretCount[secret[i]-'0']++
			guessCount[guess[i]-'0']++
		}
	}

	for i := 0; i < 10; i++ {
		if secretCount[i] < guessCount[i] {
			cows += secretCount[i]
		} else {
			cows += guessCount[i]
		}
	}

	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	fmt.Println(getHint("1", "0"))
}
