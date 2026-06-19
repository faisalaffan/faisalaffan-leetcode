package main

// LeetCode #1854: Maximum Population Year
// https://leetcode.com/problems/maximum-population-year/
// Difficulty: Easy

import "fmt"

// Time: O(n + range), Space: O(range)
func MaximumPopulation(logs [][]int) int {
	delta := make([]int, 101) // 1950 to 2050
	for _, log := range logs {
		delta[log[0]-1950]++
		delta[log[1]-1950]--
	}
	maxPop := 0
	currentPop := 0
	bestYear := 1950
	for i := 0; i < 101; i++ {
		currentPop += delta[i]
		if currentPop > maxPop {
			maxPop = currentPop
			bestYear = 1950 + i
		}
	}
	return bestYear
}

func main() {
	fmt.Println(MaximumPopulation([][]int{{1993, 1999}, {2000, 2010}}))
	fmt.Println(MaximumPopulation([][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}))
}
