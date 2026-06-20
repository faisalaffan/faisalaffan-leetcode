# 1491 — Average Salary Excluding The Minimum And Maximum Salary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func average(salary []int) float64

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1491: Average Salary Excluding the Minimum and Maximum Salary
// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
// Difficulty: Easy
//
// LeetCode submission: func average(salary []int) float64

import "fmt"

func main() {
	fmt.Println(AverageSalaryExcludingTheMinimumAndMaximumSalary([]int{4000, 3000, 1000, 2000})) // 2500
	fmt.Println(AverageSalaryExcludingTheMinimumAndMaximumSalary([]int{1000, 2000, 3000}))       // 2000
}

// Time: O(n), Space: O(1)
func AverageSalaryExcludingTheMinimumAndMaximumSalary(salary []int) float64 {
	min, max := salary[0], salary[0]
	sum := 0
	for _, v := range salary {
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return float64(sum-min-max) / float64(len(salary)-2)
}
```
