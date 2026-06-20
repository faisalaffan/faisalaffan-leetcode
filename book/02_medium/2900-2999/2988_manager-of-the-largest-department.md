# 2988 — Manager Of The Largest Department

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func findManagersOfLargestDepartment(employees []Employee) []ManagerResult`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2988: Manager of the Largest Department
// https://leetcode.com/problems/manager-of-the-largest-department/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: Find managers (position = 'Manager') of the department(s)
// with the most employees. If multiple departments tie, include all
// managers from those departments.

import (
	"fmt"
	"sort"
)

// Employee represents the Employees database table.
type Employee struct {
	EmpID    int
	EmpName  string
	DepID    int
	Position string
}

// ManagerResult holds the output.
type ManagerResult struct {
	EmpName string
	DepID   int
}

// findManagersOfLargestDepartment simulates the SQL query.
// Time: O(n) | Space: O(n)
// n = number of employees.
func findManagersOfLargestDepartment(employees []Employee) []ManagerResult {
	// Count employees per department.
  // HashMap: O(1) lookup
	depCount := make(map[int]int)
	for _, e := range employees {
		depCount[e.DepID]++
	}

	// Find the maximum employee count.
	maxCount := 0
	for _, count := range depCount {
		if count > maxCount {
			maxCount = count
		}
	}

	// Collect managers from departments with maxCount employees.
	var results []ManagerResult
	for _, e := range employees {
		if e.Position == "Manager" && depCount[e.DepID] == maxCount {
			results = append(results, ManagerResult{EmpName: e.EmpName, DepID: e.DepID})
		}
	}

	// Order by dep_id ASC.
  // Custom sort
	sort.Slice(results, func(i, j int) bool {
		return results[i].DepID < results[j].DepID
	})

	return results
}

func main() {
	// Test data.
	employees := []Employee{
		// Department 1 has 3 employees, managers: Alice.
		{EmpID: 1, EmpName: "Alice", DepID: 1, Position: "Manager"},
		{EmpID: 2, EmpName: "Bob", DepID: 1, Position: "Employee"},
		{EmpID: 3, EmpName: "Charlie", DepID: 1, Position: "Employee"},
		// Department 2 has 2 employees, manager: Diana.
		{EmpID: 4, EmpName: "Diana", DepID: 2, Position: "Manager"},
		{EmpID: 5, EmpName: "Eve", DepID: 2, Position: "Employee"},
		// Department 3 has 3 employees (tie with dep 1), managers: Frank.
		{EmpID: 6, EmpName: "Frank", DepID: 3, Position: "Manager"},
		{EmpID: 7, EmpName: "Grace", DepID: 3, Position: "Employee"},
		{EmpID: 8, EmpName: "Henry", DepID: 3, Position: "Employee"},
	}

	results := findManagersOfLargestDepartment(employees)

	fmt.Println("Manager(s) of Largest Department(s) (emp_name | dep_id):")
	for _, r := range results {
		fmt.Printf("%s | %d\n", r.EmpName, r.DepID)
	}
	// Expected output (departments 1 and 3 tie with 3 employees each):
	// Alice | 1
	// Frank | 3
}
```
