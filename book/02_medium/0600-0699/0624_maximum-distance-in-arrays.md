# 0624 — Maximum Distance In Arrays

## Deskripsi

**Soal:** [0624. Maximum Distance In Arrays](https://leetcode.com/problems/maximum-distance-in-arrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = number of arrays  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #624: Maximum Distance in Arrays
// https://leetcode.com/problems/maximum-distance-in-arrays/
// Difficulty: Medium
// Time: O(n) where n = number of arrays
// Space: O(1)

import (
	"fmt"
)

func main() {
	arrays := [][]int{
		{1, 2, 3},
		{4, 5},
		{1, 2, 3},
	}
	fmt.Println(MaxDistance(arrays))
}

func MaxDistance(arrays [][]int) int {
	minVal := arrays[0][0]
	maxVal := arrays[0][len(arrays[0])-1]
	maxDist := 0

	for i := 1; i < len(arrays); i++ {
		arr := arrays[i]
		dist1 := abs(arr[len(arr)-1] - minVal)
		dist2 := abs(maxVal - arr[0])
		if dist1 > maxDist {
			maxDist = dist1
		}
		if dist2 > maxDist {
			maxDist = dist2
		}
		if arr[0] < minVal {
			minVal = arr[0]
		}
		if arr[len(arr)-1] > maxVal {
			maxVal = arr[len(arr)-1]
		}
	}

	return maxDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
