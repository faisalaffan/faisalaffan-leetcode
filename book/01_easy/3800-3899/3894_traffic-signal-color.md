# 3894 — Traffic Signal Color

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TrafficSignalColor(timer int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3894: Traffic Signal Color
// https://leetcode.com/problems/traffic-signal-color/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TrafficSignalColor(60))
	fmt.Println(TrafficSignalColor(5))
	fmt.Println(TrafficSignalColor(0))
	fmt.Println(TrafficSignalColor(30))
	fmt.Println(TrafficSignalColor(90))
}

// Time: O(1)
// Space: O(1)
func TrafficSignalColor(timer int) string {
	if timer == 0 {
		return "Green"
	}
	if timer == 30 {
		return "Orange"
	}
	if timer > 30 && timer <= 90 {
		return "Red"
	}
	return "Invalid"
}
```
