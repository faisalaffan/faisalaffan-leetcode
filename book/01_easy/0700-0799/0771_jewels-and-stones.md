# 0771 — Jewels And Stones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numJewelsInStones(jewels string, stones string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(j + s). Space: O(j).  
**Kompleksitas Ruang:** O(j).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #771: Jewels and Stones
// https://leetcode.com/problems/jewels-and-stones/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) // 3
	fmt.Println(numJewelsInStones("z", "ZZ"))       // 0
	fmt.Println(numJewelsInStones("", "abc"))       // 0
}

// numJewelsInStones counts how many stones are also jewels.
// Time: O(j + s). Space: O(j).
func numJewelsInStones(jewels string, stones string) int {
  // Membuat map (HashMap) — pencarian O(1)
	jSet := make(map[rune]bool)
	for _, c := range jewels {
		jSet[c] = true
	}
	count := 0
	for _, c := range stones {
		if jSet[c] {
			count++
		}
	}
	return count
}
```
