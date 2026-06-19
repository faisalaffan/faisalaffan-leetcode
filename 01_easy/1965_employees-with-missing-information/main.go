package main

// LeetCode #1965: Employees With Missing Information
// https://leetcode.com/problems/employees-with-missing-information/
// Difficulty: Easy (SQL)

import (
	"fmt"
	"sort"
)

func main() {
	// Demonstration: employees (ID, name) and salaries (ID, salary)
	// Find employees missing name or salary
	employees := [][2]string{{"1", "Alice"}, {"2", "Bob"}, {"3", "Charlie"}}
	salaries := [][2]string{{"1", "5000"}, {"3", "6000"}}
	fmt.Println(EmployeesWithMissingInformation(employees, salaries)) // [2]
}

// Time: O(n log n), Space: O(n)
func EmployeesWithMissingInformation(employees, salaries [][2]string) []int {
	present := make(map[int]bool)
	for _, e := range employees {
		id := 0
		for _, c := range e[0] {
			id = id*10 + int(c-'0')
		}
		present[id] = true
	}
	for _, s := range salaries {
		id := 0
		for _, c := range s[0] {
			id = id*10 + int(c-'0')
		}
		present[id] = true
	}

	// IDs that appear in only one table
	empSet := make(map[int]bool)
	for _, e := range employees {
		id := 0
		for _, c := range e[0] {
			id = id*10 + int(c-'0')
		}
		empSet[id] = true
	}
	salSet := make(map[int]bool)
	for _, s := range salaries {
		id := 0
		for _, c := range s[0] {
			id = id*10 + int(c-'0')
		}
		salSet[id] = true
	}

	var result []int
	for id := range present {
		if empSet[id] != salSet[id] {
			result = append(result, id)
		}
	}
	sort.Ints(result)
	return result
}
