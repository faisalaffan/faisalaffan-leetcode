# 0739 — Daily Temperatures

## Deskripsi

**Soal:** [0739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

## Solusi Go

```go
package main

// LeetCode #739: Daily Temperatures
// https://leetcode.com/problems/daily-temperatures/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = i - idx
		}
		stack = append(stack, i)
	}

	return result
}
```
