# 0176 — Second Highest Salary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func secondHighestSalary(salaries []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #176: Second Highest Salary
// https://leetcode.com/problems/second-highest-salary/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

import "math"

func secondHighestSalary(salaries []int) int {
	first, second := math.MinInt32, math.MinInt32

	for _, s := range salaries {
		if s > first {
			second = first
			first = s
		} else if s > second && s < first {
			second = s
		}
	}

	if second == math.MinInt32 {
		return 0
	}
	return second
}

func main() {
	fmt.Println(secondHighestSalary([]int{100, 200, 300}))
	fmt.Println(secondHighestSalary([]int{100, 100}))
	fmt.Println(secondHighestSalary([]int{50, 100, 100, 75}))
}
```
