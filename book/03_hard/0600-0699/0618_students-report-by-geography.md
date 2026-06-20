# 0618 — Students Report By Geography

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func studentsReportByGeography(students []Student) map[string][]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(N log N) for sorting within each continent, Space: O(N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #618: Students Report By Geography
// https://leetcode.com/problems/students-report-by-geography/
// Difficulty: Hard [Paid]
//
// Pivot the Student table so that each continent becomes a column.
// Student names within each continent are listed alphabetically.
// Number of rows = maximum number of students in any continent.
// Empty cells are NULL (empty string).

// Student represents a row in the Student table.
type Student struct {
	Name      string
	Continent string
}

// studentsReportByGeography pivots students by continent.
// Time: O(N log N) for sorting within each continent, Space: O(N)
func studentsReportByGeography(students []Student) map[string][]string {
	// Group by continent, sort names within each group.
  // HashMap: O(1) lookup
	byContinent := make(map[string][]string)
	for _, s := range students {
		byContinent[s.Continent] = append(byContinent[s.Continent], s.Name)
	}

	// Sort names within each continent.
	for continent := range byContinent {
		sort.Strings(byContinent[continent])
	}

	// Build output columns.
  // HashMap: O(1) lookup
	result := make(map[string][]string)

	// Get sorted continent names for deterministic iteration.
	var continents []string
	for c := range byContinent {
		continents = append(continents, c)
	}
	sort.Strings(continents)

	// Find max rows needed.
	maxRows := 0
	for _, names := range byContinent {
		if len(names) > maxRows {
			maxRows = len(names)
		}
	}

	for _, continent := range continents {
		names := byContinent[continent]
		col := make([]string, maxRows)
		for i := 0; i < maxRows; i++ {
			if i < len(names) {
				col[i] = names[i]
			} else {
				col[i] = "" // NULL in SQL
			}
		}
		result[continent] = col
	}

	return result
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0618 Students Report By Geography ===")

	students := []Student{
		{Name: "Alice", Continent: "America"},
		{Name: "Bob", Continent: "America"},
		{Name: "Carol", Continent: "Asia"},
		{Name: "David", Continent: "Europe"},
		{Name: "Eve", Continent: "America"},
		{Name: "Frank", Continent: "Asia"},
		{Name: "Grace", Continent: "Europe"},
	}

	fmt.Println("Students:")
	for _, s := range students {
		fmt.Printf("  %s (%s)\n", s.Name, s.Continent)
	}

	result := studentsReportByGeography(students)

	fmt.Println("\nPivoted Report (ordered by continent):")
	// Print header.
	var continents []string
	for c := range result {
		continents = append(continents, c)
	}
	sort.Strings(continents)

	// Determine number of rows.
	maxRows := 0
	for _, col := range result {
		if len(col) > maxRows {
			maxRows = len(col)
		}
	}

	// Print table.
	for _, c := range continents {
		fmt.Printf("| %-10s ", c)
	}
	fmt.Println("|")
	// Separator.
	for range continents {
		fmt.Printf("|%s", "-----------")
	}
	fmt.Println("|")

	for row := 0; row < maxRows; row++ {
		for _, c := range continents {
			val := result[c][row]
			if val == "" {
				val = "NULL"
			}
			fmt.Printf("| %-10s ", val)
		}
		fmt.Println("|")
	}

	// --- Edge cases ---
	fmt.Println("\n--- Edge Cases ---")

	// Single continent.
	students2 := []Student{
		{Name: "Zoe", Continent: "Antarctica"},
		{Name: "Adam", Continent: "Antarctica"},
	}
	r2 := studentsReportByGeography(students2)
	fmt.Println("Single continent:")
	for cont, names := range r2 {
		fmt.Printf("  %s: %v\n", cont, names)
	}

	// Empty.
	r3 := studentsReportByGeography(nil)
	fmt.Println("Empty:", len(r3))

	// Single student per continent.
	students4 := []Student{
		{Name: "A", Continent: "X"},
		{Name: "B", Continent: "Y"},
		{Name: "C", Continent: "Z"},
	}
	r4 := studentsReportByGeography(students4)
	fmt.Println("One per continent:")
	for cont, names := range r4 {
		fmt.Printf("  %s: %v\n", cont, names)
	}
}
```
