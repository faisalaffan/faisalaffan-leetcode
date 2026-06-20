# 2010 — The Number Of Seniors And Juniors To Join The Company Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(budget int, employees [][]interface{}) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2010: The Number of Seniors and Juniors to Join the Company II
// https://leetcode.com/problems/the-number-of-seniors-and-juniors-to-join-the-company-ii/
// Difficulty: Hard [Paid]
//
// Variation of 2004. Instead of hiring as many Seniors as possible first,
// we need to find the maximum TOTAL number of employees we can hire within
// budget. We can choose any combination of Seniors and Juniors.
//
// Return [senior_count, junior_count] for the optimal hiring strategy
// that maximizes total headcount. If there are multiple solutions with
// the same total, prefer the one with more Seniors (or less Juniors).

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
	// Option A: 1 Senior (50000) + 1 Junior (25000 or 30000) = 2 employees
	// Option B: 0 Seniors + 3 Juniors (25000+30000) = 55000, remaining 45000 can't afford more = 2
	// Option C: 2 Seniors (50000+80000) = 130000 > budget, no
	// Max total = 2. Prefer more Seniors: [1, 1]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(100000, employees))

	// Test case 2: more juniors gives higher total
	employees = [][]interface{}{
		{1, "Senior", 100000},
		{2, "Senior", 60000},
		{3, "Junior", 15000},
		{4, "Junior", 15000},
		{5, "Junior", 15000},
	}
	// Budget = 100000
	// Option A: 1 Senior (60000) + 2 Juniors (30000) = 3 employees
	// Option B: 0 Seniors + 6 Juniors... only have 3 juniors = 3 employees
	// Option C: 1 Senior (100000) + 0 Juniors = 1 employee
	// Max total = 3 with more Seniors: [1, 2]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(100000, employees))

	// Test case 3: tie-breaking
	employees = [][]interface{}{
		{1, "Senior", 30000},
		{2, "Junior", 30000},
	}
	// Budget = 60000
	// Option A: 1 Senior (30000) + 1 Junior (30000) = 2 employees
	// Option B: 2 Seniors... only 1 senior = 1
	// Option C: 2 Juniors... only 1 junior = 1
	// [1, 1]
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(60000, employees))

	// Test case 4: budget not enough for any
	employees = [][]interface{}{
		{1, "Senior", 50000},
		{2, "Junior", 30000},
	}
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(20000, employees))

	// Test case 5: complex case
	employees = [][]interface{}{
		{1, "Senior", 10000},
		{2, "Senior", 20000},
		{3, "Senior", 30000},
		{4, "Junior", 5000},
		{5, "Junior", 5000},
		{6, "Junior", 5000},
		{7, "Junior", 10000},
	}
	// Budget = 40000
	// Seniors: [10000, 20000, 30000]
	// Juniors: [5000, 5000, 5000, 10000]
	// Try various combinations to maximize total count:
	// Option: 1 Senior (10000) + up to 6 juniors... max juniors with 30000: 3*5000+10000=25000 (4)
	//   Total: 5
	// Option: 2 Seniors (10000+20000=30000) + remaining 10000: 2*5000=10000 (2 juniors)
	//   Total: 4
	// Option: 0 Seniors + 4 Juniors (5000*3+10000=25000) -> remaining 15000 -> 4 total
	// Option: 3 Seniors: 60000 > budget
	// Best: 1 Senior + 4 Juniors = 5 total
	fmt.Println(theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(40000, employees))
}

func theNumberOfSeniorsAndJuniorsToJoinTheCompanyIi(budget int, employees [][]interface{}) []int {
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

	// Prefix sums for quick cost calculation
  // Alokasi slice
	seniorPrefix := make([]int, len(seniorSalaries)+1)
	for i, s := range seniorSalaries {
		seniorPrefix[i+1] = seniorPrefix[i] + s
	}

  // Alokasi slice
	juniorPrefix := make([]int, len(juniorSalaries)+1)
	for i, s := range juniorSalaries {
		juniorPrefix[i+1] = juniorPrefix[i] + s
	}

	bestTotal := 0
	bestSenior := 0
	bestJunior := 0

	// Try all possible senior counts
	for s := 0; s <= len(seniorSalaries); s++ {
		seniorCost := seniorPrefix[s]
		if seniorCost > budget {
			break
		}
		remaining := budget - seniorCost

		// Maximum juniors with remaining budget
		j := 0
		for j < len(juniorSalaries) && juniorPrefix[j+1] <= remaining {
			j++
		}

		total := s + j
		if total > bestTotal || (total == bestTotal && s > bestSenior) {
			bestTotal = total
			bestSenior = s
			bestJunior = j
		}
	}

	return []int{bestSenior, bestJunior}
}
```
