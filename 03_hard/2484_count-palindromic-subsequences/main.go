package main

// LeetCode #2484: Count Palindromic Subsequences
// https://leetcode.com/problems/count-palindromic-subsequences/
// Difficulty: Hard
//
// Count distinct palindromic subsequences of length 5 (a b c b a).
// For each middle position j and each digit pair (a,b):
//   leftPairs[a][b] = number of (a,b) ordered pairs before j
//   rightPairs[b][a] = number of (b,a) ordered pairs after j
//   result += leftPairs[a][b] * rightPairs[b][a]

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

	// Compute all pairs in the entire string
	totalPairs := [10][10]int64{}
	for i := 0; i < n; i++ {
		d := nums[i]
		for a := 0; a < 10; a++ {
			totalPairs[a][d] += int64(prefixCnt[i][a])
		}
	}

	// rightPairs starts as totalPairs, leftPairs starts as zeros
	rightPairs := [10][10]int64{}
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			rightPairs[a][b] = totalPairs[a][b]
		}
	}
	leftPairs := [10][10]int64{}

	var result int64 = 0

	for j := 0; j < n; j++ {
		d := nums[j]

		// Remove pairs involving position j
		for a := 0; a < 10; a++ {
			// pairs (a, d) where a is before j and d is at j — remove from right
			rightPairs[a][d] -= int64(prefixCnt[j][a])
		}

		// Now compute contributions with j as middle
		for a := 0; a < 10; a++ {
			for b := 0; b < 10; b++ {
				// leftPairs[a][b] = (a, b) pairs in prefix (before j)
				// rightPairs[b][a] = (b, a) pairs in suffix (after j)
				result = (result + leftPairs[a][b]*rightPairs[b][a]) % MOD
			}
		}

		// Add pairs where d is the second element (d is now in the prefix for next iterations)
		for a := 0; a < 10; a++ {
			leftPairs[a][d] += int64(prefixCnt[j][a])
		}
	}

	return int(result)
}
