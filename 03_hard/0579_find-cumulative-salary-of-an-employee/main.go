package main

import (
	"fmt"
	"sort"
)

// LeetCode #579: Find Cumulative Salary of an Employee
// https://leetcode.com/problems/find-cumulative-salary-of-an-employee/
// Difficulty: Hard [Paid]
//
// For each employee, for each month except the most recent, compute the cumulative
// salary over the last 3 months (the current month and the two preceding months).
// Employees who worked fewer than 3 months should still have a cumulative sum
// of whatever months are available.

// EmployeeMonth represents a row in the Employee table.
type EmployeeMonth struct {
	ID     int
	Month  int  // 1..12
	Salary int
}

// CumulativeSalary holds one result row.
type CumulativeSalary struct {
	ID              int
	Month           int
	CumulativeSalary int
}

// findCumulativeSalaryOfAnEmployee computes 3-month running total for each employee
// excluding their most recent month.
// Time: O(E log E) for sorting, Space: O(E)
func findCumulativeSalaryOfAnEmployee(records []EmployeeMonth) []CumulativeSalary {
	// Group by employee ID.
	byID := make(map[int][]EmployeeMonth)
	for _, r := range records {
		byID[r.ID] = append(byID[r.ID], r)
	}

	var result []CumulativeSalary

	for id, emps := range byID {
		// Sort by month ascending.
		sort.Slice(emps, func(i, j int) bool {
			return emps[i].Month < emps[j].Month
		})

		if len(emps) == 0 {
			continue
		}

		// Exclude the most recent month.
		lastMonth := emps[len(emps)-1].Month

		// Build a month->salary map for O(1) lookups.
		salaryByMonth := make(map[int]int)
		for _, e := range emps {
			salaryByMonth[e.Month] = e.Salary
		}

		for _, e := range emps {
			if e.Month == lastMonth {
				continue // skip the most recent month
			}
			cum := e.Salary
			if s, ok := salaryByMonth[e.Month-1]; ok {
				cum += s
			}
			if s, ok := salaryByMonth[e.Month-2]; ok {
				cum += s
			}
			result = append(result, CumulativeSalary{
				ID:               id,
				Month:            e.Month,
				CumulativeSalary: cum,
			})
		}
	}

	// Sort output by ID asc, then month desc (as per problem spec).
	sort.Slice(result, func(i, j int) bool {
		if result[i].ID != result[j].ID {
			return result[i].ID < result[j].ID
		}
		return result[i].Month > result[j].Month
	})

	return result
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0579 Find Cumulative Salary of an Employee ===")

	records := []EmployeeMonth{
		{ID: 1, Month: 1, Salary: 2000},
		{ID: 1, Month: 2, Salary: 3000},
		{ID: 1, Month: 3, Salary: 4000},
		{ID: 1, Month: 4, Salary: 5000},
		{ID: 2, Month: 1, Salary: 2500},
		{ID: 2, Month: 2, Salary: 3500},
	}

	fmt.Println("Employee Records:")
	for _, r := range records {
		fmt.Printf("  Emp %d, Month %d, $%d\n", r.ID, r.Month, r.Salary)
	}

	results := findCumulativeSalaryOfAnEmployee(records)
	fmt.Println("\nCumulative Salary (excluding most recent month):")
	for _, r := range results {
		fmt.Printf("  Emp %d, Month %d, Cumulative $%d\n", r.ID, r.Month, r.CumulativeSalary)
	}

	// --- Edge cases ---
	fmt.Println("\n--- Edge Cases ---")

	// Single month (should be excluded since it's the most recent).
	recs2 := []EmployeeMonth{
		{ID: 10, Month: 5, Salary: 5000},
	}
	r2 := findCumulativeSalaryOfAnEmployee(recs2)
	fmt.Println("Single month (should be empty):", len(r2))

	// Two months.
	recs3 := []EmployeeMonth{
		{ID: 20, Month: 1, Salary: 1000},
		{ID: 20, Month: 2, Salary: 2000},
	}
	fmt.Println("Two months:")
	for _, r := range findCumulativeSalaryOfAnEmployee(recs3) {
		fmt.Printf("  Emp %d, Month %d, $%d\n", r.ID, r.Month, r.CumulativeSalary)
	}

	// Non-consecutive months.
	recs4 := []EmployeeMonth{
		{ID: 30, Month: 1, Salary: 1000},
		{ID: 30, Month: 3, Salary: 3000},
		{ID: 30, Month: 6, Salary: 6000},
	}
	fmt.Println("Non-consecutive months:")
	for _, r := range findCumulativeSalaryOfAnEmployee(recs4) {
		fmt.Printf("  Emp %d, Month %d, $%d\n", r.ID, r.Month, r.CumulativeSalary)
	}

	// Empty.
	fmt.Println("Empty:", len(findCumulativeSalaryOfAnEmployee(nil)))
}
