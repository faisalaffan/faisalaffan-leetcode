package main

// LeetCode #3259: Maximum Energy Boost From Two Drinks
// https://leetcode.com/problems/maximum-energy-boost-from-two-drinks/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxEnergyBoost([]int{1, 3, 1}, []int{3, 1, 1}))       // 5
	fmt.Println(maxEnergyBoost([]int{4, 1, 1}, []int{1, 1, 3}))       // 7
	fmt.Println(maxEnergyBoost([]int{2, 2, 2, 2}, []int{3, 3, 3, 3})) // 12
}

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	n := len(energyDrinkA)
	dpA, dpB := int64(0), int64(0)

	for i := 0; i < n; i++ {
		newA := max(dpA+int64(energyDrinkA[i]), dpB)
		newB := max(dpB+int64(energyDrinkB[i]), dpA)
		dpA, dpB = newA, newB
	}

	if dpA > dpB {
		return dpA
	}
	return dpB
}
