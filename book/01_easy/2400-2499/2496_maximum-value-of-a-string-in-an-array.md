# 2496 — Maximum Value Of A String In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumValueOfAStringInAnArray(strs []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2496: Maximum Value of a String in an Array
// https://leetcode.com/problems/maximum-value-of-a-string-in-an-array/
// Difficulty: Easy
// Time O(n * m) | Space O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(MaximumValueOfAStringInAnArray([]string{"alic3", "bob", "3", "4", "00000"})) // 5
	fmt.Println(MaximumValueOfAStringInAnArray([]string{"1", "01", "001", "0001"}))           // 1
}

func MaximumValueOfAStringInAnArray(strs []string) int {
	maxVal := 0
	for _, s := range strs {
		val := 0
		if isNumeric(s) {
			val, _ = strconv.Atoi(s)
		} else {
			val = len(s)
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func isNumeric(s string) bool {
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return len(s) > 0
}
```
