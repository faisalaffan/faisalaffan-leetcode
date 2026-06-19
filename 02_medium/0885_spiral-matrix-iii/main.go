package main

// LeetCode #885: Spiral Matrix III
// https://leetcode.com/problems/spiral-matrix-iii/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SpiralMatrixIii(1, 4, 0, 0))
	fmt.Println(SpiralMatrixIii(5, 6, 1, 4))
}

// Time: O(rows * cols) | Space: O(rows * cols)
func SpiralMatrixIii(rows int, cols int, rStart int, cStart int) [][]int {
	total := rows * cols
	ans := make([][]int, 0, total)
	ans = append(ans, []int{rStart, cStart})
	if total == 1 {
		return ans
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	r, c := rStart, cStart
	step := 1
	dir := 0

	for len(ans) < total {
		for i := 0; i < 2; i++ {
			for j := 0; j < step; j++ {
				r += dirs[dir][0]
				c += dirs[dir][1]
				if r >= 0 && r < rows && c >= 0 && c < cols {
					ans = append(ans, []int{r, c})
				}
			}
			dir = (dir + 1) % 4
		}
		step++
	}

	return ans
}
