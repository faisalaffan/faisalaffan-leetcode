package main

// LeetCode #3025: Find the Number of Ways to Place People I
// https://leetcode.com/problems/find-the-number-of-ways-to-place-people-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(numberOfPairs([][]int{{1, 1}, {2, 2}, {3, 3}}))
	fmt.Println(numberOfPairs([][]int{{6, 2}, {4, 4}, {2, 6}}))
	fmt.Println(numberOfPairs([][]int{{3, 1}, {1, 3}, {1, 1}}))
}

func numberOfPairs(points [][]int) (ans int) {
	n := len(points)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]
			if x1 > x2 || y1 < y2 {
				continue
			}
			ok := true
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}
				x3, y3 := points[k][0], points[k][1]
				if x3 >= x2 && x3 <= x1 && y3 >= y2 && y3 <= y1 {
					if x3 == x2 && y3 == y2 || x3 == x1 && y3 == y1 {
						continue
					}
					if x3 >= x2 && x3 <= x1 && y3 >= y2 && y3 <= y1 {
						ok = false
						break
					}
				}
			}
			if ok {
				ans++
			}
		}
	}
	return
}
