# 0181 — Employees Earning More Than Their Managers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func EmployeesEarningMoreThanTheirManagers() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #181: Employees Earning More Than Their Managers
// https://leetcode.com/problems/employees-earning-more-than-their-managers/
// Difficulty: Easy

import "fmt"

func EmployeesEarningMoreThanTheirManagers() string {
	return "SELECT e.name AS Employee FROM Employee e JOIN Employee m ON e.managerId = m.id WHERE e.salary > m.salary"
}

func main() {
	fmt.Println(EmployeesEarningMoreThanTheirManagers())
}
```
