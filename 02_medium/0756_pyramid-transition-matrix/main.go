package main

// LeetCode #756: Pyramid Transition Matrix
// https://leetcode.com/problems/pyramid-transition-matrix/
// Difficulty: Medium
// Time: O(7^b) worst case where b is number of blocks
// Space: O(7^b)

import "fmt"

func main() {
	fmt.Println(pyramidTransition("BCD", []string{"BCG", "CDE", "GEA", "FFF"}))
	fmt.Println(pyramidTransition("AAAA", []string{"AAB", "AAC", "BCD", "BBE", "DEF"}))
}

func pyramidTransition(bottom string, allowed []string) bool {
	memo := make(map[string]bool)
	patterns := make(map[string][]byte)

	for _, a := range allowed {
		key := a[:2]
		patterns[key] = append(patterns[key], a[2])
	}

	var dfs func(row string, next string, idx int) bool
	dfs = func(row string, next string, idx int) bool {
		if len(row) == 1 {
			return true
		}

		key := row + "#" + next
		if val, ok := memo[key]; ok {
			return val
		}

		if idx == len(row)-1 {
			if dfs(next, "", 0) {
				memo[key] = true
				return true
			}
			memo[key] = false
			return false
		}

		chars := patterns[row[idx:idx+2]]
		for _, c := range chars {
			if dfs(row, next+string(c), idx+1) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false
		return false
	}

	return dfs(bottom, "", 0)
}
