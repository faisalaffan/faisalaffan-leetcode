# 2810 — Faulty Keyboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FaultyKeyboard(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2810: Faulty Keyboard
// https://leetcode.com/problems/faulty-keyboard/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
)

func main() {
	fmt.Println(FaultyKeyboard("string"))
	fmt.Println(FaultyKeyboard("poiinter"))
}

func FaultyKeyboard(s string) string {
	result := []rune{}
	for _, ch := range s {
		if ch == 'i' {
			// Reverse the current result
			for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
				result[i], result[j] = result[j], result[i]
			}
		} else {
			result = append(result, ch)
		}
	}
	return string(result)
}
```
