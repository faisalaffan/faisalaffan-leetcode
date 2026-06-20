# 2668 — Find Latest Salaries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindLatestSalaries(salaries []EmployeeSalary) []EmployeeSalary
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2668: Find Latest Salaries
// https://leetcode.com/problems/find-latest-salaries/
// Difficulty: Easy [Paid]
// Time: O(n log n) | Space: O(n)
// Note: SQL/JS problem adapted to Go. Find latest salary per department.

import (
	"fmt"
	"sort"
)

func main() {
	salaries := []EmployeeSalary{
		{1, "Alice", 50000, "Eng", "2023-01-01"},
		{1, "Alice", 55000, "Eng", "2023-06-01"},
		{2, "Bob", 60000, "Sales", "2023-01-01"},
	}
	fmt.Println(FindLatestSalaries(salaries))
}

type EmployeeSalary struct {
	EmpID  int
	Name   string
	Salary int
	Dept   string
	Date   string
}

func FindLatestSalaries(salaries []EmployeeSalary) []EmployeeSalary {
	if len(salaries) == 0 {
		return nil
	}

  // Custom sort dengan comparator
	sort.Slice(salaries, func(i, j int) bool {
		if salaries[i].EmpID != salaries[j].EmpID {
			return salaries[i].EmpID < salaries[j].EmpID
		}
		return salaries[i].Date > salaries[j].Date
	})

	result := []EmployeeSalary{}
	seen := map[int]bool{}
	for _, s := range salaries {
		if !seen[s.EmpID] {
			seen[s.EmpID] = true
			result = append(result, s)
		}
	}
	return result
}
```
