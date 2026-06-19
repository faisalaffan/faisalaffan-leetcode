package main

// LeetCode #3057: Employees Project Allocation (SQL simulation)
// https://leetcode.com/problems/employees-project-allocation/
// Difficulty: Hard [Paid]

import (
	"fmt"
	"sort"
)

type Project struct {
	ProjectID  int
	EmployeeID int
	Workload   int
}
type Employee struct {
	EmployeeID int
	Name       string
	Team       string
}

type allocResult struct {
	EmployeeID      int
	ProjectID       int
	EmployeeName    string
	ProjectWorkload int
}

func employeesProjectAllocation(projects []Project, employees []Employee) []allocResult {
	empMap := make(map[int]Employee)
	for _, e := range employees {
		empMap[e.EmployeeID] = e
	}
	type teamSum struct {
		total int
		count int
	}
	teamStats := make(map[string]*teamSum)
	for _, p := range projects {
		emp, ok := empMap[p.EmployeeID]
		if !ok { continue }
		if teamStats[emp.Team] == nil {
			teamStats[emp.Team] = &teamSum{}
		}
		teamStats[emp.Team].total += p.Workload
		teamStats[emp.Team].count++
	}
	var result []allocResult
	for _, p := range projects {
		emp := empMap[p.EmployeeID]
		stats := teamStats[emp.Team]
		if stats == nil || stats.count == 0 { continue }
		avg := float64(stats.total) / float64(stats.count)
		if float64(p.Workload) > avg {
			result = append(result, allocResult{p.EmployeeID, p.ProjectID, emp.Name, p.Workload})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].EmployeeID != result[j].EmployeeID {
			return result[i].EmployeeID < result[j].EmployeeID
		}
		return result[i].ProjectID < result[j].ProjectID
	})
	return result
}

func main() {
	projects := []Project{
		{1, 1, 80}, {2, 1, 60}, {1, 2, 40}, {2, 3, 90},
	}
	employees := []Employee{
		{1, "Alice", "Engineering"},
		{2, "Bob", "Engineering"},
		{3, "Charlie", "Marketing"},
	}
	for _, r := range employeesProjectAllocation(projects, employees) {
		fmt.Printf("%d %d %s %d\n", r.EmployeeID, r.ProjectID, r.EmployeeName, r.ProjectWorkload)
	}
}
