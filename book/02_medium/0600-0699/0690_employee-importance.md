# 0690 — Employee Importance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getImportance(employees []*Employee, id int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #690: Employee Importance
// https://leetcode.com/problems/employee-importance/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	employees := []*Employee{
		{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
		{Id: 2, Importance: 3, Subordinates: []int{}},
		{Id: 3, Importance: 3, Subordinates: []int{}},
	}
	fmt.Println(getImportance(employees, 1))
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
  // Membuat map (HashMap) — pencarian O(1)
	empMap := make(map[int]*Employee)
	for _, e := range employees {
		empMap[e.Id] = e
	}

	var dfs func(id int) int
	dfs = func(id int) int {
		emp := empMap[id]
		total := emp.Importance
		for _, subId := range emp.Subordinates {
			total += dfs(subId)
		}
		return total
	}

	return dfs(id)
}
```
