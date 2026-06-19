package main

// LeetCode #531: Lonely Pixel I
// https://leetcode.com/problems/lonely-pixel-i/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m + n)

import "fmt"

func main() {
	picture := [][]byte{
		{'W', 'W', 'B'},
		{'W', 'B', 'W'},
		{'B', 'W', 'W'},
	}
	fmt.Println(FindLonelyPixel(picture))
}

func FindLonelyPixel(picture [][]byte) int {
	m, n := len(picture), len(picture[0])
	rows := make([]int, m)
	cols := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' {
				rows[i]++
				cols[j]++
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' && rows[i] == 1 && cols[j] == 1 {
				count++
			}
		}
	}

	return count
}
