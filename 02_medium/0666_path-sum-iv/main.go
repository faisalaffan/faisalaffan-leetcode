package main

// LeetCode #666: Path Sum IV
// https://leetcode.com/problems/path-sum-iv/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(pathSumIV([]int{113, 215, 221}))
	fmt.Println(pathSumIV([]int{113, 221}))
}

func pathSumIV(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tree := make(map[int]int)
	for _, num := range nums {
		tree[num/10] = num % 10
	}

	total := 0
	var dfs func(key int, sum int)
	dfs = func(key int, sum int) {
		depth := key / 10
		pos := key % 10
		leftKey := (depth+1)*10 + pos*2 - 1
		rightKey := (depth+1)*10 + pos*2

		curSum := sum + tree[key]

		_, hasLeft := tree[leftKey]
		_, hasRight := tree[rightKey]

		if !hasLeft && !hasRight {
			total += curSum
			return
		}

		if hasLeft {
			dfs(leftKey, curSum)
		}
		if hasRight {
			dfs(rightKey, curSum)
		}
	}

	dfs(nums[0]/10, 0)
	return total
}
