# 3057 — Employees Project Allocation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func employeesProjectAllocation(projects []Project, employees []Employee) []allocResult`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3057: Employees Project Allocation (SQL simulation)
// https://leetcode.com/problems/employees-project-allocation/
// Difficulty: Hard [Paid]
//
// Approach: Find employees whose workload exceeds the average workload of their team.
// Assignments are grouped by team to compute averages, then filtered by > avg condition.

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
  // HashMap: O(1) lookup
	empMap := make(map[int]Employee)
	for _, e := range employees {
		empMap[e.EmployeeID] = e
	}
	type teamSum struct {
		total int
		count int
	}
  // HashMap: O(1) lookup
	teamStats := make(map[string]*teamSum)
	for _, p := range projects {
		emp, ok := empMap[p.EmployeeID]
		if !ok {
			continue
		}
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
		if stats == nil || stats.count == 0 {
			continue
		}
		avg := float64(stats.total) / float64(stats.count)
		if float64(p.Workload) > avg {
			result = append(result, allocResult{p.EmployeeID, p.ProjectID, emp.Name, p.Workload})
		}
	}
  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		if result[i].EmployeeID != result[j].EmployeeID {
			return result[i].EmployeeID < result[j].EmployeeID
		}
		return result[i].ProjectID < result[j].ProjectID
	})
	return result
}

func main() {
	// Example 1: mixed teams
	projects1 := []Project{
		{1, 1, 80}, {2, 1, 60}, {1, 2, 40}, {2, 3, 90},
	}
	employees1 := []Employee{
		{1, "Alice", "Engineering"},
		{2, "Bob", "Engineering"},
		{3, "Charlie", "Marketing"},
	}
	fmt.Println("Example 1:")
	for _, r := range employeesProjectAllocation(projects1, employees1) {
		fmt.Printf("  Emp=%d Proj=%d Name=%s Workload=%d\n", r.EmployeeID, r.ProjectID, r.EmployeeName, r.ProjectWorkload)
	}
	// Engineering avg: (80+60+40)/3 = 60, Alice P1(80)>60, Alice P2(60)=60 not >60, Bob P1(40)<60
	// Marketing avg: 90/1 = 90, Charlie P2(90)=90 not >90

	// Example 2: single employee team
	projects2 := []Project{
		{1, 1, 100}, {2, 1, 50},
	}
	employees2 := []Employee{
		{1, "Dave", "Sales"},
	}
	fmt.Println("Example 2 (single team):")
	for _, r := range employeesProjectAllocation(projects2, employees2) {
		fmt.Printf("  Emp=%d Proj=%d Name=%s Workload=%d\n", r.EmployeeID, r.ProjectID, r.EmployeeName, r.ProjectWorkload)
	}
	// Sales avg: (100+50)/2 = 75, Dave P1(100)>75

	// Example 3: no matching employees
	projects3 := []Project{
		{1, 99, 10},
	}
	employees3 := []Employee{
		{1, "Eve", "Engineering"},
	}
	fmt.Println("Example 3 (no match):")
	for _, r := range employeesProjectAllocation(projects3, employees3) {
		fmt.Printf("  Emp=%d Proj=%d Name=%s Workload=%d\n", r.EmployeeID, r.ProjectID, r.EmployeeName, r.ProjectWorkload)
	}

	// Example 4: all below average
	projects4 := []Project{
		{1, 1, 5}, {1, 2, 10},
	}
	employees4 := []Employee{
		{1, "Frank", "QA"}, {2, "Grace", "QA"},
	}
	fmt.Println("Example 4 (all below avg):")
	for _, r := range employeesProjectAllocation(projects4, employees4) {
		fmt.Printf("  Emp=%d Proj=%d Name=%s Workload=%d\n", r.EmployeeID, r.ProjectID, r.EmployeeName, r.ProjectWorkload)
	}
	// QA avg: (5+10)/2 = 7.5, both below or equal
}
```
