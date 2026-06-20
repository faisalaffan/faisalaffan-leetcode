# 1266 — Minimum Time Visiting All Points

## Deskripsi

**Soal:** [1266. Minimum Time Visiting All Points](https://leetcode.com/problems/minimum-time-visiting-all-points/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1266: Minimum Time Visiting All Points
// https://leetcode.com/problems/minimum-time-visiting-all-points/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}})) // 7
	fmt.Println(minTimeToVisitAllPoints([][]int{{3, 2}, {-2, 2}}))         // 5
}

// LeetCode submission: minTimeToVisitAllPoints
func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	for i := 1; i < len(points); i++ {
		dx := points[i][0] - points[i-1][0]
		dy := points[i][1] - points[i-1][1]
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		if dx > dy {
			ans += dx
		} else {
			ans += dy
		}
	}
	return ans
}
```
