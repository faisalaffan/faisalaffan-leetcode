package main

// LeetCode #826: Most Profit Assigning Work
// https://leetcode.com/problems/most-profit-assigning-work/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MostProfitAssigningWork([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}))
	fmt.Println(MostProfitAssigningWork([]int{85, 47, 57}, []int{24, 66, 99}, []int{40, 25, 25}))
	fmt.Println(MostProfitAssigningWork([]int{68, 35, 52, 47, 86}, []int{67, 17, 1, 81, 3}, []int{92, 10, 85, 84, 82}))
}

// Time: O(n log n + m log m) | Space: O(n)
func MostProfitAssigningWork(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	jobs := make([][2]int, n)
	for i := range difficulty {
		jobs[i] = [2]int{difficulty[i], profit[i]}
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})
	sort.Ints(worker)

	ans := 0
	idx := 0
	maxProfit := 0

	for _, w := range worker {
		for idx < n && jobs[idx][0] <= w {
			if jobs[idx][1] > maxProfit {
				maxProfit = jobs[idx][1]
			}
			idx++
		}
		ans += maxProfit
	}

	return ans
}
