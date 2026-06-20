# 2769 — Find The Maximum Achievable Number

## Deskripsi

**Soal:** [2769. Find The Maximum Achievable Number](https://leetcode.com/problems/find-the-maximum-achievable-number/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2769: Find the Maximum Achievable Number
// https://leetcode.com/problems/find-the-maximum-achievable-number/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindTheMaximumAchievableNumber(4, 1))
	fmt.Println(FindTheMaximumAchievableNumber(3, 2))
}

func FindTheMaximumAchievableNumber(num int, t int) int {
	return num + 2*t
}
```
