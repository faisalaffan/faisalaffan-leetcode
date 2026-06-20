# 2703 — Return Length Of Arguments Passed

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ReturnLengthOfArgumentsPassed(args ...interface{}) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2703: Return Length of Arguments Passed
// https://leetcode.com/problems/return-length-of-arguments-passed/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns number of arguments.

import "fmt"

func main() {
	fmt.Println(ReturnLengthOfArgumentsPassed(1, 2, 3))
	fmt.Println(ReturnLengthOfArgumentsPassed("a", "b"))
}

func ReturnLengthOfArgumentsPassed(args ...interface{}) int {
	return len(args)
}
```
