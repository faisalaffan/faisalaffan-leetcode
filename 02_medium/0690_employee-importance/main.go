package main

// LeetCode #690: Employee Importance
// https://leetcode.com/problems/employee-importance/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	employees := []*Employee{
		{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
		{Id: 2, Importance: 3, Subordinates: []int{}},
		{Id: 3, Importance: 3, Subordinates: []int{}},
	}
	fmt.Println(getImportance(employees, 1))
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	empMap := make(map[int]*Employee)
	for _, e := range employees {
		empMap[e.Id] = e
	}

	var dfs func(id int) int
	dfs = func(id int) int {
		emp := empMap[id]
		total := emp.Importance
		for _, subId := range emp.Subordinates {
			total += dfs(subId)
		}
		return total
	}

	return dfs(id)
}
