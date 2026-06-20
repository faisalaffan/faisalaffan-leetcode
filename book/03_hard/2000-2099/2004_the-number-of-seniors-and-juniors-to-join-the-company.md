# 2004 — The Number Of Seniors And Juniors To Join The Company

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func theNumberOfSeniorsAndJuniorsToJoinTheCompany(budget int, employees [][]interface{}) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2004: The Number of Seniors and Juniors to Join the Company
// https://leetcode.com/problems/the-number-of-seniors-and-juniors-to-join-the-company/
// Difficulty: Hard [Paid]
//
// Given tables Employees (employee_id, experience='Senior'|'Junior', salary)
// and a company budget, find the maximum number of employees to hire.
// Hiring rule: first hire as many Seniors as possible within budget,
// then hire as many Juniors as possible with the remaining budget.
//
// Return [senior_count, junior_count].

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: simple
	employees := [][]interface{}{
		{1, "Senior", 50000},
		{2, "Senior", 80000},
		{3, "Junior", 30000},
		{4, "Junior", 25000},
	}
	// Budget = 100000
	// Seniors: can hire 1 (50000) or (80000) -> 1 senior with min salary to maximize count
	// Actually "hire as many Seniors as possible" means maximize count
	// Senior: 50000 works, 80000 works but only 1. 50000 is better for remaining budget
	// With 50000 senior, remaining=50000, juniors: 30000+25000=55000 > 50000, so only 1 junior
	// Result: [1, 1]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompany(100000, employees))

	// Test case 2: hire many juniors
	employees = [][]interface{}{
		{1, "Senior", 10000},
		{2, "Senior", 20000},
		{3, "Junior", 5000},
		{4, "Junior", 5000},
		{5, "Junior", 5000},
	}
	// Budget = 25000
	// Seniors: hire 1 (10000) -> remaining 15000
	// Juniors: 3 * 5000 = 15000 -> 3 juniors
	// Result: [1, 3]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompany(25000, employees))

	// Test case 3: budget too low
	employees = [][]interface{}{
		{1, "Senior", 100000},
		{2, "Junior", 80000},
	}
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompany(50000, employees))

	// Test case 4: no budget
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompany(0, employees))

	// Test case 5: larger example
	employees = [][]interface{}{
		{1, "Senior", 40000},
		{2, "Senior", 30000},
		{3, "Senior", 60000},
		{4, "Junior", 15000},
		{5, "Junior", 10000},
		{6, "Junior", 20000},
	}
	// Budget = 70000
	// Seniors: can hire 2 (40000+30000=70000) -> remaining 0
	// No budget for juniors
	// Result: [2, 0]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompany(70000, employees))
}

func theNumberOfSeniorsAndJuniorsToJoinTheCompany(budget int, employees [][]interface{}) []int {
	var seniorSalaries, juniorSalaries []int
	for _, emp := range employees {
		salary := emp[2].(int)
		if emp[1].(string) == "Senior" {
			seniorSalaries = append(seniorSalaries, salary)
		} else {
			juniorSalaries = append(juniorSalaries, salary)
		}
	}

  // Sort O(n log n)
	sort.Ints(seniorSalaries)
  // Sort O(n log n)
	sort.Ints(juniorSalaries)

	// Hire as many Seniors as possible (maximize count)
	seniorCount := 0
	seniorCost := 0
	for _, s := range seniorSalaries {
		if seniorCost+s <= budget {
			seniorCost += s
			seniorCount++
		} else {
			break
		}
	}

	// Hire as many Juniors as possible with remaining budget
	remaining := budget - seniorCost
	juniorCount := 0
	juniorCost := 0
	for _, s := range juniorSalaries {
		if juniorCost+s <= remaining {
			juniorCost += s
			juniorCount++
		} else {
			break
		}
	}

	return []int{seniorCount, juniorCount}
}
```
