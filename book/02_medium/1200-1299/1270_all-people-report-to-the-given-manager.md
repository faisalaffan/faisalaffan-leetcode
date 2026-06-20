# 1270 — All People Report To The Given Manager

## Deskripsi

**Soal:** [1270. All People Report To The Given Manager](https://leetcode.com/problems/all-people-report-to-the-given-manager/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func allPeopleReportTo(employees [][]int) []int`

## Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1270: All People Report to the Given Manager
// https://leetcode.com/problems/all-people-report-to-the-given-manager/
// Difficulty: Medium [Paid]

// Find all employees who directly or indirectly report to the head
// (manager_id = 1 or manager_id is in the reporting chain).

// Time: O(n)
// Space: O(n)

func allPeopleReportTo(employees [][]int) []int {
	// employees[i] = [employee_id, manager_id]
	// Find all employees who report to employee_id=1 (directly or indirectly)

  // Membuat map untuk pencarian O(1): key → value
	adj := make(map[int][]int)
	for _, e := range employees {
		empID, mgrID := e[0], e[1]
		if mgrID != 0 { // 0 means no manager (head)
			adj[mgrID] = append(adj[mgrID], empID)
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0)
	queue := []int{1}

	for len(queue) > 0 {
		mgr := queue[0]
		queue = queue[1:]
		for _, emp := range adj[mgr] {
			result = append(result, emp)
			queue = append(queue, emp)
		}
	}

	sort.Ints(result)
	return result
}

func main() {
	employees := [][]int{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 3},
	}
	fmt.Printf("%v (expected: [2 3 4 5])\n", allPeopleReportTo(employees))

	employees2 := [][]int{
		{1, 0},
		{2, 1},
		{3, 1},
	}
	fmt.Printf("%v (expected: [2 3])\n", allPeopleReportTo(employees2))
}
```
