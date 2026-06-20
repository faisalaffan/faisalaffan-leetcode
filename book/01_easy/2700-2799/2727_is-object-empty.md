# 2727 — Is Object Empty

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func IsObjectEmpty(obj interface{}) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2727: Is Object Empty
// https://leetcode.com/problems/is-object-empty/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Checks if array/slice is empty.

import "fmt"

func main() {
	fmt.Println(IsObjectEmpty([]int{}))
	fmt.Println(IsObjectEmpty([]int{1, 2}))
}

func IsObjectEmpty(obj interface{}) bool {
	switch v := obj.(type) {
	case []int:
		return len(v) == 0
	case []string:
		return len(v) == 0
	default:
		return false
	}
}
```
