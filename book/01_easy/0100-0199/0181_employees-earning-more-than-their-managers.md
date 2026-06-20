# 0181 — Employees Earning More Than Their Managers

## Deskripsi

**Soal:** [0181. Employees Earning More Than Their Managers](https://leetcode.com/problems/employees-earning-more-than-their-managers/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func EmployeesEarningMoreThanTheirManagers() string`

## Solusi Go

```go
package main

// LeetCode #181: Employees Earning More Than Their Managers
// https://leetcode.com/problems/employees-earning-more-than-their-managers/
// Difficulty: Easy

import "fmt"

func EmployeesEarningMoreThanTheirManagers() string {
	return "SELECT e.name AS Employee FROM Employee e JOIN Employee m ON e.managerId = m.id WHERE e.salary > m.salary"
}

func main() {
	fmt.Println(EmployeesEarningMoreThanTheirManagers())
}
```
