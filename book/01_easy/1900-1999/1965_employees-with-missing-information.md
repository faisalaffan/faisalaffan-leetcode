# 1965 — Employees With Missing Information

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func EmployeesWithMissingInformation(employees, salaries [][2]string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1965: Employees With Missing Information
// https://leetcode.com/problems/employees-with-missing-information/
// Difficulty: Easy (SQL)

import (
	"fmt"
	"sort"
)

func main() {
	// Demonstration: employees (ID, name) and salaries (ID, salary)
	// Find employees missing name or salary
	employees := [][2]string{{"1", "Alice"}, {"2", "Bob"}, {"3", "Charlie"}}
	salaries := [][2]string{{"1", "5000"}, {"3", "6000"}}
	fmt.Println(EmployeesWithMissingInformation(employees, salaries)) // [2]
}

// Time: O(n log n), Space: O(n)
func EmployeesWithMissingInformation(employees, salaries [][2]string) []int {
  // HashMap: O(1) lookup
	present := make(map[int]bool)
	for _, e := range employees {
		id := 0
		for _, c := range e[0] {
			id = id*10 + int(c-'0')
		}
		present[id] = true
	}
	for _, s := range salaries {
		id := 0
		for _, c := range s[0] {
			id = id*10 + int(c-'0')
		}
		present[id] = true
	}

	// IDs that appear in only one table
  // HashMap: O(1) lookup
	empSet := make(map[int]bool)
	for _, e := range employees {
		id := 0
		for _, c := range e[0] {
			id = id*10 + int(c-'0')
		}
		empSet[id] = true
	}
  // HashMap: O(1) lookup
	salSet := make(map[int]bool)
	for _, s := range salaries {
		id := 0
		for _, c := range s[0] {
			id = id*10 + int(c-'0')
		}
		salSet[id] = true
	}

	var result []int
	for id := range present {
		if empSet[id] != salSet[id] {
			result = append(result, id)
		}
	}
  // Sort O(n log n)
	sort.Ints(result)
	return result
}
```
