# 2797 — Partial Function With Placeholders

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func PartialFunctionWithPlaceholders(fn func(int, int, int) int, args ...interface
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2797: Partial Function with Placeholders
// https://leetcode.com/problems/partial-function-with-placeholders/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)
// Note: JS problem, adapted to Go. Partially applies arguments.

import "fmt"

func main() {
	add := func(a, b, c int) int { return a + b + c }
	partial := PartialFunctionWithPlaceholders(add, 1, nil, 3)
	fmt.Println(partial(2))
}

func PartialFunctionWithPlaceholders(fn func(int, int, int) int, args ...interface{}) func(int) int {
	return func(x int) int {
		realArgs := [3]int{}
		argIdx := 0
		for i, a := range args {
			if a == nil {
				realArgs[i] = x
			} else {
				realArgs[i] = a.(int)
				argIdx++
			}
		}
		return fn(realArgs[0], realArgs[1], realArgs[2])
	}
}
```
