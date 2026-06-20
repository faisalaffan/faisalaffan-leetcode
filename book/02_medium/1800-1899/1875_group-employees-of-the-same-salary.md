# 1875 — Group Employees Of The Same Salary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GroupEmployees(employees [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Membuat map (HashMap) — pencarian O(1)
	salaryMap := make(map[int][]int)
	for _, emp := range employees {
		id, salary := emp[0], emp[1]
		salaryMap[salary] = append(salaryMap[salary], id)
	}

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)
	for _, ids := range salaryMap {
		if len(ids) >= 2 {
  // Urutkan secara ascending — O(n log n)
			sort.Ints(ids)
			result = append(result, ids)
		}
	}

	// Sort by first employee ID for deterministic output
  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}
```
