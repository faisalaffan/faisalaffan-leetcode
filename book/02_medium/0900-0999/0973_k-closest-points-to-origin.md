# 0973 — K Closest Points To Origin

## Deskripsi

**Soal:** [0973. K Closest Points To Origin](https://leetcode.com/problems/k-closest-points-to-origin/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) average, O(n^2) worst-case  
**Kompleksitas Ruang:** O(1) excluding output

**Algoritma:** —

> **Ide Kunci:** QuickSelect (partition-based selection)

## Solusi Go

```go
package main

// LeetCode #973: K Closest Points to Origin
// https://leetcode.com/problems/k-closest-points-to-origin/
// Difficulty: Medium
//
// Approach: QuickSelect (partition-based selection)
// Time: O(n) average, O(n^2) worst-case
// Space: O(1) excluding output

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(kClosest([][]int{{1, 3}, {-2, 2}}, 1))         // [[-2,2]]
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)) // [[3,3],[-2,4]]
}

func kClosest(points [][]int, k int) [][]int {
	if len(points) <= k {
		return points
	}

	// QuickSelect
	left, right := 0, len(points)-1
  // Loop two-pointer: kiri vs kanan
	for left < right {
		pivot := partition(points, left, right)
		if pivot == k {
			break
		} else if pivot < k {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return points[:k]
}

func partition(points [][]int, left, right int) int {
	pivotIdx := left + rand.Intn(right-left+1)
	points[pivotIdx], points[right] = points[right], points[pivotIdx]
	pivotDist := dist(points[right])
	storeIdx := left

	for i := left; i < right; i++ {
		if dist(points[i]) <= pivotDist {
			points[storeIdx], points[i] = points[i], points[storeIdx]
			storeIdx++
		}
	}
	points[storeIdx], points[right] = points[right], points[storeIdx]
	return storeIdx
}

func dist(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}
```
