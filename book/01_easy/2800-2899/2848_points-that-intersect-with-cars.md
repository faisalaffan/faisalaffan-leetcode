# 2848 — Points That Intersect With Cars

## Deskripsi

**Soal:** [2848. Points That Intersect With Cars](https://leetcode.com/problems/points-that-intersect-with-cars/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n * range)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2848: Points That Intersect With Cars
// https://leetcode.com/problems/points-that-intersect-with-cars/
// Difficulty: Easy
// Time: O(n * range) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(PointsThatIntersectWithCars([][]int{{3, 6}, {1, 5}, {4, 7}}))
	fmt.Println(PointsThatIntersectWithCars([][]int{{1, 3}, {5, 8}}))
}

func PointsThatIntersectWithCars(nums [][]int) int {
  // Membuat slice untuk menyimpan hasil
	points := make([]bool, 101)
	for _, car := range nums {
		for p := car[0]; p <= car[1]; p++ {
			points[p] = true
		}
	}
	count := 0
	for _, v := range points {
		if v {
			count++
		}
	}
	return count
}
```
