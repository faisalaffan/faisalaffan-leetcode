# 1077 — Project Employees Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func projectEmployeesIII(project [][]int, employee [][]int) [][]int
```

> **💡 Hint:** Group employee experience by project, find max per project

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n) where n = len(project)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1077: Project Employees III
// https://leetcode.com/problems/project-employees-iii/
// Difficulty: Medium
//
// Approach: Group employee experience by project, find max per project
// Time: O(n log n) where n = len(project)
// Space: O(n)

import "fmt"

func main() {
	// (project_id, employee_id, experience_years)
	project := [][]int{{1, 1}, {1, 2}, {2, 3}, {2, 4}}
	employee := [][]int{{1, 5}, {2, 3}, {3, 7}, {4, 2}}
	fmt.Println(projectEmployeesIII(project, employee))
}

func projectEmployeesIII(project [][]int, employee [][]int) [][]int {
  // Membuat map (HashMap) — pencarian O(1)
	expMap := make(map[int]int)
	for _, e := range employee {
		expMap[e[0]] = e[1]
	}

	// For each project, find max experience and which employees have it
	type projInfo struct {
		maxExp int
		empID  int
	}
  // Membuat map (HashMap) — pencarian O(1)
	projMax := make(map[int]projInfo)

	for _, p := range project {
		projID, empID := p[0], p[1]
		exp := expMap[empID]

		if info, ok := projMax[projID]; !ok || exp > info.maxExp {
			projMax[projID] = projInfo{exp, empID}
		}
	}

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)
	for _, p := range project {
		projID, empID := p[0], p[1]
		info := projMax[projID]
		if empID == info.empID {
			result = append(result, []int{projID, empID})
		}
	}

	return result
}
```
