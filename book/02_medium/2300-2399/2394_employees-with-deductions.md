# 2394 — Employees With Deductions

## Deskripsi

**Soal:** [2394. Employees With Deductions](https://leetcode.com/problems/employees-with-deductions/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2394: Employees With Deductions
// https://leetcode.com/problems/employees-with-deductions/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// Calculate actual hours worked per employee and compare with required hours.

import "fmt"

type Log struct {
	EmpID        int
	Timestamp    int
	IsLogin      bool
}

func main() {
	logs := []Log{
		{1, 100, true},
		{1, 200, false},
		{2, 50, true},
		{2, 150, false},
		{1, 300, true},
		{1, 400, false},
	}
	// Employee 1: total 200 min = 3.33 hrs, needs 4 hrs
	// Employee 2: total 100 min = 1.67 hrs, needs 4 hrs
	fmt.Println(calculateDeductions(logs, []int{4, 4})) // [1, 2] (both under)

	logs2 := []Log{
		{1, 0, true},
		{1, 240, false},
		{2, 0, true},
		{2, 480, false},
	}
	fmt.Println(calculateDeductions(logs2, []int{4, 8})) // [2]
}

func calculateDeductions(logs []Log, requiredHours []int) []int {
  // Membuat map untuk pencarian O(1): key → value
	hours := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	loginTime := make(map[int]int)

	for _, l := range logs {
		if l.IsLogin {
			loginTime[l.EmpID] = l.Timestamp
		} else {
			if start, ok := loginTime[l.EmpID]; ok {
				hours[l.EmpID] += l.Timestamp - start
				delete(loginTime, l.EmpID)
			}
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0)
	for empID := 1; empID <= len(requiredHours); empID++ {
		workedMin := hours[empID]
		neededMin := requiredHours[empID-1] * 60
		if workedMin < neededMin {
			result = append(result, empID)
		}
	}
	return result
}
```
