# 0868 — Binary Gap

## Deskripsi

**Soal:** [0868. Binary Gap](https://leetcode.com/problems/binary-gap/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #868: Binary Gap
// https://leetcode.com/problems/binary-gap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(binaryGap(22))  // 2 (10110)
	fmt.Println(binaryGap(8))   // 0 (1000)
	fmt.Println(binaryGap(5))   // 2 (101)
	fmt.Println(binaryGap(6))   // 1 (110)
}

// binaryGap finds the longest distance between two consecutive 1s in binary representation.
// Time: O(log n). Space: O(1).
func binaryGap(n int) int {
	last := -1
	maxDist := 0
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				if i-last > maxDist {
					maxDist = i - last
				}
			}
			last = i
		}
		n >>= 1
	}
	return maxDist
}
```
