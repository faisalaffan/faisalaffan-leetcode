package main

// LeetCode #1978: Employees Whose Manager Left the Company
// https://leetcode.com/problems/employees-whose-manager-left-the-company/
// Difficulty: Easy (SQL)

import (
	"fmt"
	"sort"
)

func main() {
	// Employees: (employee_id, manager_id, salary)
	employees := [][3]string{{"1", "3", "1000"}, {"2", "3", "2000"}, {"3", "", "3000"}, {"4", "5", "4000"}}
	fmt.Println(EmployeesWhoseManagerLeftTheCompany(employees)) // [4]
}

// Time: O(n log n), Space: O(n)
func EmployeesWhoseManagerLeftTheCompany(employees [][3]string) []int {
	empSet := make(map[int]bool)
	for _, e := range employees {
		id := 0
		for _, c := range e[0] {
			id = id*10 + int(c-'0')
		}
		empSet[id] = true
	}

	var result []int
	for _, e := range employees {
		id := 0
		for _, c := range e[0] {
			id = id*10 + int(c-'0')
		}
		managerID := 0
		if e[1] != "" {
			for _, c := range e[1] {
				managerID = managerID*10 + int(c-'0')
			}
		}
		salary := 0
		for _, c := range e[2] {
			salary = salary*10 + int(c-'0')
		}
		if managerID > 0 && !empSet[managerID] && salary < 30000 {
			result = append(result, id)
		}
	}
	sort.Ints(result)
	return result
}
