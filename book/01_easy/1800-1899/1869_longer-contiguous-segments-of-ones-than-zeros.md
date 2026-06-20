# 1869 — Longer Contiguous Segments Of Ones Than Zeros

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckZeroOnes(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1869: Longer Contiguous Segments of Ones Than Zeros
// https://leetcode.com/problems/longer-contiguous-segments-of-ones-than-zeros/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CheckZeroOnes(s string) bool {
	maxOnes, maxZeros := 0, 0
	curOnes, curZeros := 0, 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			curOnes++
			curZeros = 0
			if curOnes > maxOnes {
				maxOnes = curOnes
			}
		} else {
			curZeros++
			curOnes = 0
			if curZeros > maxZeros {
				maxZeros = curZeros
			}
		}
	}
	return maxOnes > maxZeros
}

func main() {
	fmt.Println(CheckZeroOnes("1101"))
	fmt.Println(CheckZeroOnes("111000"))
	fmt.Println(CheckZeroOnes("110100010"))
}
```
