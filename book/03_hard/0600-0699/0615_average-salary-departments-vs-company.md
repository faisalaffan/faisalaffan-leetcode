# 0615 — Average Salary Departments Vs Company

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func calcAvg(vals []int) float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(S + E + M*D) where S=salaries, E=employees, M=months, D=depts  |  **Ruang:** O(S + E)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #615: Average Salary: Departments VS Company
// https://leetcode.com/problems/average-salary-departments-vs-company/
// Difficulty: Hard [Paid]
//
// For each month, compare the average salary of each department to the company-wide
// average. Return 'higher', 'lower', or 'same'.

// Employee represents a row in the Employee table.
type Employee struct {
	ID           int
	DepartmentID int
}

// SalaryRecord represents a row in the Salary table.
type SalaryRecord struct {
	ID         int
	EmployeeID int
	Amount     int
	PayDate    string // "YYYY-MM"
}

// DeptComparison holds one result row.
type DeptComparison struct {
	PayMonth     string
	DepartmentID int
	Comparison   string // "higher", "lower", "same"
}

// calcAvg computes rounded average of int slice.
func calcAvg(vals []int) float64 {
	if len(vals) == 0 {
		return 0
	}
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return float64(sum) / float64(len(vals))
}

// averageSalaryDepartmentsVsCompany compares monthly dept avg vs company avg.
// Time: O(S + E + M*D) where S=salaries, E=employees, M=months, D=depts
// Space: O(S + E)
func averageSalaryDepartmentsVsCompany(employees []Employee, salaries []SalaryRecord) []DeptComparison {
	// Map employee ID to department ID.
  // HashMap: O(1) lookup
	empDept := make(map[int]int)
	for _, e := range employees {
		empDept[e.ID] = e.DepartmentID
	}

	// Group salaries by month (company-wide).
  // HashMap: O(1) lookup
	monthSalaries := make(map[string][]int)
	// Group salaries by month then department.
  // HashMap: O(1) lookup
	deptMonthSalaries := make(map[string]map[int][]int) // month -> deptID -> []amount

	for _, s := range salaries {
		deptID, ok := empDept[s.EmployeeID]
		if !ok {
			continue // employee not found
		}
		monthSalaries[s.PayDate] = append(monthSalaries[s.PayDate], s.Amount)

		if deptMonthSalaries[s.PayDate] == nil {
			deptMonthSalaries[s.PayDate] = make(map[int][]int)
		}
		deptMonthSalaries[s.PayDate][deptID] = append(deptMonthSalaries[s.PayDate][deptID], s.Amount)
	}

	// Sort months.
	var months []string
	for m := range monthSalaries {
		months = append(months, m)
	}
	sort.Strings(months)

	var result []DeptComparison

	for _, month := range months {
		companyAvg := calcAvg(monthSalaries[month])

		for deptID, deptAmounts := range deptMonthSalaries[month] {
			deptAvg := calcAvg(deptAmounts)
			comparison := "same"
			if deptAvg > companyAvg {
				comparison = "higher"
			} else if deptAvg < companyAvg {
				comparison = "lower"
			}
			result = append(result, DeptComparison{
				PayMonth:     month,
				DepartmentID: deptID,
				Comparison:   comparison,
			})
		}
	}

	// Sort for deterministic output.
  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		if result[i].PayMonth != result[j].PayMonth {
			return result[i].PayMonth < result[j].PayMonth
		}
		return result[i].DepartmentID < result[j].DepartmentID
	})

	return result
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0615 Average Salary: Departments VS Company ===")

	employees := []Employee{
		{ID: 1, DepartmentID: 1},
		{ID: 2, DepartmentID: 2},
		{ID: 3, DepartmentID: 2},
	}

	salaries := []SalaryRecord{
		{ID: 1, EmployeeID: 1, Amount: 8000, PayDate: "2017-03"},
		{ID: 2, EmployeeID: 2, Amount: 4000, PayDate: "2017-03"},
		{ID: 3, EmployeeID: 3, Amount: 6000, PayDate: "2017-03"},
		{ID: 4, EmployeeID: 1, Amount: 7000, PayDate: "2017-04"},
		{ID: 5, EmployeeID: 2, Amount: 3000, PayDate: "2017-04"},
		{ID: 6, EmployeeID: 3, Amount: 7000, PayDate: "2017-04"},
	}

	fmt.Println("Salaries:")
	for _, s := range salaries {
		fmt.Printf("  Emp %d, Dept %d, $%d, %s\n",
			s.EmployeeID, employees[s.EmployeeID-1].DepartmentID, s.Amount, s.PayDate)
	}

	results := averageSalaryDepartmentsVsCompany(employees, salaries)
	fmt.Println("\nDepartment vs Company Average Comparison:")
	for _, r := range results {
		fmt.Printf("  %s | Dept %d | %s\n", r.PayMonth, r.DepartmentID, r.Comparison)
	}

	// --- Edge cases ---
	fmt.Println("\n--- Edge Cases ---")

	// Single department.
	emps2 := []Employee{{ID: 1, DepartmentID: 1}}
	sals2 := []SalaryRecord{
		{ID: 10, EmployeeID: 1, Amount: 5000, PayDate: "2020-01"},
	}
	fmt.Println("Single employee/dept:")
	for _, r := range averageSalaryDepartmentsVsCompany(emps2, sals2) {
		fmt.Printf("  %s | Dept %d | %s (expected same)\n", r.PayMonth, r.DepartmentID, r.Comparison)
	}

	// Same avg across depts.
	emps3 := []Employee{
		{ID: 10, DepartmentID: 1},
		{ID: 20, DepartmentID: 2},
	}
	sals3 := []SalaryRecord{
		{ID: 1, EmployeeID: 10, Amount: 1000, PayDate: "2020-01"},
		{ID: 2, EmployeeID: 20, Amount: 1000, PayDate: "2020-01"},
	}
	fmt.Println("Equal dept and company avg:")
	for _, r := range averageSalaryDepartmentsVsCompany(emps3, sals3) {
		fmt.Printf("  %s | Dept %d | %s (expected same)\n", r.PayMonth, r.DepartmentID, r.Comparison)
	}

	// Empty.
	fmt.Println("Empty:", len(averageSalaryDepartmentsVsCompany(nil, nil)))
}
```
