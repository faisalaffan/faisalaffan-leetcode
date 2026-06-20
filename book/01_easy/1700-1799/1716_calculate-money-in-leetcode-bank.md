# 1716 — Calculate Money In Leetcode Bank

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TotalMoney(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1716: Calculate Money in Leetcode Bank
// https://leetcode.com/problems/calculate-money-in-leetcode-bank/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func TotalMoney(n int) int {
	weeks := n / 7
	days := n % 7
	// Sum of arithmetic progression: first week = 28, each week adds 7
	total := weeks*28 + 7*weeks*(weeks-1)/2
	// Remaining days
	total += days*(weeks+1) + days*(days-1)/2
	return total
}

func main() {
	fmt.Println(TotalMoney(4))
	fmt.Println(TotalMoney(10))
	fmt.Println(TotalMoney(20))
}
```
