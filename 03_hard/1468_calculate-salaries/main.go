package main

// LeetCode #1468: Calculate Salaries
// https://leetcode.com/problems/calculate-salaries/
// Difficulty: Medium (listed here as Hard) [Paid]
//
// Calculate salary after tax. The tax rate depends on the company's
// max salary in the same department. If max salary in the department
// is < 1000, tax is 0%. If max is between 1000 and 10000 (inclusive),
// tax is 24%. If max is > 10000, tax is 49%.
// Also, salary after tax is rounded.

import "fmt"

// Employee represents an employee record.
type Employee struct {
	ID         int
	Department string
	Salary     int
}

// calculateSalaries computes post-tax salaries for all employees.
func calculateSalaries(employees []Employee) map[int]int {
	// Find max salary per department
	deptMax := make(map[string]int)
	for _, e := range employees {
		if e.Salary > deptMax[e.Department] {
			deptMax[e.Department] = e.Salary
		}
	}

	// Calculate tax rate per department and apply to each employee
	result := make(map[int]int)
	for _, e := range employees {
		maxSal := deptMax[e.Department]
		var taxRate float64
		if maxSal < 1000 {
			taxRate = 0.0
		} else if maxSal <= 10000 {
			taxRate = 0.24
		} else {
			taxRate = 0.49
		}

		afterTax := float64(e.Salary) * (1.0 - taxRate)
		result[e.ID] = roundSalary(afterTax)
	}

	return result
}

// roundSalary rounds to nearest integer, with .5 going up.
func roundSalary(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

func main() {
	employees := []Employee{
		{1, "Engineering", 8000},
		{2, "Engineering", 6000},
		{3, "Marketing", 500},
		{4, "Marketing", 700},
		{5, "Sales", 12000},
		{6, "Sales", 11000},
	}

	salaries := calculateSalaries(employees)
	fmt.Println("Salaries after tax:")
	employeeNames := map[int]string{1: "Alice", 2: "Bob", 3: "Charlie", 4: "David", 5: "Eve", 6: "Frank"}
	for _, e := range employees {
		fmt.Printf("  %s (%s): $%d -> $%d\n", employeeNames[e.ID], e.Department, e.Salary, salaries[e.ID])
	}

	// Test 2: Dept with max < 1000 (0% tax)
	employees2 := []Employee{
		{1, "Support", 800},
		{2, "Support", 900},
	}
	salaries2 := calculateSalaries(employees2)
	fmt.Println("\nTest 2 - Low salary dept (0% tax):")
	for _, e := range employees2 {
		fmt.Printf("  ID %d: $%d -> $%d (expected $%d)\n", e.ID, e.Salary, salaries2[e.ID], e.Salary)
	}

	// Test 3: Multiple departments with different rates
	employees3 := []Employee{
		{1, "DeptA", 500},    // max=500, 0% -> 500
		{2, "DeptA", 300},    // max=500, 0% -> 300
		{3, "DeptB", 5000},   // max=5000, 24% -> 3800
		{4, "DeptB", 2000},   // max=5000, 24% -> 1520
		{5, "DeptC", 20000},  // max=20000, 49% -> 10200
		{6, "DeptC", 15000},  // max=20000, 49% -> 7650
	}
	salaries3 := calculateSalaries(employees3)
	fmt.Println("\nTest 3 - Mixed departments:")
	for _, e := range employees3 {
		fmt.Printf("  ID %d (dept %s): $%d -> $%d\n", e.ID, e.Department, e.Salary, salaries3[e.ID])
	}
}
