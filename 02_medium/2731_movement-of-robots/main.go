package main

// LeetCode #2731: Movement of Robots
// https://leetcode.com/problems/movement-of-robots/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func MovementOfRobots(nums []int, s string, d int) int {
	n := len(nums)
	pos := make([]int, n)
	for i, num := range nums {
		pos[i] = num
		if s[i] == 'R' {
			pos[i] += d
		} else {
			pos[i] -= d
		}
	}

	sort.Ints(pos)

	const mod = 1_000_000_007
	var sum int64
	var prefix int64
	for i, p := range pos {
		sum = (sum + int64(p)*int64(i) - prefix) % mod
		prefix += int64(p)
	}

	return int(sum)
}

func main() {
	fmt.Println(MovementOfRobots([]int{1, 0}, "RL", 2))
	fmt.Println(MovementOfRobots([]int{-2, 0, 2}, "RLL", 3))
}
