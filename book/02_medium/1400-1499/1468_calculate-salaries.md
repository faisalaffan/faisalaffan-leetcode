# 1468 — Calculate Salaries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func calculateSalaries(salaries []struct {
	companyID    int
	employeeID   int
	employeeName string
	salary       int
}, companies []struct {
	companyID int
	name      string
}) []taxResult
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n) for sorting  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1468: Calculate Salaries
// https://leetcode.com/problems/calculate-salaries/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// SQL problem - simulating in Go
	result := calculateSalaries(
		[]struct {
			companyID    int
			employeeID   int
			employeeName string
			salary       int
		}{
			{1, 1, "Alice", 100000},
			{1, 2, "Bob", 90000},
			{1, 3, "Charlie", 80000},
			{1, 4, "David", 70000},
			{1, 5, "Eve", 60000},
			{2, 6, "Frank", 120000},
			{2, 7, "Grace", 110000},
			{2, 8, "Henry", 100000},
		},
		[]struct {
			companyID int
			name      string
		}{
			{1, "Acme"},
			{2, "Globex"},
		},
	)
	for _, r := range result {
		fmt.Printf("%d %s %s %d\n", r.companyID, r.companyName, r.employeeName, r.tax)
	}
}

type taxResult struct {
	companyID     int
	companyName   string
	employeeName  string
	tax           int
}

// Time: O(n log n) for sorting
// Space: O(n)
func calculateSalaries(salaries []struct {
	companyID    int
	employeeID   int
	employeeName string
	salary       int
}, companies []struct {
	companyID int
	name      string
}) []taxResult {
	// Group salaries by company
  // Membuat map (HashMap) — pencarian O(1)
	companySalaries := make(map[int][]struct {
		employeeID int
		name       string
		salary     int
	})
	for _, s := range salaries {
		companySalaries[s.companyID] = append(companySalaries[s.companyID], struct {
			employeeID int
			name       string
			salary     int
		}{s.employeeID, s.employeeName, s.salary})
	}

  // Membuat map (HashMap) — pencarian O(1)
	companyNames := make(map[int]string)
	for _, c := range companies {
		companyNames[c.companyID] = c.name
	}

	var results []taxResult
	for companyID, emps := range companySalaries {
		// Sort by salary descending
  // Custom sort dengan comparator
		sort.Slice(emps, func(i, j int) bool {
			return emps[i].salary > emps[j].salary
		})

		totalEmp := len(emps)

		for _, emp := range emps {
			tax := emp.salary
			if emp.salary > 100000 {
				tax = emp.salary
			} else if emp.salary < 1000 {
				tax = 0
			} else {
				// Count employees with higher salary
				rank := 1
				for i := 0; i < totalEmp; i++ {
					if emps[i].salary > emp.salary {
						rank++
					}
				}
				// Max tax rate: 100000, reduced by 10% for each higher-rank employee
				maxTax := 100000
				reduction := (rank - 1) * 10000
				if rank >= 10 {
					tax = 0
				} else {
					tax = emp.salary
					mx := maxTax - reduction
					if tax > mx {
						tax = mx
					}
				}
			}
			if tax < 0 {
				tax = 0
			}

			results = append(results, taxResult{
				companyID,
				companyNames[companyID],
				emp.name,
				tax,
			})
		}
	}

	return results
}
```
