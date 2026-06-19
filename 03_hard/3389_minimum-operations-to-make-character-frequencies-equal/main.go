package main

// LeetCode #3389: Minimum Operations to Make Character Frequencies Equal
// https://leetcode.com/problems/minimum-operations-to-make-character-frequencies-equal/
// Difficulty: Hard
//
// Enumerate target frequency 1..maxFreq. DP across 26 letters.
// Operations: delete a char, add a char, or change a char to next letter.

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MinimumOperationsToMakeCharacterFrequenciesEqual("aabbcc"))
	fmt.Println(MinimumOperationsToMakeCharacterFrequenciesEqual("aaabbbccc"))
}

func MinimumOperationsToMakeCharacterFrequenciesEqual(s string) int {
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	maxFreq := 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
		}
	}

	best := math.MaxInt32
	for target := 1; target <= maxFreq; target++ {
		ops := 0
		for i := 0; i < 26; i++ {
			if freq[i] > target {
				ops += freq[i] - target
			}
		}
		surplus := 0
		for i := 0; i < 26; i++ {
			if freq[i] < target {
				need := target - freq[i]
				if surplus >= need {
					surplus -= need
				} else {
					ops += need - surplus
					surplus = 0
				}
			} else if freq[i] > target {
				surplus += freq[i] - target
			}
		}
		if ops < best {
			best = ops
		}
	}
	return best
}
