# 1204 — Last Person To Fit In The Bus

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func lastToFit(people []person) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1204: Last Person to Fit in the Bus
// https://leetcode.com/problems/last-person-to-fit-in-the-bus/
// Difficulty: Medium

// People queue for bus with weight limit 1000.
// Find the last person name that can board without exceeding limit.

// Time: O(n log n)
// Space: O(n)

type person struct {
	name   string
	weight int
	turn   int
}

func lastToFit(people []person) string {
  // Custom sort
	sort.Slice(people, func(i, j int) bool {
		return people[i].turn < people[j].turn
	})

	total := 0
	lastName := ""
	for _, p := range people {
		if total+p.weight <= 1000 {
			total += p.weight
			lastName = p.name
		} else {
			break
		}
	}
	return lastName
}

func main() {
	people := []person{
		{"Alice", 200, 1},
		{"Bob", 300, 2},
		{"Charlie", 400, 3},
		{"Dave", 200, 4},
	}
	fmt.Printf("%q (expected: \"Charlie\")\n", lastToFit(people))

	people2 := []person{
		{"John", 500, 1},
		{"Jane", 600, 2},
	}
	fmt.Printf("%q (expected: \"John\")\n", lastToFit(people2))
}
```
