package main

// LeetCode #3733: Minimum Time to Complete All Deliveries
// https://leetcode.com/problems/minimum-time-to-complete-all-deliveries/
// Difficulty: Medium
// Time: O(log(maxTime)) | Space: O(1)

import "fmt"

func minimumTimeToCompleteAllDeliveries(d []int, r []int) int64 {
	a, b := int64(d[0]), int64(d[1])
	x, y := int64(r[0]), int64(r[1])

	// Check if t hours is enough
	check := func(t int64) bool {
		// Hours drone 2 cannot work (drone 1 only)
		only1 := t / y - t/lcm(x, y)
		// Hours drone 1 cannot work (drone 2 only)
		only2 := t / x - t/lcm(x, y)
		// Hours both can work
		both := t - t/x - t/y + t/lcm(x, y)

		needA := a - only1
		if needA < 0 {
			needA = 0
		}
		needB := b - only2
		if needB < 0 {
			needB = 0
		}
		return needA+needB <= both
	}

	lo, hi := a+b, (a+b)*max(x, y)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumTimeToCompleteAllDeliveries([]int{3, 1}, []int{2, 3}))
	fmt.Println(minimumTimeToCompleteAllDeliveries([]int{1, 3}, []int{2, 2}))
	fmt.Println(minimumTimeToCompleteAllDeliveries([]int{2, 1}, []int{3, 4}))
}
