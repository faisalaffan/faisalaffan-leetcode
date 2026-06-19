package main

// LeetCode #2513: Minimize the Maximum of Two Arrays
// https://leetcode.com/problems/minimize-the-maximum-of-two-arrays/
// Difficulty: Medium
// Time: O(log(max)) | Space: O(1)
// Binary search on answer. For value X:
//   count1 = X - X/divisor1 (numbers not divisible by divisor1)
//   count2 = X - X/divisor2 (numbers not divisible by divisor2)
//   common = X - X/lcm (numbers not divisible by either)
//   We need count1 >= uniqueCnt1, count2 >= uniqueCnt2, common >= uniqueCnt1 + uniqueCnt2

import "fmt"

func main() {
	fmt.Println(minimizeSet(2, 7, 1, 3)) // 4
	fmt.Println(minimizeSet(3, 5, 2, 1)) // 3
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	lo, hi := 1, 1<<31-1
	l := lcm(divisor1, divisor2)

	for lo < hi {
		mid := (lo + hi) / 2
		cnt1 := mid - mid/divisor1
		cnt2 := mid - mid/divisor2
		common := mid - mid/l

		if cnt1 >= uniqueCnt1 && cnt2 >= uniqueCnt2 && common >= uniqueCnt1+uniqueCnt2 {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
