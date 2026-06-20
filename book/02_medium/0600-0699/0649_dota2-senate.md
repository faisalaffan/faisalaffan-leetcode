# 0649 — Dota2 Senate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func predictPartyVictory(senate string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #649: Dota2 Senate
// https://leetcode.com/problems/dota2-senate/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(predictPartyVictory("RD"))
	fmt.Println(predictPartyVictory("RDD"))
	fmt.Println(predictPartyVictory("RRDDD"))
}

func predictPartyVictory(senate string) string {
	n := len(senate)
  // Alokasi slice integer
	radiant := make([]int, 0)
  // Alokasi slice integer
	dire := make([]int, 0)

	for i, c := range senate {
		if c == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		r := radiant[0]
		d := dire[0]
		radiant = radiant[1:]
		dire = dire[1:]

		if r < d {
			radiant = append(radiant, r+n)
		} else {
			dire = append(dire, d+n)
		}
	}

	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
```
