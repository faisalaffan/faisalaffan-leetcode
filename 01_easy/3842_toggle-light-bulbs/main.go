package main

// LeetCode #3842: Toggle Light Bulbs
// https://leetcode.com/problems/toggle-light-bulbs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ToggleLightBulbs([]int{10, 30, 20, 10}))
	fmt.Println(ToggleLightBulbs([]int{100, 100}))
	fmt.Println(ToggleLightBulbs([]int{1, 2, 3}))
}

// Time: O(n)
// Space: O(1)
func ToggleLightBulbs(bulbs []int) []int {
	var state [101]int
	for _, b := range bulbs {
		state[b] ^= 1
	}
	var result []int
	for i := 1; i <= 100; i++ {
		if state[i] == 1 {
			result = append(result, i)
		}
	}
	return result
}
