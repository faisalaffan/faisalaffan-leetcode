# 0848 — Shifting Letters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ShiftingLetters(s string, shifts []int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #848: Shifting Letters
// https://leetcode.com/problems/shifting-letters/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ShiftingLetters("abc", []int{3, 5, 9}))
	fmt.Println(ShiftingLetters("aaa", []int{1, 2, 3}))
	fmt.Println(ShiftingLetters("z", []int{52}))
}

// Time: O(n) | Space: O(n)
func ShiftingLetters(s string, shifts []int) string {
	n := len(s)
	// Calculate suffix sum of shifts
	for i := n - 2; i >= 0; i-- {
		shifts[i] = (shifts[i] + shifts[i+1]) % 26
	}

	res := []byte(s)
	for i := 0; i < n; i++ {
		res[i] = byte((int(res[i]-'a')+shifts[i])%26 + 'a')
	}

	return string(res)
}
```
