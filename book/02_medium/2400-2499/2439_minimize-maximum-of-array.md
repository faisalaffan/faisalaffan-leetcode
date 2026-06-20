# 2439 — Minimize Maximum Of Array

## Deskripsi

**Soal:** [2439. Minimize Maximum Of Array](https://leetcode.com/problems/minimize-maximum-of-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2439: Minimize Maximum of Array
// https://leetcode.com/problems/minimize-maximum-of-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Prefix average approach: we can distribute value to the left.
// The min possible max is the max prefix average (ceil).

import "fmt"

func main() {
	fmt.Println(minimizeArrayValue([]int{3, 7, 1, 6})) // 5
	fmt.Println(minimizeArrayValue([]int{10, 1}))      // 10
}

func minimizeArrayValue(nums []int) int {
	var sum int64
	ans := 0
	for i, v := range nums {
		sum += int64(v)
		avg := int((sum + int64(i)) / int64(i+1)) // ceil division
		if avg > ans {
			ans = avg
		}
	}
	return ans
}
```
