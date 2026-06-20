# 0168 — Excel Sheet Column Title

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConvertToTitle(columnNumber int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #168: Excel Sheet Column Title
// https://leetcode.com/problems/excel-sheet-column-title/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(log n)
func ConvertToTitle(columnNumber int) string {
	res := make([]byte, 0, 8)
	for columnNumber > 0 {
		columnNumber--
		res = append(res, byte('A'+columnNumber%26))
		columnNumber /= 26
	}
	// reverse
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func main() {
	fmt.Println(ConvertToTitle(1))
	fmt.Println(ConvertToTitle(28))
	fmt.Println(ConvertToTitle(701))
}
```
