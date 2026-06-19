package main

// LeetCode #3697: Compute Decimal Representation
// https://leetcode.com/problems/compute-decimal-representation/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ComputeDecimalRepresentation(537))
	fmt.Println(ComputeDecimalRepresentation(102))
	fmt.Println(ComputeDecimalRepresentation(6))
}

// Time: O(log n)
// Space: O(log n)
func ComputeDecimalRepresentation(n int) []int {
	res := make([]int, 0)
	place := 1
	for n > 0 {
		d := n % 10
		if d != 0 {
			res = append(res, d*place)
		}
		place *= 10
		n /= 10
	}

	// Reverse to descending order
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
