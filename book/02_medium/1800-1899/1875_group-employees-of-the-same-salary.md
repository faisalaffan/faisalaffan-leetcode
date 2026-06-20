# 1875 — Group Employees Of The Same Salary

## Deskripsi

**Soal:** [1875. Group Employees Of The Same Salary](https://leetcode.com/problems/group-employees-of-the-same-salary/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1875: Group Employees of the Same Salary
// https://leetcode.com/problems/group-employees-of-the-same-salary/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// employee: [id, salary]
	employees := [][]int{{1, 50000}, {2, 60000}, {3, 50000}, {4, 70000}, {5, 60000}}
	fmt.Println(GroupEmployees(employees))
}

// Time: O(n log n), Space: O(n)
func GroupEmployees(employees [][]int) [][]int {
  // Membuat map untuk pencarian O(1): key → value
	salaryMap := make(map[int][]int)
	for _, emp := range employees {
		id, salary := emp[0], emp[1]
		salaryMap[salary] = append(salaryMap[salary], id)
	}

  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, 0)
	for _, ids := range salaryMap {
		if len(ids) >= 2 {
			sort.Ints(ids)
			result = append(result, ids)
		}
	}

	// Sort by first employee ID for deterministic output
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}
```
