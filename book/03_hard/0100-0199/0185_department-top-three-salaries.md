# 0185 — Department Top Three Salaries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func departmentTopThreeSalaries(employees []Employee, departments []Department) []Result
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(E log E + D log D) for sorting, Space: O(E + D)  
**Kompleksitas Ruang:** O(E + D)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #185: Department Top Three Salaries
// https://leetcode.com/problems/department-top-three-salaries/
// Difficulty: Hard
//
// For each department, find employees who earn one of the top 3 distinct salaries.
// If there are ties, include all employees with those salaries.

// Employee represents a row in the Employee table.
type Employee struct {
	ID           int
	Name         string
	Salary       int
	DepartmentID int
}

// Department represents a row in the Department table.
type Department struct {
	ID   int
	Name string
}

// Result represents one output row.
type Result struct {
	Department string
	Employee   string
	Salary     int
}

// departmentTopThreeSalaries returns employees with top 3 distinct salaries per dept.
// Time: O(E log E + D log D) for sorting, Space: O(E + D)
func departmentTopThreeSalaries(employees []Employee, departments []Department) []Result {
	// Build department name lookup.
  // Membuat map (HashMap) — pencarian O(1)
	deptName := make(map[int]string)
	for _, d := range departments {
		deptName[d.ID] = d.Name
	}

	// Group employees by department.
  // Membuat map (HashMap) — pencarian O(1)
	byDept := make(map[int][]Employee)
	for _, e := range employees {
		byDept[e.DepartmentID] = append(byDept[e.DepartmentID], e)
	}

	var results []Result

	for deptID, emps := range byDept {
		// Sort descending by salary.
  // Custom sort dengan comparator
		sort.Slice(emps, func(i, j int) bool {
			return emps[i].Salary > emps[j].Salary
		})

		// Collect top 3 distinct salaries.
  // Alokasi slice integer
		distinctSalaries := make([]int, 0)
		for _, e := range emps {
			if len(distinctSalaries) == 0 || e.Salary != distinctSalaries[len(distinctSalaries)-1] {
				distinctSalaries = append(distinctSalaries, e.Salary)
				if len(distinctSalaries) == 3 {
					break
				}
			}
		}

		// Build a set of qualifying salaries.
  // Membuat map (HashMap) — pencarian O(1)
		qualifying := make(map[int]bool)
		for _, s := range distinctSalaries {
			qualifying[s] = true
		}

		// Gather employees whose salary is in the qualifying set.
		// Sort by salary desc, then name asc for stable output.
		var matched []Employee
		for _, e := range emps {
			if qualifying[e.Salary] {
				matched = append(matched, e)
			}
		}
  // Custom sort dengan comparator
		sort.Slice(matched, func(i, j int) bool {
			if matched[i].Salary != matched[j].Salary {
				return matched[i].Salary > matched[j].Salary
			}
			return matched[i].Name < matched[j].Name
		})

		for _, e := range matched {
			results = append(results, Result{
				Department: deptName[deptID],
				Employee:   e.Name,
				Salary:     e.Salary,
			})
		}
	}

	// Sort by department name for deterministic output.
  // Custom sort dengan comparator
	sort.Slice(results, func(i, j int) bool {
		if results[i].Department != results[j].Department {
			return results[i].Department < results[j].Department
		}
		if results[i].Salary != results[j].Salary {
			return results[i].Salary > results[j].Salary
		}
		return results[i].Employee < results[j].Employee
	})

	return results
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0185 Department Top Three Salaries ===")

	departments := []Department{
		{ID: 1, Name: "IT"},
		{ID: 2, Name: "Sales"},
	}

	employees := []Employee{
		{ID: 1, Name: "Joe", Salary: 85000, DepartmentID: 1},
		{ID: 2, Name: "Henry", Salary: 80000, DepartmentID: 2},
		{ID: 3, Name: "Sam", Salary: 60000, DepartmentID: 2},
		{ID: 4, Name: "Max", Salary: 90000, DepartmentID: 1},
		{ID: 5, Name: "Janet", Salary: 69000, DepartmentID: 1},
		{ID: 6, Name: "Randy", Salary: 85000, DepartmentID: 1},
		{ID: 7, Name: "Will", Salary: 70000, DepartmentID: 1},
	}

	fmt.Println("Employees:")
	for _, e := range employees {
		fmt.Printf("  %s (Dept %d, Salary %d)\n", e.Name, e.DepartmentID, e.Salary)
	}

	fmt.Println("\nTop 3 Salaries Per Department:")
	results := departmentTopThreeSalaries(employees, departments)
	for _, r := range results {
		fmt.Printf("  %s | %s | %d\n", r.Department, r.Employee, r.Salary)
	}

	// --- Edge cases ---
	fmt.Println("\n--- Edge Cases ---")

	// Single employee department.
	deps2 := []Department{
		{ID: 10, Name: "Engineering"},
		{ID: 20, Name: "HR"},
	}
	emps2 := []Employee{
		{ID: 1, Name: "Alice", Salary: 100000, DepartmentID: 10},
		{ID: 2, Name: "Bob", Salary: 50000, DepartmentID: 20},
		{ID: 3, Name: "Carol", Salary: 45000, DepartmentID: 20},
	}
	fmt.Println("Test - single employee dept:")
	for _, r := range departmentTopThreeSalaries(emps2, deps2) {
		fmt.Printf("  %s | %s | %d\n", r.Department, r.Employee, r.Salary)
	}

	// All same salary (ties).
	deps3 := []Department{{ID: 1, Name: "Support"}}
	emps3 := []Employee{
		{ID: 10, Name: "A", Salary: 50000, DepartmentID: 1},
		{ID: 20, Name: "B", Salary: 50000, DepartmentID: 1},
		{ID: 30, Name: "C", Salary: 50000, DepartmentID: 1},
		{ID: 40, Name: "D", Salary: 50000, DepartmentID: 1},
	}
	fmt.Println("Test - all same salary (all qualify):")
	for _, r := range departmentTopThreeSalaries(emps3, deps3) {
		fmt.Printf("  %s | %s | %d\n", r.Department, r.Employee, r.Salary)
	}

	// Empty departments.
	fmt.Println("Test - empty employees:", len(departmentTopThreeSalaries(nil, departments)))
}
```
