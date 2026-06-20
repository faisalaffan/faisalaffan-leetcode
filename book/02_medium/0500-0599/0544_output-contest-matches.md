# 0544 — Output Contest Matches

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindContestMatch(n int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #544: Output Contest Matches
// https://leetcode.com/problems/output-contest-matches/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FindContestMatch(2))
	fmt.Println(FindContestMatch(4))
	fmt.Println(FindContestMatch(8))
}

func FindContestMatch(n int) string {
	teams := make([]string, n)
	for i := 0; i < n; i++ {
		teams[i] = strconv.Itoa(i + 1)
	}

	for n > 1 {
		for i := 0; i < n/2; i++ {
			teams[i] = "(" + teams[i] + "," + teams[n-1-i] + ")"
		}
		n /= 2
	}

	return teams[0]
}
```
