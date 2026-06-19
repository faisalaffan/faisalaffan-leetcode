package main

// LeetCode #3633: Earliest Finish Time for Land and Water Rides I
// https://leetcode.com/problems/earliest-finish-time-for-land-and-water-rides-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(EarliestFinishTimeForLandAndWaterRidesI([]int{2, 8}, []int{4, 1}, []int{6}, []int{3}))
	fmt.Println(EarliestFinishTimeForLandAndWaterRidesI([]int{5}, []int{3}, []int{1}, []int{10}))
}

// Time: O(n*m) where n = len(landStartTime), m = len(waterStartTime)
// Space: O(1)
func EarliestFinishTimeForLandAndWaterRidesI(landStartTime, landDuration, waterStartTime, waterDuration []int) int {
	ans := int(^uint(0) >> 1) // max int

	// Land -> Water
	for i := range landStartTime {
		finishLand := landStartTime[i] + landDuration[i]
		for j := range waterStartTime {
			startWater := max(finishLand, waterStartTime[j])
			finish := startWater + waterDuration[j]
			if finish < ans {
				ans = finish
			}
		}
	}

	// Water -> Land
	for i := range waterStartTime {
		finishWater := waterStartTime[i] + waterDuration[i]
		for j := range landStartTime {
			startLand := max(finishWater, landStartTime[j])
			finish := startLand + landDuration[j]
			if finish < ans {
				ans = finish
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
