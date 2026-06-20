# 3220 — Odd And Even Transactions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func oddAndEvenTransactions(transactions [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3220: Odd and Even Transactions
// https://leetcode.com/problems/odd-and-even-transactions/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func oddAndEvenTransactions(transactions [][]int) []int {
	oddSum, evenSum := 0, 0
	for _, t := range transactions {
		amount := t[1]
		if amount%2 == 0 {
			evenSum += amount
		} else {
			oddSum += amount
		}
	}
	return []int{oddSum, evenSum}
}

func main() {
	fmt.Println(oddAndEvenTransactions([][]int{{1, 10}, {2, 15}, {3, 20}})) // Expected: [25 30]
	fmt.Println(oddAndEvenTransactions([][]int{{1, 1}, {2, 2}, {3, 3}}))    // Expected: [4 2]
}
```
