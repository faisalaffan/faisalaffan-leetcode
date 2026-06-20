# 0690 — Employee Importance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func getImportance(employees []*Employee, id int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #690: Employee Importance
// https://leetcode.com/problems/employee-importance/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	employees := []*Employee{
		{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
		{Id: 2, Importance: 3, Subordinates: []int{}},
		{Id: 3, Importance: 3, Subordinates: []int{}},
	}
	fmt.Println(getImportance(employees, 1))
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
  // HashMap: O(1) lookup
	empMap := make(map[int]*Employee)
	for _, e := range employees {
		empMap[e.Id] = e
	}

	var dfs func(id int) int
	dfs = func(id int) int {
		emp := empMap[id]
		total := emp.Importance
		for _, subId := range emp.Subordinates {
			total += dfs(subId)
		}
		return total
	}

	return dfs(id)
}
```
