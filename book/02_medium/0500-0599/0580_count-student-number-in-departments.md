# 0580 — Count Student Number In Departments

## Deskripsi

**Soal:** [0580. Count Student Number In Departments](https://leetcode.com/problems/count-student-number-in-departments/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m) where n = departments, m = students  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #580: Count Student Number in Departments
// https://leetcode.com/problems/count-student-number-in-departments/
// Difficulty: Medium [Paid]
// Time: O(n + m) where n = departments, m = students
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	// departments: {dept_id, dept_name}
	departments := map[int]string{
		1: "Engineering",
		2: "Science",
		3: "Arts",
	}
	// students: {student_id, dept_id}
	students := []int{1, 1, 1, 2, 3}

	result := CountStudents(departments, students)
	for _, r := range result {
		fmt.Printf("%s: %d\n", r[0].(string), r[1].(int))
	}
}

type DeptCount struct {
	Name  string
	Count int
}

func CountStudents(departments map[int]string, studentDepts []int) [][]interface{} {
  // Membuat map untuk pencarian O(1): key → value
	counts := make(map[int]int)
	for _, deptID := range studentDepts {
		counts[deptID]++
	}

	result := [][]interface{}{}
	for deptID, deptName := range departments {
		result = append(result, []interface{}{deptName, counts[deptID]})
	}
	// Include departments with 0 students
	for deptID := range departments {
		if _, exists := counts[deptID]; !exists {
			counts[deptID] = 0
		}
	}

	// Rebuild with all departments
	result = nil
	for deptID, deptName := range departments {
		result = append(result, []interface{}{deptName, counts[deptID]})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0].(string) < result[j][0].(string)
	})

	return result
}
```
