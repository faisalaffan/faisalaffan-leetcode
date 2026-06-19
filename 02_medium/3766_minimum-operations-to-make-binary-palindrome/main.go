package main

// LeetCode #3766: Minimum Operations to Make Binary Palindrome
// https://leetcode.com/problems/minimum-operations-to-make-binary-palindrome/
// Difficulty: Medium
// Time: O(n * log M) | Space: O(M)

import (
	"fmt"
	"sort"
	"strconv"
)

var binaryPalindromes []int

func init() {
	for i := 0; i < (1 << 14); i++ {
		s := strconv.FormatInt(int64(i), 2)
		if isPalindromeStr(s) {
			binaryPalindromes = append(binaryPalindromes, i)
		}
	}
}

func isPalindromeStr(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func minimumOperationsToMakeBinaryPalindrome(nums []int) []int {
	ans := make([]int, len(nums))
	for idx, x := range nums {
		pos := sort.SearchInts(binaryPalindromes, x)
		best := 1 << 30
		if pos < len(binaryPalindromes) {
			if binaryPalindromes[pos]-x < best {
				best = binaryPalindromes[pos] - x
			}
		}
		if pos > 0 {
			if x-binaryPalindromes[pos-1] < best {
				best = x - binaryPalindromes[pos-1]
			}
		}
		ans[idx] = best
	}
	return ans
}

func main() {
	fmt.Println(minimumOperationsToMakeBinaryPalindrome([]int{1, 2, 3, 4, 5}))
	fmt.Println(minimumOperationsToMakeBinaryPalindrome([]int{10, 20}))
	fmt.Println(minimumOperationsToMakeBinaryPalindrome([]int{7}))
}
