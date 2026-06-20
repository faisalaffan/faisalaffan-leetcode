# 2303 — Calculate Amount Paid In Taxes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CalculateAmountPaidInTaxes(brackets [][]int, income int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2303: Calculate Amount Paid in Taxes
// https://leetcode.com/problems/calculate-amount-paid-in-taxes/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CalculateAmountPaidInTaxes([][]int{{3, 50}, {7, 10}, {12, 25}}, 10)) // 2.65
	fmt.Println(CalculateAmountPaidInTaxes([][]int{{1, 0}, {4, 25}, {5, 50}}, 2))    // 0.25
}

func CalculateAmountPaidInTaxes(brackets [][]int, income int) float64 {
	tax := 0.0
	prev := 0
	for _, b := range brackets {
		upper := b[0]
		percent := float64(b[1]) / 100.0
		taxable := min(upper, income) - prev
		if taxable > 0 {
			tax += float64(taxable) * percent
		}
		prev = upper
		if income <= upper {
			break
		}
	}
	return tax
}
```
