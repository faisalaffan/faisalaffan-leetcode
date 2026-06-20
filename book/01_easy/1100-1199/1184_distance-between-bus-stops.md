# 1184 — Distance Between Bus Stops

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func distanceBetweenBusStops(distance []int, start, destination int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1184: Distance Between Bus Stops
// https://leetcode.com/problems/distance-between-bus-stops/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1)) // 1
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2)) // 3
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3)) // 4
}

// LeetCode submission: distanceBetweenBusStops
func distanceBetweenBusStops(distance []int, start, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	forward := 0
	for i := start; i < destination; i++ {
		forward += distance[i]
	}
	total := 0
	for _, d := range distance {
		total += d
	}
	backward := total - forward
	if backward < forward {
		return backward
	}
	return forward
}
```
