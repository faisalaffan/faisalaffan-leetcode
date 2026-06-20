# 0577 — Employee Bonus

## Deskripsi

**Soal:** [0577. Employee Bonus](https://leetcode.com/problems/employee-bonus/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func EmployeeBonus() string`

## Solusi Go

```go
package main

// LeetCode #577: Employee Bonus
// https://leetcode.com/problems/employee-bonus/
// Difficulty: Easy

import "fmt"

func EmployeeBonus() string {
	return "SELECT e.name, b.bonus FROM Employee e LEFT JOIN Bonus b ON e.empId = b.empId WHERE b.bonus < 1000 OR b.bonus IS NULL"
}

func main() {
	fmt.Println(EmployeeBonus())
}
```
