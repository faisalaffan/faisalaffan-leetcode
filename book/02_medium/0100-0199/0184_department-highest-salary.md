# 0184 — Department Highest Salary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func departmentHighestSalary(employees []Employee) []DeptSalary
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(d) where d is number of departments  
**Kompleksitas Ruang:** O(d) where d is number of departments

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #184: Department Highest Salary
// https://leetcode.com/problems/department-highest-salary/
// Difficulty: Medium
// Time: O(n), Space: O(d) where d is number of departments

import "fmt"

type Employee struct {
	Name       string
	Salary     int
	Department string
}

type DeptSalary struct {
	Department string
	Employee   string
	Salary     int
}

func departmentHighestSalary(employees []Employee) []DeptSalary {
  // Membuat map (HashMap) — pencarian O(1)
	deptMax := make(map[string]int)
	for _, e := range employees {
		if e.Salary > deptMax[e.Department] {
			deptMax[e.Department] = e.Salary
		}
	}

	result := []DeptSalary{}
	for _, e := range employees {
		if e.Salary == deptMax[e.Department] {
			result = append(result, DeptSalary{e.Department, e.Name, e.Salary})
		}
	}

	return result
}

func main() {
	emps := []Employee{
		{"Joe", 85000, "IT"},
		{"Jim", 90000, "IT"},
		{"Henry", 80000, "Sales"},
		{"Sam", 60000, "Sales"},
		{"Max", 90000, "IT"},
	}
	fmt.Println(departmentHighestSalary(emps))

	emps2 := []Employee{
		{"A", 50000, "Eng"},
		{"B", 60000, "Eng"},
	}
	fmt.Println(departmentHighestSalary(emps2))

	fmt.Println(departmentHighestSalary(nil))
}
```
