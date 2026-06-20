# 0177 — Nth Highest Salary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func nthHighestSalary(salaries []int, n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n) average, Space: O(n) for quickselect  
**Kompleksitas Ruang:** O(n) for quickselect

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #177: Nth Highest Salary
// https://leetcode.com/problems/nth-highest-salary/
// Difficulty: Medium
// Time: O(n log n) average, Space: O(n) for quickselect

import (
	"fmt"
	"sort"
)

func nthHighestSalary(salaries []int, n int) int {
	if n <= 0 || n > len(salaries) {
		return 0
	}

	// Use sort + deduplicate
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	unique := []int{}
	for _, s := range salaries {
		if !seen[s] {
			seen[s] = true
			unique = append(unique, s)
		}
	}

	if n > len(unique) {
		return 0
	}

  // Custom sort dengan comparator
	sort.Slice(unique, func(i, j int) bool {
		return unique[i] > unique[j]
	})

	return unique[n-1]
}

func main() {
	fmt.Println(nthHighestSalary([]int{100, 200, 300, 200}, 2))
	fmt.Println(nthHighestSalary([]int{100, 100}, 2))
	fmt.Println(nthHighestSalary([]int{60, 70, 80, 90, 100}, 3))
}
```
