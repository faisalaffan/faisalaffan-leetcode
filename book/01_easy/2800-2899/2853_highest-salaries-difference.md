# 2853 — Highest Salaries Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func HighestSalariesDifference(salaries []struct {
	Name       string
	Salary     int
	Department string
}) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2853: Highest Salaries Difference
// https://leetcode.com/problems/highest-salaries-difference/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)
// Note: SQL problem, adapted to Go. Difference between highest salaries in two departments.

import "fmt"

func main() {
	salaries := []struct {
		Name       string
		Salary     int
		Department string
	}{
		{"Alice", 100000, "Engineering"},
		{"Bob", 90000, "Engineering"},
		{"Charlie", 80000, "Marketing"},
		{"David", 95000, "Marketing"},
	}
	fmt.Println(HighestSalariesDifference(salaries))
}

func HighestSalariesDifference(salaries []struct {
	Name       string
	Salary     int
	Department string
}) int {
	maxEng, maxMkt := 0, 0
	for _, s := range salaries {
		if s.Department == "Engineering" && s.Salary > maxEng {
			maxEng = s.Salary
		} else if s.Department == "Marketing" && s.Salary > maxMkt {
			maxMkt = s.Salary
		}
	}
	if maxEng > maxMkt {
		return maxEng - maxMkt
	}
	return maxMkt - maxEng
}
```
