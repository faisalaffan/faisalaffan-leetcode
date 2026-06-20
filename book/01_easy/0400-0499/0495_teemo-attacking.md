# 0495 — Teemo Attacking

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TeemoAttacking(timeSeries []int, duration int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #495: Teemo Attacking
// https://leetcode.com/problems/teemo-attacking/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func TeemoAttacking(timeSeries []int, duration int) int {
	total := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(timeSeries)-1; i++ {
		total += min(duration, timeSeries[i+1]-timeSeries[i])
	}
	if len(timeSeries) > 0 {
		total += duration
	}
	return total
}

func main() {
	fmt.Println(TeemoAttacking([]int{1, 4}, 2))
	fmt.Println(TeemoAttacking([]int{1, 2}, 2))
}
```
