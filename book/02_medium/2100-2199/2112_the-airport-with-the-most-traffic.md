# 2112 — The Airport With The Most Traffic

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func airportWithMostTraffic(flights [][]string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2112: The Airport With the Most Traffic
// https://leetcode.com/problems/the-airport-with-the-most-traffic/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func airportWithMostTraffic(flights [][]string) string {
  // Membuat map (HashMap) — pencarian O(1)
	traffic := make(map[string]int)
	for _, f := range flights {
		departure, arrival := f[0], f[1]
		traffic[departure]++
		traffic[arrival]++
	}

	maxTraffic := 0
	airport := ""
	for a, t := range traffic {
		if t > maxTraffic || (t == maxTraffic && (airport == "" || a < airport)) {
			maxTraffic = t
			airport = a
		}
	}
	return airport
}

func main() {
	// Test case 1
	flights1 := [][]string{{"JFK", "LGA"}, {"JFK", "LAX"}, {"LAX", "SFO"}, {"LGA", "ORD"}}
	fmt.Println("Test 1:", airportWithMostTraffic(flights1))
	// Expected: "JFK" (3 flights)

	// Test case 2
	flights2 := [][]string{{"A", "B"}, {"B", "C"}, {"C", "A"}}
	fmt.Println("Test 2:", airportWithMostTraffic(flights2))
	// Expected: "A" (tie, alphabetically first)

	// Test case 3
	fmt.Println("Test 3:", airportWithMostTraffic([][]string{{"X", "Y"}}))
	// Expected: "X" (tie, alphabetically: X < Y)
}
```
