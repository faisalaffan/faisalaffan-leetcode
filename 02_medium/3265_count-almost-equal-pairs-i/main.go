package main

// LeetCode #3265: Count Almost Equal Pairs I
// https://leetcode.com/problems/count-almost-equal-pairs-i/
// Difficulty: Medium
// Time: O(n^2 * d) Space: O(d) where d = number of digits

import "fmt"

func main() {
	fmt.Println(countPairs([]int{1, 1, 1}))                                           // 3
	fmt.Println(countPairs([]int{3, 12, 30, 17, 21}))                                 // 2
	fmt.Println(countPairs([]int{1023, 3012, 1230, 2301, 0, 0}))                     // 4
}

func countPairs(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if isAlmostEqual(nums[i], nums[j]) {
				ans++
			}
		}
	}
	return ans
}

func isAlmostEqual(a, b int) bool {
	sa, sb := fmt.Sprintf("%07d", a), fmt.Sprintf("%07d", b)
	diff := 0
	ca, cb := make([]int, 10), make([]int, 10)
	for k := 0; k < 7; k++ {
		if sa[k] != sb[k] {
			diff++
			if diff > 2 {
				return false
			}
		}
		ca[sa[k]-'0']++
		cb[sb[k]-'0']++
	}
	for k := 0; k < 10; k++ {
		if ca[k] != cb[k] {
			return false
		}
	}
	return diff <= 2
}
