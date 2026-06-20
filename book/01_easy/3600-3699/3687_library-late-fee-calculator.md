# 3687 — Library Late Fee Calculator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LibraryLateFeeCalculator(daysLate []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3687: Library Late Fee Calculator
// https://leetcode.com/problems/library-late-fee-calculator/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(LibraryLateFeeCalculator([]int{5, 1, 7}))
	fmt.Println(LibraryLateFeeCalculator([]int{1, 1}))
}

// Time: O(n)
// Space: O(1)
func LibraryLateFeeCalculator(daysLate []int) int {
	ans := 0
	for _, x := range daysLate {
		if x == 1 {
			ans += 1
		} else if x > 5 {
			ans += 3 * x
		} else {
			ans += 2 * x
		}
	}
	return ans
}
```
