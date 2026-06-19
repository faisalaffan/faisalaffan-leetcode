package main

// LeetCode #733: Flood Fill
// https://leetcode.com/problems/flood-fill/
// Difficulty: Easy

import "fmt"

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(floodFill(image, 1, 1, 2))
	// [[2,2,2],[2,2,0],[2,0,1]]

	image2 := [][]int{{0, 0, 0}, {0, 0, 0}}
	fmt.Println(floodFill(image2, 0, 0, 0))
	// [[0,0,0],[0,0,0]]
}

// floodFill performs a flood fill on the image starting from (sr, sc).
// Time: O(m*n). Space: O(m*n) for recursion stack.
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}
	dfs(image, sr, sc, originalColor, color)
	return image
}

func dfs(image [][]int, r, c, originalColor, newColor int) {
	if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) {
		return
	}
	if image[r][c] != originalColor {
		return
	}
	image[r][c] = newColor
	dfs(image, r+1, c, originalColor, newColor)
	dfs(image, r-1, c, originalColor, newColor)
	dfs(image, r, c+1, originalColor, newColor)
	dfs(image, r, c-1, originalColor, newColor)
}
