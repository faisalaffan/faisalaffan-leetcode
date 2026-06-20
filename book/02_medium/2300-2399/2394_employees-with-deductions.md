# 2394 — Employees With Deductions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func calculateDeductions(logs []Log, requiredHours []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2394: Employees With Deductions
// https://leetcode.com/problems/employees-with-deductions/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// Calculate actual hours worked per employee and compare with required hours.

import "fmt"

type Log struct {
	EmpID        int
	Timestamp    int
	IsLogin      bool
}

func main() {
	logs := []Log{
		{1, 100, true},
		{1, 200, false},
		{2, 50, true},
		{2, 150, false},
		{1, 300, true},
		{1, 400, false},
	}
	// Employee 1: total 200 min = 3.33 hrs, needs 4 hrs
	// Employee 2: total 100 min = 1.67 hrs, needs 4 hrs
	fmt.Println(calculateDeductions(logs, []int{4, 4})) // [1, 2] (both under)

	logs2 := []Log{
		{1, 0, true},
		{1, 240, false},
		{2, 0, true},
		{2, 480, false},
	}
	fmt.Println(calculateDeductions(logs2, []int{4, 8})) // [2]
}

func calculateDeductions(logs []Log, requiredHours []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	hours := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	loginTime := make(map[int]int)

	for _, l := range logs {
		if l.IsLogin {
			loginTime[l.EmpID] = l.Timestamp
		} else {
			if start, ok := loginTime[l.EmpID]; ok {
				hours[l.EmpID] += l.Timestamp - start
				delete(loginTime, l.EmpID)
			}
		}
	}

  // Alokasi slice integer
	result := make([]int, 0)
	for empID := 1; empID <= len(requiredHours); empID++ {
		workedMin := hours[empID]
		neededMin := requiredHours[empID-1] * 60
		if workedMin < neededMin {
			result = append(result, empID)
		}
	}
	return result
}
```
