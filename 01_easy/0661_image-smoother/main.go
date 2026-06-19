package main

// LeetCode #661: Image Smoother
// https://leetcode.com/problems/image-smoother/
// Difficulty: Easy

import "fmt"

func main() {
	img := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println(imageSmoother(img))
	// [[0,0,0],[0,0,0],[0,0,0]]

	img2 := [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}
	fmt.Println(imageSmoother(img2))
}

// imageSmoother applies a 3x3 smoother to each cell of the image.
// Time: O(m*n). Space: O(m*n).
func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 0}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum, count := 0, 0
			for _, d := range directions {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					sum += img[ni][nj]
					count++
				}
			}
			result[i][j] = sum / count
		}
	}
	return result
}
