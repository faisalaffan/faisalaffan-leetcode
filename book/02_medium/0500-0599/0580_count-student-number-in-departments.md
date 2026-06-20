# 0580 — Count Student Number In Departments

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountStudents(departments map[int]string, studentDepts []int) [][]interface
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n + m) where n = departments, m = students  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
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

  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		return result[i][0].(string) < result[j][0].(string)
	})

	return result
}
```
