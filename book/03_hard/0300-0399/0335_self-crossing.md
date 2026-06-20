# 0335 — Self Crossing

## Deskripsi

**Soal:** [0335. Self Crossing](https://leetcode.com/problems/self-crossing/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func isSelfCrossing(distance []int) bool`

## Solusi Go

```go
package main

// LeetCode #335: Self Crossing
// https://leetcode.com/problems/self-crossing/
// Difficulty: Hard

import "fmt"

func isSelfCrossing(distance []int) bool {
	n := len(distance)
	if n < 4 {
		return false
	}

	for i := 3; i < n; i++ {
		// Case 1: i-th line crosses (i-3)-th line
		// x[i-3] >= x[i-1] and x[i] >= x[i-2]
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}

		// Case 2: i-th line crosses (i-4)-th line (overlap of 5 edges)
		if i >= 4 {
			if distance[i-1] == distance[i-3] &&
				distance[i] >= distance[i-2]-distance[i-4] {
				return true
			}
		}

		// Case 3: i-th line crosses (i-5)-th line (overlap of 6 edges)
		if i >= 5 {
			if distance[i-2] >= distance[i-4] &&
				distance[i-3]-distance[i-5] >= 0 &&
				distance[i-1] >= distance[i-3]-distance[i-5] &&
				distance[i-1] <= distance[i-3] &&
				distance[i] >= distance[i-2]-distance[i-4] {
				return true
			}
		}
	}
	return false
}

func main() {
	// Example 1: crossing
	fmt.Println(isSelfCrossing([]int{2, 1, 1, 2}))
	// true

	// Example 2: no crossing
	fmt.Println(isSelfCrossing([]int{1, 2, 3, 4}))
	// false

	// Example 3: crossing
	fmt.Println(isSelfCrossing([]int{1, 1, 1, 2, 1}))
	// true

	// Example 4: no crossing
	fmt.Println(isSelfCrossing([]int{1, 1, 2, 2, 3, 3, 4, 4, 10, 4, 4, 3, 3, 2, 2, 1, 1}))
	// false
}
```
