# 1556 — Thousand Separator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func thousandSeparator(n int) string

import (
	"fmt"
	"strconv"
)

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n), Space: O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1556: Thousand Separator
// https://leetcode.com/problems/thousand-separator/
// Difficulty: Easy
//
// LeetCode submission: func thousandSeparator(n int) string

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ThousandSeparator(987))      // "987"
	fmt.Println(ThousandSeparator(1234))     // "1.234"
	fmt.Println(ThousandSeparator(1000000))  // "1.000.000"
}

// Time: O(log n), Space: O(log n)
func ThousandSeparator(n int) string {
	s := strconv.Itoa(n)
	res := make([]byte, 0, len(s)+len(s)/3)
	for i, ch := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			res = append(res, '.')
		}
		res = append(res, byte(ch))
	}
	return string(res)
}
```
