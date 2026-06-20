# 1118 — Number Of Days In A Month

## Deskripsi

**Soal:** [1118. Number Of Days In A Month](https://leetcode.com/problems/number-of-days-in-a-month/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1118: Number of Days in a Month
// https://leetcode.com/problems/number-of-days-in-a-month/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(numberOfDays(1992, 7))  // 31
	fmt.Println(numberOfDays(2000, 2))  // 29
	fmt.Println(numberOfDays(1900, 2))  // 28
}

// LeetCode submission: numberOfDays
func numberOfDays(year, month int) int {
	leap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if leap {
		days[2] = 29
	}
	return days[month]
}
```
