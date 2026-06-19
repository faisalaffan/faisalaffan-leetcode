package main

// LeetCode #3635: Earliest Finish Time for Land and Water Rides II
// https://leetcode.com/problems/earliest-finish-time-for-land-and-water-rides-ii/
// Difficulty: Medium
// Time: O(n + m) | Space: O(1)

import "fmt"

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	// Try land -> water
	minLandEnd := int(^uint(0) >> 1)
	for i := range landStartTime {
		end := landStartTime[i] + landDuration[i]
		if end < minLandEnd {
			minLandEnd = end
		}
	}
	ans := int(^uint(0) >> 1)
	for i := range waterStartTime {
		start := waterStartTime[i]
		if minLandEnd > start {
			start = minLandEnd
		}
		end := start + waterDuration[i]
		if end < ans {
			ans = end
		}
	}

	// Try water -> land
	minWaterEnd := int(^uint(0) >> 1)
	for i := range waterStartTime {
		end := waterStartTime[i] + waterDuration[i]
		if end < minWaterEnd {
			minWaterEnd = end
		}
	}
	for i := range landStartTime {
		start := landStartTime[i]
		if minWaterEnd > start {
			start = minWaterEnd
		}
		end := start + landDuration[i]
		if end < ans {
			ans = end
		}
	}

	return ans
}

func main() {
	fmt.Println(earliestFinishTime(
		[]int{1, 2, 3}, []int{3, 2, 1},
		[]int{2, 3}, []int{4, 1},
	))
	fmt.Println(earliestFinishTime(
		[]int{5}, []int{10},
		[]int{2, 8}, []int{3, 2},
	))
	fmt.Println(earliestFinishTime(
		[]int{0, 10}, []int{5, 2},
		[]int{3, 7}, []int{4, 2},
	))
}
