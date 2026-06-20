# 2666 — Allow One Function Call

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func AllowOneFunctionCall(fn func(int) int) func(int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2666: Allow One Function Call
// https://leetcode.com/problems/allow-one-function-call/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Ensures fn is called at most once.

import "fmt"

func main() {
	fn := func(x int) int { return x * 2 }
	onceFn := AllowOneFunctionCall(fn)
	fmt.Println(onceFn(5))
	fmt.Println(onceFn(10))
}

func AllowOneFunctionCall(fn func(int) int) func(int) int {
	called := false
	return func(x int) int {
		if called {
			return 0
		}
		called = true
		return fn(x)
	}
}
```
