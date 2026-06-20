# 0681 — Next Closest Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func nextClosestTime(time string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) since there are at most 4^4 = 256 combinations  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #681: Next Closest Time
// https://leetcode.com/problems/next-closest-time/
// Difficulty: Medium [Paid]
// Time: O(1) since there are at most 4^4 = 256 combinations
// Space: O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(nextClosestTime("19:34"))
	fmt.Println(nextClosestTime("23:59"))
	fmt.Println(nextClosestTime("13:55"))
}

func nextClosestTime(time string) string {
  // Membuat map (HashMap) — pencarian O(1)
	digits := make(map[byte]bool)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(time); i++ {
		if time[i] != ':' {
			digits[time[i]] = true
		}
	}

	hours, _ := strconv.Atoi(time[:2])
	minutes, _ := strconv.Atoi(time[3:])

	current := hours*60 + minutes

	for elapsed := 1; elapsed <= 24*60; elapsed++ {
		t := (current + elapsed) % (24 * 60)
		h := t / 60
		m := t % 60

		hs := fmt.Sprintf("%02d%02d", h, m)
		valid := true
		for i := 0; i < 4; i++ {
			if !digits[hs[i]] {
				valid = false
				break
			}
		}

		if valid {
			return fmt.Sprintf("%02d:%02d", h, m)
		}
	}

	return ""
}
```
