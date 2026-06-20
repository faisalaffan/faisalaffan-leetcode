# 1965 — Employees With Missing Information

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func EmployeesWithMissingInformation(employees, salaries [][2]string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
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
  // Membuat map (HashMap) — pencarian O(1)
	empSet := make(map[int]bool)
	for _, e := range employees {
		id := 0
		for _, c := range e[0] {
			id = id*10 + int(c-'0')
		}
		empSet[id] = true
	}
  // Membuat map (HashMap) — pencarian O(1)
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
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}
```
