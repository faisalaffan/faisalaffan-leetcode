# 0580 — Count Student Number In Departments

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountStudents(departments map[int]string, studentDepts []int) [][]interface`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n + m) where n = departments, m = students  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
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

  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		return result[i][0].(string) < result[j][0].(string)
	})

	return result
}
```
