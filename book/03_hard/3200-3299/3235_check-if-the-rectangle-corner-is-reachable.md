# 3235 — Check If The Rectangle Corner Is Reachable

## Deskripsi

**Soal:** [3235. Check If The Rectangle Corner Is Reachable](https://leetcode.com/problems/check-if-the-rectangle-corner-is-reachable/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), DFS (Depth-First Search / pencarian kedalaman)

## Solusi Go

```go
package main

// LeetCode #3235: Check if the Rectangle Corner Is Reachable
// https://leetcode.com/problems/check-if-the-rectangle-corner-is-reachable/
// Difficulty: Hard
//
// DFS on circles. A path from (0,0) to (X,Y) is blocked iff circles form a
// connected barrier from the {left, top} edges to the {right, bottom} edges,
// or either corner lies inside a circle.
//
// Two circles intersect if distance between centers ≤ sum of radii. Their
// intersection point (weighted midpoint) must lie within the rectangle to
// count as a blocking connection.

import "fmt"

func main() {
	// Example 1: X=3,Y=4,circles=[[2,1,1]] => true
	fmt.Println(canReachCorner(3, 4, [][]int{{2, 1, 1}}))
	// Example 2: start inside circle
	fmt.Println(canReachCorner(3, 3, [][]int{{1, 1, 2}}))
	// Example 3: circle blocks top to bottom
	fmt.Println(canReachCorner(5, 5, [][]int{{3, 3, 3}}))
	// Example 4: two circles forming a barrier
	fmt.Println(canReachCorner(10, 10, [][]int{{2, 5, 3}, {8, 5, 3}}))
	// Example 5: free path
	fmt.Println(canReachCorner(10, 10, [][]int{{1, 1, 1}, {9, 9, 1}}))
}

func canReachCorner(X int, Y int, circles [][]int) bool {
	n := len(circles)

	// Check if corners are inside any circle
	for _, c := range circles {
		x, y, r := c[0], c[1], c[2]
		if pointInCircle(0, 0, x, y, r) || pointInCircle(X, Y, x, y, r) {
			return false
		}
	}

	// Circle touches left or top edge
	touchesLeftTop := func(cx, cy, r int) bool {
		// Circle extends past left wall AND center is within rectangle vertically
		if abs(cx) <= r && 0 <= cy && cy <= Y {
			return true
		}
		// Circle extends past top wall AND center is within rectangle horizontally
		if abs(cy-Y) <= r && 0 <= cx && cx <= X {
			return true
		}
		return false
	}

	// Circle touches right or bottom edge
	touchesRightBottom := func(cx, cy, r int) bool {
		// Circle extends past right wall AND center is within rectangle vertically
		if abs(cx-X) <= r && 0 <= cy && cy <= Y {
			return true
		}
		// Circle extends past bottom wall AND center is within rectangle horizontally
		if abs(cy) <= r && 0 <= cx && cx <= X {
			return true
		}
		return false
	}

  // Membuat slice untuk menyimpan hasil
	visited := make([]bool, n)

	var dfs func(int) bool
	dfs = func(i int) bool {
		x1, y1, r1 := circles[i][0], circles[i][1], circles[i][2]
		if touchesRightBottom(x1, y1, r1) {
			return true
		}
		visited[i] = true

		for j := 0; j < n; j++ {
			if visited[j] {
				continue
			}
			x2, y2, r2 := circles[j][0], circles[j][1], circles[j][2]
			dx, dy := x1-x2, y1-y2
			distSq := dx*dx + dy*dy
			sumR := r1 + r2
			if distSq > sumR*sumR {
				continue
			}
			// Weighted midpoint must lie inside the rectangle
			// (x1*r2 + x2*r1) / (r1+r2) < X AND (y1*r2 + y2*r1) / (r1+r2) < Y
			if x1*r2+x2*r1 < sumR*X && y1*r2+y2*r1 < sumR*Y {
				if dfs(j) {
					return true
				}
			}
		}
		return false
	}

	for i, c := range circles {
		x, y, r := c[0], c[1], c[2]
		if !visited[i] && touchesLeftTop(x, y, r) {
			if dfs(i) {
				return false
			}
		}
	}

	return true
}

func pointInCircle(px, py, cx, cy, r int) bool {
	dx, dy := px-cx, py-cy
	return dx*dx+dy*dy <= r*r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
