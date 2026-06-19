package main

// LeetCode #786: K-th Smallest Prime Fraction
// https://leetcode.com/problems/k-th-smallest-prime-fraction/
// Difficulty: Medium
// Time: O(n log max)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
	fmt.Println(kthSmallestPrimeFraction([]int{1, 7}, 1))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	left, right := 0.0, 1.0

	for left < right {
		mid := (left + right) / 2.0
		count := 0
		maxFraction := 0.0
		p, q := 0, 1

		j := 1
		for i := 0; i < n; i++ {
			for j < n && float64(arr[i])/float64(arr[j]) > mid {
				j++
			}
			if j == n {
				break
			}
			count += n - j

			fraction := float64(arr[i]) / float64(arr[j])
			if fraction > maxFraction {
				maxFraction = fraction
				p, q = arr[i], arr[j]
			}
		}

		if count == k {
			return []int{p, q}
		} else if count < k {
			left = mid
		} else {
			right = mid
		}
	}

	return nil
}
