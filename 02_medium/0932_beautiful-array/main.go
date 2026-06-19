package main

// LeetCode #932: Beautiful Array
// https://leetcode.com/problems/beautiful-array/
// Difficulty: Medium

import "fmt"

// Time: O(n log n) | Space: O(n)
func beautifulArray(n int) []int {
	memo := make(map[int][]int)
	var dfs func(int) []int
	dfs = func(n int) []int {
		if v, ok := memo[n]; ok {
			return v
		}
		res := make([]int, n)
		if n == 1 {
			res[0] = 1
		} else {
			left := dfs((n + 1) / 2)
			right := dfs(n / 2)
			for i, v := range left {
				res[i] = 2*v - 1
			}
			for i, v := range right {
				res[(n+1)/2+i] = 2 * v
			}
		}
		memo[n] = res
		return res
	}
	return dfs(n)
}

func main() {
	fmt.Println(beautifulArray(4))
	fmt.Println(beautifulArray(5))
	fmt.Println(beautifulArray(1))
}
