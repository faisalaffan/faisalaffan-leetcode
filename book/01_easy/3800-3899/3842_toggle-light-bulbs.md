# 3842 — Toggle Light Bulbs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ToggleLightBulbs(bulbs []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3842: Toggle Light Bulbs
// https://leetcode.com/problems/toggle-light-bulbs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ToggleLightBulbs([]int{10, 30, 20, 10}))
	fmt.Println(ToggleLightBulbs([]int{100, 100}))
	fmt.Println(ToggleLightBulbs([]int{1, 2, 3}))
}

// Time: O(n)
// Space: O(1)
func ToggleLightBulbs(bulbs []int) []int {
	var state [101]int
	for _, b := range bulbs {
		state[b] ^= 1
	}
	var result []int
	for i := 1; i <= 100; i++ {
		if state[i] == 1 {
			result = append(result, i)
		}
	}
	return result
}
```
