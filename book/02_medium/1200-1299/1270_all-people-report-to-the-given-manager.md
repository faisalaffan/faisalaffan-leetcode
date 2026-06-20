# 1270 — All People Report To The Given Manager

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func allPeopleReportTo(employees [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1270: All People Report to the Given Manager
// https://leetcode.com/problems/all-people-report-to-the-given-manager/
// Difficulty: Medium [Paid]

// Find all employees who directly or indirectly report to the head
// (manager_id = 1 or manager_id is in the reporting chain).

// Time: O(n)
// Space: O(n)

func allPeopleReportTo(employees [][]int) []int {
	// employees[i] = [employee_id, manager_id]
	// Find all employees who report to employee_id=1 (directly or indirectly)

  // Membuat map (HashMap) — pencarian O(1)
	adj := make(map[int][]int)
	for _, e := range employees {
		empID, mgrID := e[0], e[1]
		if mgrID != 0 { // 0 means no manager (head)
			adj[mgrID] = append(adj[mgrID], empID)
		}
	}

  // Alokasi slice integer
	result := make([]int, 0)
	queue := []int{1}

	for len(queue) > 0 {
		mgr := queue[0]
		queue = queue[1:]
		for _, emp := range adj[mgr] {
			result = append(result, emp)
			queue = append(queue, emp)
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	employees := [][]int{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 3},
	}
	fmt.Printf("%v (expected: [2 3 4 5])\n", allPeopleReportTo(employees))

	employees2 := [][]int{
		{1, 0},
		{2, 1},
		{3, 1},
	}
	fmt.Printf("%v (expected: [2 3])\n", allPeopleReportTo(employees2))
}
```
