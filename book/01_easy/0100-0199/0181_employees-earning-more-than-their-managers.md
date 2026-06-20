# 0181 — Employees Earning More Than Their Managers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func EmployeesEarningMoreThanTheirManagers() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


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
