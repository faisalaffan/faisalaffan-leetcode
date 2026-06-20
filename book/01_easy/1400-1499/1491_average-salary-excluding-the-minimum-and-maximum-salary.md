# 1491 — Average Salary Excluding The Minimum And Maximum Salary

## Deskripsi

**Soal:** [1491. Average Salary Excluding The Minimum And Maximum Salary](https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func average(salary []int) float64`

## Solusi Go

```go
package main

// LeetCode #1491: Average Salary Excluding the Minimum and Maximum Salary
// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
// Difficulty: Easy
//
// LeetCode submission: func average(salary []int) float64

import "fmt"

func main() {
	fmt.Println(AverageSalaryExcludingTheMinimumAndMaximumSalary([]int{4000, 3000, 1000, 2000})) // 2500
	fmt.Println(AverageSalaryExcludingTheMinimumAndMaximumSalary([]int{1000, 2000, 3000}))       // 2000
}

// Time: O(n), Space: O(1)
func AverageSalaryExcludingTheMinimumAndMaximumSalary(salary []int) float64 {
	min, max := salary[0], salary[0]
	sum := 0
	for _, v := range salary {
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return float64(sum-min-max) / float64(len(salary)-2)
}
```
