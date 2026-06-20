# 1436 — Destination City

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func destCity(paths [][]string) string

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1436: Destination City
// https://leetcode.com/problems/destination-city/
// Difficulty: Easy
//
// LeetCode submission: func destCity(paths [][]string) string

import "fmt"

func main() {
	paths1 := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Println(DestinationCity(paths1)) // "Sao Paulo"

	paths2 := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	fmt.Println(DestinationCity(paths2)) // "A"
}

// Time: O(n), Space: O(n)
func DestinationCity(paths [][]string) string {
  // Membuat map (HashMap) — pencarian O(1)
	outgoing := make(map[string]bool, len(paths))
	for _, p := range paths {
		outgoing[p[0]] = true
	}
	for _, p := range paths {
		if !outgoing[p[1]] {
			return p[1]
		}
	}
	return ""
}
```
