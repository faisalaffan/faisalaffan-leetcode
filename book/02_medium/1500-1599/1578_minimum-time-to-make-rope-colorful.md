# 1578 — Minimum Time To Make Rope Colorful

## Deskripsi

**Soal:** [1578. Minimum Time To Make Rope Colorful](https://leetcode.com/problems/minimum-time-to-make-rope-colorful/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1578: Minimum Time to Make Rope Colorful
// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinCost("abaac", []int{1, 2, 3, 4, 5}))
	fmt.Println(MinCost("abc", []int{1, 2, 3}))
	fmt.Println(MinCost("aabaa", []int{1, 2, 3, 4, 1}))
}

func MinCost(colors string, neededTime []int) int {
	// Time: O(N), Space: O(1)
	n := len(colors)
	totalTime := 0

	i := 0
	for i < n {
		j := i
		maxTime := 0
		sum := 0
		for j < n && colors[j] == colors[i] {
			maxTime = maxInt(maxTime, neededTime[j])
			sum += neededTime[j]
			j++
		}
		// Keep the max time balloon, remove the rest
		totalTime += sum - maxTime
		i = j
	}

	return totalTime
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
