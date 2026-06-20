# 0585 — Investments In 2016

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindInvestmentSum(insurance [][]interface{}) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #585: Investments in 2016
// https://leetcode.com/problems/investments-in-2016/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Insurance records: {pid, tiv_2015, tiv_2016, lat, lon}
	insurance := [][]interface{}{
		{1, 100, 200, 10, 20},
		{2, 100, 300, 30, 40},
		{3, 200, 400, 50, 60},
		{4, 100, 500, 10, 20}, // same lat/lon as pid=1
	}
	fmt.Println(FindInvestmentSum(insurance))
}

func FindInvestmentSum(insurance [][]interface{}) float64 {
  // Membuat map (HashMap) — pencarian O(1)
	tiv2015 := make(map[float64]int)
  // Membuat map (HashMap) — pencarian O(1)
	locationCount := make(map[string]int)
  // Membuat map (HashMap) — pencarian O(1)
	tiv2016Sum := make(map[int]float64)

	for _, record := range insurance {
		pid := record[0].(int)
		tiv15 := record[1].(float64)
		tiv16 := record[2].(float64)
		lat := record[3].(float64)
		lon := record[4].(float64)

		tiv2015[tiv15]++
		locKey := fmt.Sprintf("%f,%f", lat, lon)
		locationCount[locKey]++
		tiv2016Sum[pid] = tiv16
	}

	total := 0.0
	for _, record := range insurance {
		pid := record[0].(int)
		tiv15 := record[1].(float64)
		lat := record[3].(float64)
		lon := record[4].(float64)
		locKey := fmt.Sprintf("%f,%f", lat, lon)

		if tiv2015[tiv15] > 1 && locationCount[locKey] == 1 {
			total += tiv2016Sum[pid]
		}
	}

	return total
}
```
