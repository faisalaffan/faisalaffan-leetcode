# 0860 — Lemonade Change

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func lemonadeChange(bills []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #860: Lemonade Change
// https://leetcode.com/problems/lemonade-change/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20})) // true
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20})) // false
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 5, 20, 5, 10, 5, 20})) // true
}

// lemonadeChange checks if we can provide correct change for each customer.
// Time: O(n). Space: O(1).
func lemonadeChange(bills []int) bool {
	fives, tens := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			fives++
		case 10:
			if fives == 0 {
				return false
			}
			fives--
			tens++
		case 20:
			if tens > 0 && fives > 0 {
				tens--
				fives--
			} else if fives >= 3 {
				fives -= 3
			} else {
				return false
			}
		}
	}
	return true
}
```
