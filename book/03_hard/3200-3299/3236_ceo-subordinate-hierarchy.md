# 3236 — Ceo Subordinate Hierarchy

## Deskripsi

**Soal:** [3236. Ceo Subordinate Hierarchy](https://leetcode.com/problems/ceo-subordinate-hierarchy/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #3236: CEO Subordinate Hierarchy
// https://leetcode.com/problems/ceo-subordinate-hierarchy/
// Difficulty: Hard [Paid]
//
// Given an employee-manager hierarchy, return for each subordinate their
// hierarchy level and salary difference from the CEO.
// CEO has manager_id == -1 (or nil). The result is ordered by hierarchy_level
// ascending, then subordinate_id ascending.
//
// Tree DFS approach: build adjacency list from manager to subordinates, find
// the CEO (no manager), then traverse to compute levels and salary differences.

import (
	"fmt"
	"sort"
)

type Employee struct {
	ID     int
	Name   string
	Salary int
}

type Result struct {
	SubordinateID   int
	SubordinateName string
	HierarchyLevel  int
	SalaryDiff      int
}

func main() {
	// Example 1: Simple chain
	employees := []Employee{
		{1, "CEO", 100000},
		{2, "VP", 80000},
		{3, "Manager", 60000},
		{4, "Engineer", 40000},
	}
	managerID := map[int]int{
		1: -1, // CEO
		2: 1,  // reports to CEO
		3: 2,  // reports to VP
		4: 3,  // reports to Manager
	}
	fmt.Println(ceoSubordinateHierarchy(employees, managerID))

	// Example 2: Flat structure
	employees2 := []Employee{
		{1, "CEO", 200000},
		{2, "DirectorA", 150000},
		{3, "DirectorB", 140000},
		{4, "DirectorC", 130000},
	}
	managerID2 := map[int]int{
		1: -1,
		2: 1,
		3: 1,
		4: 1,
	}
	fmt.Println(ceoSubordinateHierarchy(employees2, managerID2))

	// Example 3: Multi-level tree
	employees3 := []Employee{
		{1, "CEO", 500000},
		{2, "EVP", 400000},
		{3, "SVP", 300000},
		{4, "VP", 200000},
		{5, "Director", 150000},
		{6, "Manager", 100000},
		{7, "IC", 80000},
	}
	managerID3 := map[int]int{
		1: -1,
		2: 1,
		3: 2,
		4: 3,
		5: 4,
		6: 5,
		7: 6,
	}
	fmt.Println(ceoSubordinateHierarchy(employees3, managerID3))

	// Example 4: Single employee (just CEO)
	employees4 := []Employee{
		{1, "CEO", 100000},
	}
	managerID4 := map[int]int{
		1: -1,
	}
	fmt.Println(ceoSubordinateHierarchy(employees4, managerID4))

	// Example 5: Two employees
	employees5 := []Employee{
		{1, "CEO", 100000},
		{2, "Assistant", 50000},
	}
	managerID5 := map[int]int{
		1: -1,
		2: 1,
	}
	fmt.Println(ceoSubordinateHierarchy(employees5, managerID5))
}

func ceoSubordinateHierarchy(employees []Employee, manager map[int]int) []Result {
  // Membuat map untuk pencarian O(1): key → value
	empMap := make(map[int]Employee)
  // Membuat map untuk pencarian O(1): key → value
	subordinates := make(map[int][]int)
	ceoID := -1

	for _, e := range employees {
		empMap[e.ID] = e
		mgr := manager[e.ID]
		if mgr == -1 {
			ceoID = e.ID
		} else {
			subordinates[mgr] = append(subordinates[mgr], e.ID)
		}
	}

	if ceoID == -1 {
		return nil
	}

	ceoSalary := empMap[ceoID].Salary

	var results []Result

	var dfs func(empID int, level int)
	dfs = func(empID int, level int) {
		emp := empMap[empID]
		if level > 0 {
			results = append(results, Result{
				SubordinateID:   emp.ID,
				SubordinateName: emp.Name,
				HierarchyLevel:  level,
				SalaryDiff:      emp.Salary - ceoSalary,
			})
		}
		for _, subID := range subordinates[empID] {
			dfs(subID, level+1)
		}
	}

	dfs(ceoID, 0)

	sort.Slice(results, func(i, j int) bool {
		if results[i].HierarchyLevel != results[j].HierarchyLevel {
			return results[i].HierarchyLevel < results[j].HierarchyLevel
		}
		return results[i].SubordinateID < results[j].SubordinateID
	})

	return results
}
```
