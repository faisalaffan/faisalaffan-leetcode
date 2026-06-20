# 3338 — Second Highest Salary Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func secondHighestSalary(employees []Employee) []DeptSalary
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(e log e) Space: O(e)  
**Kompleksitas Ruang:** O(e)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3338: Second Highest Salary II
// https://leetcode.com/problems/second-highest-salary-ii/
// Difficulty: Medium
// Time: O(e log e) Space: O(e)

import (
	"fmt"
	"sort"
)

func main() {
	employees := []Employee{
		{1, "IT", 100000},
		{2, "IT", 80000},
		{3, "IT", 80000},
		{4, "HR", 90000},
		{5, "HR", 70000},
		{6, "HR", 60000},
	}
	fmt.Println(secondHighestSalary(employees))
}

type Employee struct {
	ID     int
	Dept   string
	Salary int
}

type DeptSalary struct {
	Dept  string
	Salary int
}

func secondHighestSalary(employees []Employee) []DeptSalary {
  // Membuat map (HashMap) — pencarian O(1)
	deptSalaries := make(map[string][]int)
	for _, e := range employees {
		deptSalaries[e.Dept] = append(deptSalaries[e.Dept], e.Salary)
	}

	var result []DeptSalary
	for dept, salaries := range deptSalaries {
  // Custom sort dengan comparator
		sort.Slice(salaries, func(i, j int) bool {
			return salaries[i] > salaries[j]
		})
		// Find second highest distinct salary
		seen := 1
		second := -1
		for _, s := range salaries {
			if s < salaries[0] {
				if seen == 1 || s < second {
					second = s
					seen++
				}
				if second != s {
					continue
				}
			}
		}
		// Actually simpler: just find 2nd distinct
		distinct := []int{salaries[0]}
		for _, s := range salaries {
			if s != distinct[len(distinct)-1] {
				distinct = append(distinct, s)
			}
		}
		if len(distinct) >= 2 {
			result = append(result, DeptSalary{dept, distinct[1]})
		}
	}

  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		return result[i].Dept < result[j].Dept
	})
	return result
}
```
