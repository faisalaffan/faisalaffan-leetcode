# 1907 — Count Salary Categories

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountSalaryCategories(accounts [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1907: Count Salary Categories
// https://leetcode.com/problems/count-salary-categories/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// accounts: [account_id, income]
	accounts := [][]int{{1, 20000}, {2, 50000}, {3, 100000}, {4, 80000}, {5, 30000}}
	fmt.Println(CountSalaryCategories(accounts))
}

// Time: O(n), Space: O(1)
func CountSalaryCategories(accounts [][]int) []int {
	low := 0
	mid := 0
	high := 0

	for _, a := range accounts {
		income := a[1]
		if income < 20000 {
			low++
		} else if income <= 50000 {
			mid++
		} else {
			high++
		}
	}
	return []int{low, mid, high}
}
```
