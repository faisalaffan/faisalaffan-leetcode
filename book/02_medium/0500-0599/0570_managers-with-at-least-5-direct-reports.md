# 0570 — Managers With At Least 5 Direct Reports

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindManagers(employees [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #570: Managers with at Least 5 Direct Reports
// https://leetcode.com/problems/managers-with-at-least-5-direct-reports/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Employees: {id, name, department, managerId}
	// managerId == -1 means top-level manager
	employees := [][]int{
		{101, 1, 0, -1},  // John, manager
		{102, 2, 0, 101}, // Dan
		{103, 3, 0, 101}, // James
		{104, 4, 0, 101}, // Amy
		{105, 5, 0, 101}, // Ben
		{106, 6, 0, 101}, // Sam
		{107, 7, 1, -1},  // Ron, another manager
		{108, 8, 1, 107}, // Tom
	}
	fmt.Println(FindManagers(employees))
}

func FindManagers(employees [][]int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	reportCount := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	managerNames := make(map[int]int) // managerId -> manager name (for simplicity, just id)

	for _, emp := range employees {
		id, name, _, managerId := emp[0], emp[1], emp[2], emp[3]
		managerNames[id] = name
		if managerId != -1 {
			reportCount[managerId]++
		}
	}

	result := []int{}
	for mid, count := range reportCount {
		if count >= 5 {
			result = append(result, managerNames[mid])
		}
	}

	return result
}
```
