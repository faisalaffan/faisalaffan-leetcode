# 1523 — Count Odd Numbers In An Interval Range

## Deskripsi

**Soal:** [1523. Count Odd Numbers In An Interval Range](https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func countOdds(low int, high int) int`

## Solusi Go

```go
package main

// LeetCode #1523: Count Odd Numbers in an Interval Range
// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/
// Difficulty: Easy
//
// LeetCode submission: func countOdds(low int, high int) int

import "fmt"

func main() {
	fmt.Println(CountOddNumbersInAnIntervalRange(3, 7)) // 3
	fmt.Println(CountOddNumbersInAnIntervalRange(8, 10)) // 1
	fmt.Println(CountOddNumbersInAnIntervalRange(0, 0)) // 0
}

// Time: O(1), Space: O(1)
func CountOddNumbersInAnIntervalRange(low int, high int) int {
	return (high + 1) / 2 - low / 2
}
```
