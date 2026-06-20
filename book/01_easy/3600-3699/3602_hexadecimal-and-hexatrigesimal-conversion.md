# 3602 — Hexadecimal And Hexatrigesimal Conversion

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func HexadecimalAndHexatrigesimalConversion(n int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3602: Hexadecimal and Hexatrigesimal Conversion
// https://leetcode.com/problems/hexadecimal-and-hexatrigesimal-conversion/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

const digits = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	fmt.Println(HexadecimalAndHexatrigesimalConversion(4))
}

// Time: O(log n)
// Space: O(log n)
func HexadecimalAndHexatrigesimalConversion(n int) string {
	sq := n * n
	cube := n * n * n
	return toBase(sq, 16) + toBase(cube, 36)
}

func toBase(num, base int) string {
	if num == 0 {
		return "0"
	}
	var res strings.Builder
	for num > 0 {
		res.WriteByte(digits[num%base])
		num /= base
	}

	// reverse
	s := []byte(res.String())
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
```
