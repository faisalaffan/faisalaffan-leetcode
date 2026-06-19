package main

// LeetCode #2910: Minimum Number of Groups to Create a Valid Assignment
// https://leetcode.com/problems/minimum-number-of-groups-to-create-a-valid-assignment/
// Difficulty: Medium
// Time: O(n + m*minFreq) | Space: O(m)

import "fmt"

func main() {
	fmt.Println(minGroupsForValidAssignment([]int{3, 3, 3, 3, 3, 1, 1}))
	fmt.Println(minGroupsForValidAssignment([]int{10, 10, 10, 10, 10}))
	fmt.Println(minGroupsForValidAssignment([]int{1, 1, 1, 2, 2, 2}))
}

func minGroupsForValidAssignment(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	k := len(nums)
	for _, v := range cnt {
		if v < k {
			k = v
		}
	}
	for ; ; k-- {
		ans := 0
		ok := true
		for _, v := range cnt {
			if v/k < v%k {
				ok = false
				break
			}
			ans += (v + k) / (k + 1)
		}
		if ok {
			return ans
		}
	}
}
