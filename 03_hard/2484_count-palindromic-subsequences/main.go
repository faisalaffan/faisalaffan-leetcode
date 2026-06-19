package main

// LeetCode #2484: Count Palindromic Subsequences
// https://leetcode.com/problems/count-palindromic-subsequences/
// Difficulty: Hard
//
// Count distinct palindromic subsequences of length 5 (a b c b a).
// For each middle position j and each pair (a,b):
//   leftCount[a][b] = number of (a,b) ordered pairs in s[0:j]
//   rightCount[a][b] = number of (a,b) ordered pairs in s[j+1:n]
//   result += leftCount[a][b] * rightCount[b][a]

import "fmt"

const MOD = 1000000007

func main() {
	// Example 1: "103301" => 2
	fmt.Println(countPalindromicSubsequences("103301"))
	// Example 2: "0000000" => 21
	fmt.Println(countPalindromicSubsequences("0000000"))
	// Example 3: "9999900000" => 96
	fmt.Println(countPalindromicSubsequences("9999900000"))
	// Edge: length 5 palindrome
	fmt.Println(countPalindromicSubsequences("12321"))
	// Edge: all distinct
	fmt.Println(countPalindromicSubsequences("12345"))
}

func countPalindromicSubsequences(s string) int {
	n := len(s)
	if n < 5 {
		return 0
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = int(s[i] - '0')
	}

	// prefix counts
	prefixCnt := make([][10]int, n+1)
	for i := 0; i < n; i++ {
		for d := 0; d < 10; d++ {
			prefixCnt[i+1][d] = prefixCnt[i][d]
		}
		prefixCnt[i+1][nums[i]]++
	}

	// suffix counts
	suffixCnt := make([][10]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for d := 0; d < 10; d++ {
			suffixCnt[i][d] = suffixCnt[i+1][d]
		}
		suffixCnt[i][nums[i]]++
	}

	// Compute total ordered pairs in the full string
	// rightPairs[a][b] = number of (a,b) pairs where a appears before b
	rightPairs := [10][10]int64{}
	for i := 0; i < n; i++ {
		d := nums[i]
		for a := 0; a < 10; a++ {
			rightPairs[a][d] += int64(prefixCnt[i][a])
		}
	}

	leftPairs := [10][10]int64{}
	var result int64 = 0

	for j := 0; j < n; j++ {
		d := nums[j]

		// Remove pairs involving position j from rightPairs
		// Pairs (a, d) where a is before j: these are no longer "after j"
		for a := 0; a < 10; a++ {
			rightPairs[a][d] -= int64(prefixCnt[j][a])
		}
		// Pairs (d, b) where b is after j: remove them too
		for b := 0; b < 10; b++ {
			rightPairs[d][b] -= int64(suffixCnt[j+1][b])
		}

		// For each (a,b), count palindromes with j as middle
		for a := 0; a < 10; a++ {
			for b := 0; b < 10; b++ {
				result = (result + leftPairs[a][b]*rightPairs[b][a]) % MOD
			}
		}

		// Add pairs (a, d) where d is the second element, now part of left for next iterations
		for a := 0; a < 10; a++ {
			leftPairs[a][d] += int64(prefixCnt[j][a])
		}
	}

	return int(result)
}
