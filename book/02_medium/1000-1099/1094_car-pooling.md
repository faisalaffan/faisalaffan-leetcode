# 1094 — Car Pooling

## Deskripsi

**Soal:** [1094. Car Pooling](https://leetcode.com/problems/car-pooling/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + maxLocation)  
**Kompleksitas Ruang:** O(maxLocation)

**Algoritma:** Prefix Sum (jumlah kumulatif)

> **Ide Kunci:** Difference array (prefix sum)

## Solusi Go

```go
package main

// LeetCode #1094: Car Pooling
// https://leetcode.com/problems/car-pooling/
// Difficulty: Medium
//
// Approach: Difference array (prefix sum)
// Time: O(n + maxLocation)
// Space: O(maxLocation)

import "fmt"

func main() {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4)) // false
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5)) // true
}

func carPooling(trips [][]int, capacity int) bool {
	maxLoc := 0
	for _, t := range trips {
		if t[2] > maxLoc {
			maxLoc = t[2]
		}
	}

  // Membuat slice untuk menyimpan hasil
	diff := make([]int, maxLoc+2)
	for _, t := range trips {
		diff[t[1]] += t[0]
		diff[t[2]] -= t[0]
	}

	current := 0
	for _, d := range diff {
		current += d
		if current > capacity {
			return false
		}
	}

	return true
}
```
