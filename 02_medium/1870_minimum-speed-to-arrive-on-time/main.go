package main

// LeetCode #1870: Minimum Speed to Arrive on Time
// https://leetcode.com/problems/minimum-speed-to-arrive-on-time/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSpeedOnTime([]int{1, 3, 2}, 6.0))
	fmt.Println(MinSpeedOnTime([]int{1, 3, 2}, 2.7))
	fmt.Println(MinSpeedOnTime([]int{1, 3, 2}, 1.9))
}

// Time: O(n log maxDist), Space: O(1)
func MinSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	if hour <= float64(n-1) {
		return -1
	}

	lo, hi := 1, 10000000
	ans := -1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if canReach(dist, mid, hour) {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func canReach(dist []int, speed int, hour float64) bool {
	time := 0.0
	for i := 0; i < len(dist)-1; i++ {
		time += float64((dist[i] + speed - 1) / speed) // ceil division
	}
	time += float64(dist[len(dist)-1]) / float64(speed)
	return time <= hour
}
