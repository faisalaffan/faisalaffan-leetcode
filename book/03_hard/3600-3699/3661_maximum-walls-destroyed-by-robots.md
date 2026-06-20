# 3661 — Maximum Walls Destroyed By Robots

## Deskripsi

**Soal:** [3661. Maximum Walls Destroyed By Robots](https://leetcode.com/problems/maximum-walls-destroyed-by-robots/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Sort robots and walls. For each robot, count walls within

## Solusi Go

```go
package main

// LeetCode #3661: Maximum Walls Destroyed by Robots
// https://leetcode.com/problems/maximum-walls-destroyed-by-robots/
// Difficulty: Hard
//
// Given robots at positions with bullet ranges and walls at positions,
// find maximum number of unique walls a single robot can destroy.
//
// Approach: Sort robots and walls. For each robot, count walls within
// its reach (left and right). Track max.

import "fmt"
import "sort"

func main() {
	// Example 1
	fmt.Println(maxWalls([]int{1, 5, 10}, []int{2, 3, 4}, []int{2, 4, 6, 8, 11}))
	// Example 2
	fmt.Println(maxWalls([]int{3, 7}, []int{2, 2}, []int{1, 4, 6, 9}))
	// Edge: single robot
	fmt.Println(maxWalls([]int{5}, []int{3}, []int{1, 2, 6, 7, 8}))
	// Edge: no walls
	fmt.Println(maxWalls([]int{1}, []int{5}, []int{}))
}

func maxWalls(robots []int, distance []int, walls []int) int {
	sort.Ints(walls)
	result := 0

	for i, r := range robots {
		d := distance[i]
		left := r - d
		right := r + d
		// Count walls in [left, right]
		l := sort.SearchInts(walls, left)
		rr := sort.SearchInts(walls, right+1)
		cnt := rr - l
		if cnt > result {
			result = cnt
		}
	}
	return result
}
```
