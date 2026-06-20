# 2889 — Reshape Data Pivot

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ReshapeDataPivot(data [][]string) map[string]map[string]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2889: Reshape Data: Pivot
// https://leetcode.com/problems/reshape-data-pivot/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we pivot data from [month, city, temperature] to [month, city1, city2, ...].

import "fmt"

func main() {
	// LeetCode name: pivotTable
	// Input: [month, city, temperature]
	data := [][]string{
		{"January", "London", "5"},
		{"January", "Paris", "7"},
		{"February", "London", "6"},
		{"February", "Paris", "8"},
	}
	fmt.Println(ReshapeDataPivot(data))
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: pivotTable
func ReshapeDataPivot(data [][]string) map[string]map[string]string {
  // Membuat map (HashMap) — pencarian O(1)
	result := make(map[string]map[string]string)
	for _, row := range data {
		month, city, temp := row[0], row[1], row[2]
		if result[month] == nil {
			result[month] = make(map[string]string)
		}
		result[month][city] = temp
	}
	return result
}
```
