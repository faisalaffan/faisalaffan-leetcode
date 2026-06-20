# 1518 — Water Bottles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numWaterBottles(numBottles int, numExchange int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1518: Water Bottles
// https://leetcode.com/problems/water-bottles/
// Difficulty: Easy
//
// LeetCode submission: func numWaterBottles(numBottles int, numExchange int) int

import "fmt"

func main() {
	fmt.Println(WaterBottles(9, 3))  // 13
	fmt.Println(WaterBottles(15, 4)) // 19
	fmt.Println(WaterBottles(5, 5))  // 6
}

// Time: O(log n), Space: O(1)
func WaterBottles(numBottles int, numExchange int) int {
	total := numBottles
	empty := numBottles
	for empty >= numExchange {
		newBottles := empty / numExchange
		total += newBottles
		empty = newBottles + empty%numExchange
	}
	return total
}
```
