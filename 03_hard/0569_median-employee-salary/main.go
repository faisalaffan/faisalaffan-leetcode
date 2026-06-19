package main

import (
	"fmt"
	"sort"
)

// LeetCode #569: Median Employee Salary
// https://leetcode.com/problems/median-employee-salary/
// Difficulty: Hard [Paid]
//
// For each company, find the median salary of its employees.
// The median can be one or two values (for odd/even counts).
// Output: company, salary (one row per median value).

// Employee represents a row in the Employee table.
type Employee struct {
	ID      int
	Company string
	Salary  int
}

// medianEmployeeSalary returns median salary(s) per company.
// Time: O(E log E) for sorting, Space: O(E)
func medianEmployeeSalary(employees []Employee) []struct {
	Company string
	Salary  int
} {
	// Group by company.
	byCompany := make(map[string][]int)
	for _, e := range employees {
		byCompany[e.Company] = append(byCompany[e.Company], e.Salary)
	}

	type result struct {
		Company string
		Salary  int
	}
	var results []result

	// Process each company.
	for company, salaries := range byCompany {
		sort.Ints(salaries)
		n := len(salaries)

		if n == 0 {
			continue
		}

		if n%2 == 1 {
			// Odd count: one median.
			results = append(results, result{company, salaries[n/2]})
		} else {
			// Even count: two medians (the two middle values).
			results = append(results, result{company, salaries[n/2-1]})
			results = append(results, result{company, salaries[n/2]})
		}
	}

	// Sort by company name, then salary for deterministic output.
	sort.Slice(results, func(i, j int) bool {
		if results[i].Company != results[j].Company {
			return results[i].Company < results[j].Company
		}
		return results[i].Salary < results[j].Salary
	})

	return results
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0569 Median Employee Salary ===")

	employees := []Employee{
		{ID: 1, Company: "A", Salary: 1000},
		{ID: 2, Company: "A", Salary: 2000},
		{ID: 3, Company: "A", Salary: 3000},
		{ID: 4, Company: "B", Salary: 4000},
		{ID: 5, Company: "B", Salary: 5000},
		{ID: 6, Company: "B", Salary: 6000},
		{ID: 7, Company: "B", Salary: 7000},
	}

	fmt.Println("Employees:")
	for _, e := range employees {
		fmt.Printf("  %s $%d\n", e.Company, e.Salary)
	}

	results := medianEmployeeSalary(employees)
	fmt.Println("\nMedian Salaries Per Company:")
	for _, r := range results {
		fmt.Printf("  %s $%d\n", r.Company, r.Salary)
	}

	// --- Edge cases ---
	fmt.Println("\n--- Edge Cases ---")

	// Single employee per company.
	emps2 := []Employee{
		{ID: 10, Company: "X", Salary: 5000},
		{ID: 20, Company: "Y", Salary: 7000},
	}
	fmt.Println("Single employee companies:")
	for _, r := range medianEmployeeSalary(emps2) {
		fmt.Printf("  %s $%d\n", r.Company, r.Salary)
	}

	// Two employees in same company.
	emps3 := []Employee{
		{ID: 1, Company: "C", Salary: 1000},
		{ID: 2, Company: "C", Salary: 9000},
	}
	fmt.Println("Two-employee company (should return both):")
	for _, r := range medianEmployeeSalary(emps3) {
		fmt.Printf("  %s $%d\n", r.Company, r.Salary)
	}

	// Empty input.
	fmt.Println("Empty input:", len(medianEmployeeSalary(nil)), "results")

	// All same salary.
	emps4 := []Employee{
		{ID: 1, Company: "D", Salary: 5000},
		{ID: 2, Company: "D", Salary: 5000},
		{ID: 3, Company: "D", Salary: 5000},
	}
	fmt.Println("All same salary (odd):")
	for _, r := range medianEmployeeSalary(emps4) {
		fmt.Printf("  %s $%d\n", r.Company, r.Salary)
	}
}
